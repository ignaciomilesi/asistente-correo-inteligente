package main

import (
	"fmt"
	"go-ollama/pkg/lectormail"
	"path/filepath"
)

func main() {

	cuerpo := lectormail.ExtraerPropiedades("mailTest\\prueba.msg").Cuerpo

	fmt.Println(cuerpo)

	return
	paths := []string{"prueba.msg", "prueba2.msg", "prueba3.msg", "prueba4.msg"}

	for _, p := range paths {
		pp := filepath.Join("mailTest", p)
		r := lectormail.ExtraerPropiedades(pp)
		fmt.Println("------------------------------------------")
		fmt.Printf("%-8s : %s\n", "Asunto", r.Asunto)
		fmt.Printf("%-8s : %s\n", "De", r.De)
		fmt.Printf("%-8s : %s\n", "Para", r.Para)
		fmt.Printf("%-8s : %s\n", "CC", r.CC)
		fmt.Printf("%-8s : %s\n", "Fecha", r.Fecha.Format("02/01/2006"))
	}

}
