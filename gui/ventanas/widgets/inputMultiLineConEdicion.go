package widgets

import (
	"image/color"
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

	var textoMostrarSelectable string
	var mostrarComandos int32 // 0 muestra, 1 no muestra

	if di.seleccionEdicion {
		mostrarComandos = 0
	} else {
		mostrarComandos = 1
	}

	g.Align(g.AlignRight).To(
		g.Style().SetFontSize(10).SetColor(g.StyleColorText, color.Gray{100}).To(
			g.Stack(mostrarComandos,
				g.Label("(Comandos)"),
				g.Label(""),
			),
		),
	).Build()

	if di.seleccionEdicion {

		g.ContextMenu().MouseButton(g.MouseButtonLeft).Layout(
			g.Style().SetColor(g.StyleColorText, color.Gray{100}).To(
				g.Label("Comandos: !tipo:contenido|etiqueta -> Deben estar en una nueva linea"),
				g.Label("Ejemplo: !l:www.google.com|etiqueta - genera un hipervínculo"),
				g.Separator(),
				g.Label("Comandos disponibles:"),
				g.BulletText("!link o !l: hipervínculo. Con etiqueta se puede cambiar el texto a mostrar"),
				g.BulletText("!sap o !s: hipervínculo al detalle del SAP"),
				g.BulletText("!bullet o !-: bullet o símbolo viñeta"),
			),
		).Build()
		g.InputTextMultiline(di.input).Size(g.Auto, di.altura).Build()
		textoMostrarSelectable = "Finalizar edición"

	} else {

		g.Child().Size(g.Auto, di.altura).Border(true).Layout(
			di.parsearInput()...,
		).Build()
		textoMostrarSelectable = "Editar"
	}

	g.ContextMenu().Layout(
		g.Selectable(textoMostrarSelectable).OnClick(func() {
			di.seleccionEdicion = !di.seleccionEdicion
		}),
	).Build()
}

func (di inputMultiLineConEdicion) parsearInput() (layout []g.Widget) {

	/*
		Comandos:	identificado como !tipo:contenido|etiqueta
		Tipo:		identificado a la derecha ! (identifica el tipo de comando link, bullet, etc..)
		Contenido: 	identificado a la derecha :, es lo que se procesa
		Etiqueta: 	identificado a la derecha del |, es opcional. es el texto visible
	*/

	localizadorDeComando := `!(.+?):(.+?)(?:\|(.*))?$`

	partes := strings.Split(*di.input, "\n") // rompo el texto en saltos de linea

	for _, parte := range partes {

		reTipo := regexp.MustCompile(localizadorDeComando)
		matchComando := reTipo.FindStringSubmatch(parte)

		if len(matchComando) == 0 {
			layout = append(layout, g.Label(parte).Wrapped(true))
			continue
		}

		switch matchComando[1] {
		case "link", "l":
			var matchEtiqueta string
			if len(matchComando[3]) == 0 {
				matchEtiqueta = matchComando[2]
			} else {
				matchEtiqueta = matchComando[3]
			}
			layout = append(layout,
				g.Link(matchEtiqueta).OnClick(func() {
					exec.Command("rundll32", "url.dll,FileProtocolHandler", matchComando[2]).Start()
				}))

		case "sap", "s":
			layout = append(layout, di.linkSAP(matchComando[2], matchComando[3]))

		case "bullet", "-":
			layout = append(layout, g.BulletText(matchComando[2]))

		default:
			layout = append(layout, g.Label(parte).Wrapped(true))
		}

	}

	return
}
