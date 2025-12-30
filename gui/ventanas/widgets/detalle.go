package widgets

import (
	"fmt"
	"os"
	"os/exec"

	"go-ollama/service"

	g "github.com/AllenDang/giu"
)

type serviceInterfaceDetallePendiente interface {
	ObtenerDetallePendiente(int) (service.PendienteCompleto, error)
	ObtenerListaAdjunto(int) ([]service.Adjunto, error)
	ObtenerListaAvance(int) ([]service.Avance, error)
	ObtenerListaUsuarios() ([]string, error)
}

type detallePendiente struct {
	service serviceInterfaceDetallePendiente

	ventanaActiva bool
	pendiente     service.PendienteCompleto
	avances       []service.Avance
	adjuntos      []service.Adjunto
	listaUsuarios []string

	descripcionInput g.Widget
	cierreInput      g.Widget
}

func NuevoDetallePendiente(serviceBaseDatos serviceInterfaceDetallePendiente, idPendiente int) *detallePendiente {

	v := detallePendiente{
		service: serviceBaseDatos,
	}

	var actualizacionPendiente = make(chan service.PendienteCompleto)
	var actualizacionAvances = make(chan []service.Avance)
	var actualizacionAdjuntos = make(chan []service.Adjunto)
	var actualizacionListaUsuarios = make(chan []string)

	// pendientes
	go func() {
		dp, err := v.service.ObtenerDetallePendiente(idPendiente)
		if err != nil {
			fmt.Printf("No se pudo actualizar el detalle.\n%v\n", err)
		}
		actualizacionPendiente <- dp
	}()
	// avances
	go func() {
		a, err := v.service.ObtenerListaAvance(idPendiente)
		if err != nil {
			fmt.Printf("No se pudo actualizar los avances.\n%v\n", err)
		}
		actualizacionAvances <- a

	}()
	// adjuntos
	go func() {
		a, err := v.service.ObtenerListaAdjunto(idPendiente)
		if err != nil {
			fmt.Printf("No se pudo actualizar los avances.\n%v\n", err)
		}
		actualizacionAdjuntos <- a

	}()
	// usuarios
	v.listaUsuarios = []string{""} //Ocupo el lugar 0
	go func() {
		lista := []string{""} //Ocupo el lugar 0
		u, err := v.service.ObtenerListaUsuarios()
		if err != nil {
			fmt.Printf("No se pudo actualizar la lista de usuarios.\n%v\n", err)
		}

		actualizacionListaUsuarios <- append(lista, u...)

	}()

	// función encargada de realizar la actualización
	go func() {
		for i := 0; i < 4; i++ {
			select {
			// pendientes
			case actualizacion := <-actualizacionPendiente:
				v.pendiente = actualizacion

			// avances
			case actualizacion := <-actualizacionAvances:
				v.avances = actualizacion

			// adjuntos
			case actualizacion := <-actualizacionAdjuntos:
				v.adjuntos = actualizacion

			// usuarios
			case actualizacion := <-actualizacionListaUsuarios:
				v.listaUsuarios = actualizacion
			}

		}

	}()

	v.descripcionInput = NuevoInputMultiLineConEdicion(&v.pendiente.Descripcion, 275)
	v.cierreInput = NuevoInputMultiLineConEdicion(&v.pendiente.Cierre.String, 100)

	return &v
}
func (v *detallePendiente) EsVentanaActiva() bool {
	return true
}

func (v *detallePendiente) Build() {

	g.PrepareMsgbox().Build() //necesario para los msgBox
	g.Row(
		g.Labelf("ID: %d", v.pendiente.ID),
		g.Condition(v.pendiente.Finalizada,
			g.Label(" -  Finalizado"),
			g.Label("")),

		g.Condition(func() bool {
			ancho, _ := g.GetAvailableRegion()
			return ancho > 450
		}(),
			g.Align(g.AlignRight).To(
				g.Labelf("Fecha de Inicio: %s    /    Fecha cierre: %s",
					v.pendiente.Fecha_iniciada.Time.Format("02-01-2006"),
					v.pendiente.Fecha_finalizada.Time.Format("02-01-2006")),
			),
			g.Dummy(1, 1)),
	).Build()

	g.Separator().Build()

	// Titulo
	g.Dummy(5, 5).Build()
	g.Style().
		SetFontSize(32).To(
		g.Label(v.pendiente.Titulo),
	).Build()

	g.Dummy(5, 5).Build()

	g.Row(
		g.Label("Asignado: "),

		g.Condition(v.pendiente.Asignado.Valid && len(v.listaUsuarios) > int(v.pendiente.Asignado.Int32),
			g.Combo("", v.listaUsuarios[v.pendiente.Asignado.Int32], v.listaUsuarios, &v.pendiente.Asignado.Int32).Size(300),
			g.Dummy(5, 5),
		),
	).Build()

	g.Row(
		g.Label("Estado: "),
		g.InputText(&v.pendiente.Estado).Size(g.Auto),
	).Build()
	g.Dummy(3, 3).Build()

	g.Separator().Build()
	v.descripcionInput.Build()

	g.Dummy(3, 3).Build()
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
	).Build()

}

func (v detallePendiente) tablaDeAvances() (filas []*g.TableRowWidget) {

	for _, avance := range v.avances {

		tableRow := g.TableRow(

			// Fecha
			g.Selectable(avance.Fecha_Avance.Time.Format("02-01-2006")).
				OnDClick(func() {

					if avance.Ubicacion_mail.String == "" {
						return
					}

					_, err := os.Stat(avance.Ubicacion_mail.String)
					if err != nil {

						g.Msgbox("Error", err.Error())

						fmt.Println(err)
						return
					}
					exec.Command("rundll32", "url.dll,FileProtocolHandler", avance.Ubicacion_mail.String).Start()

				}).Flags(g.SelectableFlagsSpanAllColumns),

			// Descripción
			g.Label(avance.Descripcion).Wrapped(true),
		).MinHeight(30)
		filas = append(filas, tableRow)
	}

	return
}

func (v detallePendiente) tablaDeAdjuntos() (filas []*g.TableRowWidget) {

	for _, adjunto := range v.adjuntos {

		tableRow := g.TableRow(

			// Descripción
			g.Selectable(adjunto.Descripcion).
				OnDClick(func() {

					if adjunto.Ubicacion_archivo.String == "" {
						return
					}

					_, err := os.Stat(adjunto.Ubicacion_archivo.String)
					if err != nil {

						g.Msgbox("Error", err.Error())

						fmt.Println(err)
						return
					}
					exec.Command("rundll32", "url.dll,FileProtocolHandler", adjunto.Ubicacion_archivo.String).Start()

				}).Flags(g.SelectableFlagsSpanAllColumns),
		).MinHeight(30)

		filas = append(filas, tableRow)
	}

	return
}
