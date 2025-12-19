package ventanas

import (
	"fmt"
	"go-ollama/gui/ventanas/widgets"
	"go-ollama/service"
	"strconv"

	g "github.com/AllenDang/giu"
)

type serviceInterfaceListaPendiente interface {
	ObtenerListaPendientes() ([]service.Pendiente, error)
	ObtenerDetallePendiente(int) (service.PendienteCompleto, error)
	ObtenerListaAdjunto(int) ([]service.Adjunto, error)
	ObtenerListaAvance(int) ([]service.Avance, error)
	ObtenerListaUsuarios() ([]string, error)
}

type VentanaListaPendiente struct {
	service       serviceInterfaceListaPendiente
	ventanaActiva bool
	pendientes    []service.Pendiente
	detalle       g.Widget
	anchoSplit    float32
}

func (self *VentanaListaPendiente) EsVentanaActiva() bool {
	return true
}

func NuevaVentanaListaPendiente(serviceBaseDatos serviceInterfaceListaPendiente) *VentanaListaPendiente {

	p, err := serviceBaseDatos.ObtenerListaPendientes()
	if err != nil {
		fmt.Printf("No se pudo actualizar la lista de pendientes.\n%v\n", err)
	}

	return &VentanaListaPendiente{
		service:    serviceBaseDatos,
		pendientes: p,
	}
}

func (self *VentanaListaPendiente) Esquema() {

	g.SingleWindowWithMenuBar().Layout(

		g.Row(
			g.Child().Size(g.Auto-self.anchoSplit, g.Auto).Flags(g.WindowFlagsHorizontalScrollbar).Layout(
				g.Row(
					g.Style().
						SetFontSize(32).To(
						g.Label("Lista de Pendientes"),
					),

					g.Condition(self.anchoSplit > 0,
						g.Dummy(1, 1),
						g.Align(g.AlignRight).To(
							g.Button("Cargar nuevo pendiente"),
						),
					),
				),
				g.Dummy(3, 3),

				g.Table().Columns(
					g.TableColumn("ID").Flags(g.TableColumnFlagsNoResize|g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(20),
					g.TableColumn("Titulo").Flags(g.TableColumnFlagsNoResize|g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(300),
					g.TableColumn("Estado").Flags(g.TableColumnFlagsNoResize),
					g.TableColumn("Ultimo Avance").Flags(g.TableColumnFlagsNoResize|g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(100),
				).Rows(
					self.obtenerFilas()...,
				).Flags(g.TableFlagsRowBg|g.TableFlagsBorders),
			),

			g.Child().Size(600, g.Auto).Layout(
				g.Align(g.AlignRight).To(
					g.Button("Cerrar Detalle").OnClick(func() {
						self.anchoSplit = 0
					}),
				),
				self.detalle,
			),
		),
	)

}

func (self *VentanaListaPendiente) DropFileCallback(ventanaPrincipal g.MasterWindow, pathFiles []string) {

	if len(pathFiles) > 1 {
		g.Msgbox("Error", "Arrastre un solo archivo")
		return
	}

	posicionVentanaX, _ := ventanaPrincipal.GetPos()
	anchoVentana, _ := ventanaPrincipal.GetSize()
	posMouseX := g.GetMousePos().X

	if (anchoVentana - (posMouseX - posicionVentanaX)) < int(self.anchoSplit) {
		fmt.Println("Cargar nuevo avance")
	} else {
		fmt.Println("Cargar nuevo pendiente")
	}

}

func (self *VentanaListaPendiente) obtenerFilas() (filas []*g.TableRowWidget) {

	for _, pendiente := range self.pendientes {

		tableRow := g.TableRow(
			// ID
			g.Selectable(strconv.Itoa(pendiente.ID)).Flags(g.SelectableFlagsSpanAllColumns).OnDClick(
				func() {
					self.detalle = widgets.NuevoDetallePendiente(self.service, pendiente.ID)
					self.anchoSplit = 600
				}),
			g.ContextMenu().Layout(g.Label(pendiente.Descripcion)),

			// TItulo
			g.Label(pendiente.Titulo),

			// Estado
			g.Label(pendiente.Estado),

			// Fecha Ultimo avance
			g.Align(g.AlignCenter).To(g.Label(pendiente.FechaUltimoAvance.Time.Format("2006-01-02"))),
		)
		filas = append(filas, tableRow)
	}

	return
}
