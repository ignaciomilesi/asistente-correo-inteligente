package widgets

import (
	"os/exec"
	"regexp"
	"strings"

	g "github.com/AllenDang/giu"
)

type inputMultiLineConEdicion struct {
	input            *string
	altura           float32
	seleccionEdicion bool
	t                bool
}

func NuevoInputMultiLineConEdicion(inputTexto *string, alturaWidget float32) *inputMultiLineConEdicion {
	return &inputMultiLineConEdicion{
		input:  inputTexto,
		altura: alturaWidget,
	}
}

func (di *inputMultiLineConEdicion) Build() {

	var textoAMostrar string

	if !di.t {
		di.parsearInput()
		di.t = true
	}
	if di.seleccionEdicion {

		g.InputTextMultiline(di.input).Size(g.Auto, di.altura).Build()
		textoAMostrar = "Volver a presentación"

	} else {

		g.Child().Size(g.Auto, di.altura).Border(true).Layout(
			di.parsearInput()...,
		/*
			g.Label("prueba label").Wrapped(true),
			g.Link("PruebaLink").OnClick(func() {
				exec.Command("rundll32", "url.dll,FileProtocolHandler", "wwww.google.com").Start()
			}),
			g.BulletText("Prueba 1"),
			g.BulletText("Prueba 2"),*/
		).Build()

		//g.Child().Size(g.Auto, di.altura).Border(true).Layout(
		//	g.Markdown(*di.input)).Build()
		textoAMostrar = "Editar"
	}
	g.ContextMenu().Layout(
		g.Selectable(textoAMostrar).OnClick(func() {
			di.seleccionEdicion = !di.seleccionEdicion
		}),
	).Build()

}

func (di inputMultiLineConEdicion) parsearInput() (layout []g.Widget) {

	/*
		Comandos:	identificado entre ::
		Tipo:		identificado entre [] (identifica el tipo de comando link, bullet, etc..)
		Contenido: 	seguido al tipo, es lo que se procesa
		Etiqueta: 	identificado entre(), no siempre se usa. es el texto visible
	*/

	//localizadorDeComando := `::(.*?)::`
	localizadorDeTipo := `\[(.*?)\]`
	localizadorDeEtiqueta := `\((.*?)\)`

	partes := strings.Split(*di.input, "::")

	for _, parte := range partes {

		//comprueba salto de linea solitario
		if len(parte) == 1 && strings.HasPrefix(parte, "\n") {
			continue
		}
		// eliminamos los saltos de linea al inicio
		if strings.HasPrefix(parte, "\n") {
			parte = parte[1:]
		}

		reTipo := regexp.MustCompile(localizadorDeTipo)
		matchTipo := reTipo.FindString(parte)

		matchTipo = strings.ToLower(matchTipo)

		switch matchTipo {
		case "[link]":
			reEtiqueta := regexp.MustCompile(localizadorDeEtiqueta)
			matchEtiqueta := reEtiqueta.FindString(parte)

			link := parte[len(matchTipo) : len(parte)-len(matchEtiqueta)]

			if len(matchEtiqueta) == 0 {
				matchEtiqueta = parte[len(matchTipo) : len(parte)-len(matchEtiqueta)]
			} else {
				matchEtiqueta = matchEtiqueta[1 : len(matchEtiqueta)-1]
			}
			layout = append(layout,
				g.Link(matchEtiqueta).OnClick(func() {
					exec.Command("rundll32", "url.dll,FileProtocolHandler", link).Start()
				}))

		case "[sap]":
			layout = append(layout, di.linkSAP(parte))

		case "[bullet]", "[-]":
			layout = append(layout, g.BulletText(parte[len(matchTipo):]))

		default:
			layout = append(layout, g.Label(parte).Wrapped(true))
		}

	}

	return
}
