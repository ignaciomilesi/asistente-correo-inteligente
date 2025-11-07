package main

import (
	"fmt"
	"go-ollama/pkg/lectormail"
	"go-ollama/pkg/ollama"
	"path/filepath"
	"sync"
)

func main() {

	paths := []string{"prueba.msg", "prueba2.msg", "prueba3.msg", "prueba4.msg"}

	op := ollama.Processor{}

	op.Iniciar()

	var wg sync.WaitGroup

	for _, p := range paths {
		wg.Go(func() {
			pp := filepath.Join("mailTest", p)
			r := lectormail.ExtraerPropiedades(pp)

			msg := fmt.Sprintln("----------------------", p, "----------------------")
			msg += fmt.Sprintf("%-8s : %s\n", "Asunto", r.Asunto)
			msg += fmt.Sprintf("%-8s : %s\n", "De", r.De)
			msg += fmt.Sprintf("%-8s : %s\n", "Para", r.Para)
			msg += fmt.Sprintf("%-8s : %s\n", "CC", r.CC)
			msg += fmt.Sprintf("%-8s : %s\n", "Fecha", r.Fecha.Format("02/01/2006"))
			//msg += fmt.Sprintf("%-8s : %s\n", "Ultimo Mail", r.UltimoMail)
			msg += fmt.Sprintf("%-8s : ", "Resumen")

			if resumen, err := op.Resumir(r.UltimoMail); err == nil {
				msg += fmt.Sprintf("%s", resumen)
			} else {
				msg += fmt.Sprintf("No se ha podido resumir el mail. Error: %s", err)
			}

			fmt.Println(msg)
		})
	}

	wg.Wait()

}
