package gui

import (
	"fmt"

	g "github.com/AllenDang/giu"

	"go-ollama/gui/ventanas"
	"go-ollama/service"
)

type ventanaInterface interface {
	EsVentanaActiva() bool
	Esquema()
	DropFileCallback(g.MasterWindow, []string)
}

type serviceInterface interface {
	ObtenerListaPendientes() ([]service.Pendiente, error)
	ObtenerDetallePendiente(int) (service.PendienteCompleto, error)
	ObtenerListaAdjunto(int) ([]service.Adjunto, error)
	ObtenerListaAvance(int) ([]service.Avance, error)
	ObtenerListaUsuarios() ([]string, error)
	Cerrar()
}

type controlerGUI struct {
	service               serviceInterface
	ventanaListaPendiente ventanaInterface
}

func New(ser serviceInterface) *controlerGUI {

	return &controlerGUI{
		service: ser,
	}
}

func (cg *controlerGUI) Ejecutar() {

	cg.ventanaListaPendiente = ventanas.NuevaVentanaListaPendiente(cg.service)

	gui := g.NewMasterWindow("Buscar expediente", 1000, 600, g.MasterWindowFlagsTransparent)
	gui.SetPos(100, 100)

	// seteo el cierre de la base de datos al cerrar la ventana
	gui.SetCloseCallback(func() bool {
		fmt.Println("Cerrando la base de datos")
		cg.service.Cerrar()
		return true
	})

	gui.SetDropCallback(func(pathFilesDrop []string) {

		if cg.ventanaListaPendiente.EsVentanaActiva() {
			cg.ventanaListaPendiente.DropFileCallback(*gui, pathFilesDrop)
		}
	})

	gui.Run(func() {
		if cg.ventanaListaPendiente.EsVentanaActiva() {
			cg.ventanaListaPendiente.Esquema()
		}
	})
}
