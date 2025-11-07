package lectormail

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf16"

	"github.com/richardlehane/mscfb"
)

type MailPropiedades struct {
	Asunto         string
	De             string
	Para           string
	CC             string
	UltimoMail     string // Cuerpo del ultimo mail
	CadenaDeMail   string // Cuerpo de toda la cadena de mail
	CuerpoCompleto string // Ultimo mail + cadena
	Cabecera       string
	Fecha          time.Time
}

func ExtraerPropiedades(mailPath string) (mailParseado MailPropiedades) {

	file, err := os.Open(mailPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := mscfb.New(file)
	if err != nil {
		panic(err)
	}

	// Mapeo de los streams que queremos leer
	wanted := map[string]*string{
		"__substg1.0_0037001F": &mailParseado.Asunto,
		"__substg1.0_0C1A001F": &mailParseado.De,
		"__substg1.0_0E04001F": &mailParseado.Para,
		"__substg1.0_0E03001F": &mailParseado.CC,
		"__substg1.0_1000001F": &mailParseado.CuerpoCompleto,
		"__substg1.0_007D001F": &mailParseado.Cabecera,
	}

	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {

		if target, ok := wanted[entry.Name]; ok {

			buf := make([]byte, entry.Size)
			doc.Read(buf)
			*target = utf16ToString(buf)
		}
	}

	// limpio el texto del cuerpo
	mailParseado.CuerpoCompleto = limpiarTextoEmail(mailParseado.CuerpoCompleto)

	// extraigo la fecha de la cabecera
	mailParseado.Fecha, err = extraerFecha(mailParseado.Cabecera)
	if err != nil {
		fmt.Println(err)
	}

	UltimoMail, Cadena, ok := strings.Cut(mailParseado.CuerpoCompleto, "De: ")

	if ok {
		mailParseado.UltimoMail = UltimoMail
		mailParseado.CadenaDeMail = Cadena

	} else {
		mailParseado.UltimoMail = mailParseado.CuerpoCompleto
	}

	return mailParseado
}

// Convierte []byte (UTF-16LE) a string UTF-8
func utf16ToString(b []byte) string {
	if len(b) < 2 {
		return ""
	}
	u16s := make([]uint16, len(b)/2)
	_ = binary.Read(bytes.NewReader(b), binary.LittleEndian, &u16s)
	runes := utf16.Decode(u16s)
	return string(runes)
}

func extraerFecha(cabecera string) (time.Time, error) {
	// Tomar todo después del último ';'
	parts := strings.Split(cabecera, "Date: ")
	if len(parts) < 2 {
		return time.Time{}, fmt.Errorf("no se encontró fecha en la cadena")
	}

	partsConFecha := parts[len(parts)-1] //me quedo con la parte de la fecha

	datePart := strings.TrimSpace(partsConFecha[:25])

	// Formato esperado:
	layouts := []string{
		"Mon, 02 Jan 2006 15:04:05",
		"02 Jan 2006 15:04:05",
		"Mon, 02 Jan 2006 15:04:05 -",
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, datePart)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("no se pudo parsear la fecha: %v", err)
}

func limpiarTextoEmail(texto string) string {

	// Normaliza los saltos de línea a '\n' (por si vienen como '\r\n')

	texto = strings.ReplaceAll(texto, "\r", "\n")

	// Reemplaza saltos de línea múltiples (2 o más) por uno solo
	reSaltos := regexp.MustCompile(`\n{2,}`)
	texto = reSaltos.ReplaceAllString(texto, "\n")

	// Reemplaza espacios múltiples por uno solo
	reEspacios := regexp.MustCompile(` +`)
	texto = reEspacios.ReplaceAllString(texto, " ")

	// Elimina espacios o saltos al principio y final
	texto = strings.TrimSpace(texto)

	return texto
}
