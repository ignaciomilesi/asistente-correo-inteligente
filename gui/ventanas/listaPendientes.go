package ventanas

import (
	"fmt"
	"go-ollama/service"
	"strconv"

	g "github.com/AllenDang/giu"
)

type serviceInterface interface {
	ObtenerListaPendientes() ([]service.Pendiente, error)
}

type VentanaListaPendiente struct {
	Service       serviceInterface
	ventanaActiva bool
	pendientes    []service.Pendiente
}

func (v *VentanaListaPendiente) EsVentanaActiva() bool {
	return true
}

func (v *VentanaListaPendiente) Actualizar() {

	var err error
	v.pendientes, err = v.Service.ObtenerListaPendientes()

	if err != nil {
		fmt.Printf("No se pudo actualizar la lista de pendientes.\n%v\n", err)
	}
}

func (v *VentanaListaPendiente) Esquema() {

	g.SingleWindowWithMenuBar().Layout(
		/*
			g.MenuBar().Layout(
				g.MenuItem("Configuración").OnClick(vp.MostrarConfig),
			),
		*/
		g.Align(g.AlignRight).To(
			g.Button("Cargar nuevo pendiente"),
		),

		g.Dummy(5, 5),

		g.Table().Columns(
			g.TableColumn("ID").Flags(g.TableColumnFlagsNoResize|g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(20),
			//g.TableColumn("-").Flags(g.TableColumnFlagsNoResize|g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(30),
			g.TableColumn("Titulo").Flags(g.TableColumnFlagsNoResize|g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(300),
			g.TableColumn("Estado").Flags(g.TableColumnFlagsNoResize),
			g.TableColumn("Ultimo Avance").Flags(g.TableColumnFlagsNoResize|g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(100),
		).Rows(
			v.obtenerFilas()...,
		).Flags(g.TableFlagsRowBg|g.TableFlagsBorders),
	)

}

func (v VentanaListaPendiente) obtenerFilas() (filas []*g.TableRowWidget) {

	for _, pendiente := range v.pendientes {

		tableRow := g.TableRow(
			// ID
			g.Selectable(strconv.Itoa(pendiente.ID)).Flags(g.SelectableFlagsSpanAllColumns).OnDClick(
				func() {
					fmt.Println("Ver detalle")
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
