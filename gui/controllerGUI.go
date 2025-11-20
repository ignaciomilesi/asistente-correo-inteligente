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
	Actualizar()
}

type serviceInterface interface {
	ObtenerListaPendientes() ([]service.Pendiente, error)
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

	cg.ventanaListaPendiente = &ventanas.VentanaListaPendiente{
		Service: cg.service,
	}

	gui := g.NewMasterWindow("Buscar expediente", 1000, 600, g.MasterWindowFlagsTransparent)
	gui.SetPos(100, 100)

	// seteo el cierre de la base de datos al cerrar la ventana
	gui.SetCloseCallback(func() bool {
		fmt.Println("Cerrando la base de datos")
		cg.service.Cerrar()
		return true
	})

	cg.actualizarVentanas()

	gui.Run(cg.mostrarVentanas)
}

func (cg controlerGUI) actualizarVentanas() {
	if cg.ventanaListaPendiente.EsVentanaActiva() {
		cg.ventanaListaPendiente.Actualizar()
	}
}

func (cg controlerGUI) mostrarVentanas() {

	if cg.ventanaListaPendiente.EsVentanaActiva() {
		cg.ventanaListaPendiente.Esquema()
	}
}

/*
func (cg *ControlerGUI) setVentanaPrincipal() {

	cg.vp.Activa = true
	cg.vp.Visible = true

	cg.vp.Control = &cg.control

	cg.vp.MostrarConfig = func() {
		cg.vp.Activa = false
		cg.vc.Visible = true
	}

}

func (cg *ControlerGUI) setVentanaConfig() {
	cg.vc.Activa = true
	cg.vc.Visible = false

	cg.vc.Control = &cg.control

	cg.vc.CerrarVentana = func() {

		cg.vp.Activa = true
		cg.vc.Visible = false
	}

}*/
