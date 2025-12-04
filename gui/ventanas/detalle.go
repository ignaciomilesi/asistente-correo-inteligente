package ventanas

import (
	"fmt"
	"os"
	"os/exec"

	"go-ollama/gui/ventanas/widgets"
	"go-ollama/service"

	g "github.com/AllenDang/giu"
)

type serviceInterfaceDetallePendiente interface {
	ObtenerDetallePendiente(int) (service.PendienteCompleto, error)
	ObtenerListaAdjunto(int) ([]service.Adjunto, error)
	ObtenerListaAvance(int) ([]service.Avance, error)
	ObtenerListaUsuarios() ([]string, error)
}

type VentanaDetallePendiente struct {
	Service serviceInterfaceDetallePendiente

	ventanaActiva bool
	pendiente     service.PendienteCompleto
	avances       []service.Avance
	adjuntos      []service.Adjunto
	listaUsuarios []string

	descripcionInput g.Widget
	cierreInput      g.Widget
}

func (v *VentanaDetallePendiente) EsVentanaActiva() bool {
	return true
}

func (v *VentanaDetallePendiente) Actualizar() {

	/*
		chanPendiente := make(chan service.PendienteCompleto)
		go func(ch chan<- service.PendienteCompleto) {
			dp, err := v.Service.ObtenerDetallePendiente(4)
			if err != nil {
				fmt.Printf("No se pudo actualizar el detalle.\n%v\n", err)
			}
			ch <- dp
		}(chanPendiente)

		v.pendiente = <-chanPendiente*/

	var err error
	v.pendiente, err = v.Service.ObtenerDetallePendiente(4)
	v.avances, err = v.Service.ObtenerListaAvance(4)
	v.adjuntos, err = v.Service.ObtenerListaAdjunto(4)
	listaUsuarios, err := v.Service.ObtenerListaUsuarios()
	v.listaUsuarios = append(v.listaUsuarios, "") //Ocupo el lugar 0
	v.listaUsuarios = append(v.listaUsuarios, listaUsuarios...)

	if err != nil {
		fmt.Printf("No se pudo actualizar el detalle.\n%v\n", err)
	}

	v.descripcionInput = widgets.NuevoInputMultiLineConEdicion(&v.pendiente.Descripcion, 275)
	v.cierreInput = widgets.NuevoInputMultiLineConEdicion(&v.pendiente.Cierre, 100)
}

func (v *VentanaDetallePendiente) Esquema() {

	ventana := g.Window("Detalle del pendiente")
	ventana.Size(1000, 600)

	ventana.Layout(

		g.PrepareMsgbox(), //necesario para los msgBox
		g.Row(
			g.Labelf("ID: %d", v.pendiente.ID),
			g.Condition(v.pendiente.Finalizada,
				g.Label(" -  Finalizado"),
				g.Label("")),

			g.Condition(func() bool {
				ancho, _ := ventana.CurrentSize()
				return ancho > 450
			}(),
				g.Align(g.AlignRight).To(
					g.Labelf("Fecha de Inicio: %s    /    Fecha cierre: %s",
						v.pendiente.Fecha_iniciada.Time.Format("02-01-2006"),
						v.pendiente.Fecha_finalizada.Time.Format("02-01-2006")),
				),
				g.Dummy(1, 1)),
		),

		g.Separator(),

		// Titulo
		g.Dummy(5, 5),
		g.Style().
			SetFontSize(36).To(
			g.Label(v.pendiente.Titulo),
		),

		g.Dummy(5, 5),

		g.Row(
			g.Label("Asignado: "),
			g.Combo("", v.listaUsuarios[v.pendiente.Asignado], v.listaUsuarios, &v.pendiente.Asignado).Size(300),
		),

		g.Row(
			g.Label("Estado: "),
			g.InputText(&v.pendiente.Estado).Size(g.Auto),
		),
		g.Dummy(5, 5),

		g.Separator(),

		g.Dummy(5, 5),
		v.descripcionInput,

		g.TabBar().TabItems(

			g.TabItem("Avances").Layout(
				g.Dummy(5, 5),

				g.Table().Columns(
					g.TableColumn("Fecha").Flags(g.TableColumnFlagsNoResize|g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(80),
					g.TableColumn("Descripción").Flags(g.TableColumnFlagsNoResize),
				).Rows(
					v.tablaDeAvances()...,
				).Flags(g.TableFlagsRowBg|g.TableFlagsBorders).NoHeader(true),
			),

			g.TabItem("Adjuntos").Layout(
				g.Dummy(5, 5),
				g.Table().Columns(
					g.TableColumn("Descripción").Flags(g.TableColumnFlagsNoResize),
				).Rows(
					v.tablaDeAdjuntos()...,
				).Flags(g.TableFlagsRowBg|g.TableFlagsBorders).NoHeader(true),
			),

			g.TabItem("Cierre").Layout(
				v.cierreInput,
			),
		),
	)

}

func (v VentanaDetallePendiente) tablaDeAvances() (filas []*g.TableRowWidget) {

	for _, avance := range v.avances {

		tableRow := g.TableRow(

			// Fecha
			g.Selectable(avance.Fecha_Avance.Time.Format("02-01-2006")).
				OnDClick(func() {

					_, err := os.Stat(avance.Ubicacion_mail)
					if err != nil {

						g.Msgbox("Error", err.Error())

						fmt.Println(err)
						return
					}
					exec.Command("rundll32", "url.dll,FileProtocolHandler", avance.Ubicacion_mail).Start()

				}).Flags(g.SelectableFlagsSpanAllColumns),

			// Descripción
			g.Label(avance.Descripcion).Wrapped(true),
		).MinHeight(30)
		filas = append(filas, tableRow)
	}

	return
}

func (v VentanaDetallePendiente) tablaDeAdjuntos() (filas []*g.TableRowWidget) {

	for _, adjunto := range v.adjuntos {

		tableRow := g.TableRow(

			// Descripción
			g.Selectable(adjunto.Descripcion).
				OnDClick(func() {

					_, err := os.Stat(adjunto.Ubicacion_archivo)
					if err != nil {

						g.Msgbox("Error", err.Error())

						fmt.Println(err)
						return
					}
					exec.Command("rundll32", "url.dll,FileProtocolHandler", adjunto.Ubicacion_archivo).Start()

				}).Flags(g.SelectableFlagsSpanAllColumns),
		).MinHeight(30)

		filas = append(filas, tableRow)
	}

	return
}
