package widgets

import (
	"database/sql"
	"testing"
	"time"

	g "github.com/AllenDang/giu"

	"go-ollama/service"
)

type mockserviceDetalle struct {
	MockPedienteCompleto service.PendienteCompleto
	MockListaAdjunto     []service.Adjunto
	MockListaAvance      []service.Avance
	MockListaUsuarios    []string
}

func (msd mockserviceDetalle) ObtenerDetallePendiente(id int) (service.PendienteCompleto, error) {
	return msd.MockPedienteCompleto, nil
}

func (msd mockserviceDetalle) ObtenerListaAdjunto(id int) ([]service.Adjunto, error) {
	return msd.MockListaAdjunto, nil
}

func (msd mockserviceDetalle) ObtenerListaAvance(id int) ([]service.Avance, error) {
	return msd.MockListaAvance, nil
}

func (msd mockserviceDetalle) ObtenerListaUsuarios() ([]string, error) {
	return msd.MockListaUsuarios, nil
}

func TestEsquemaDetallePendiente(t *testing.T) {

	msd := mockserviceDetalle{
		MockPedienteCompleto: service.PendienteCompleto{
			ID:     42,
			Titulo: "Titulo de Prueba",
			Descripcion: `prueba, Descripción de pruebaDescripción de prueba, Descripción de prueba, Descripción de prueba, descripción de prueba, Descripción de prueba, Descripción de prueba, descripción de prueba, Descripción de prueba, Descripción de prueba, descripción de prueba,
	   Descripción de prueba, Descripción de prueba, Descripción de prueba

	   !sap:1000004080
	   !s:1000004080| - texto a la derecha

	   Descripción de prueba, Descripción de prueba,
	   !-:Bullet 1: prueba 1
	   !l:www.google.com|etiqueta en link
	   !bullet:Bullet 2, prueba 1
	   !link:www.google.com

	   !malComando:www.google.com
	   Descripción de prueba, Descripción de prueba, descripción de prueba,
	   Descripción de pruebaDescripción de prueba, Descripción de prueba, Descripción de prueba, Descripción de prueba
	   `,
			Estado:     "Estado de Prueba",
			Finalizada: false,
			Fecha_iniciada: sql.NullTime{
				Time: time.Date(2025, time.Month(11), 15, 0, 0, 0, 0, time.UTC),
			},
			Fecha_finalizada: sql.NullTime{},
			Cierre: sql.NullString{
				String: "Cierre de Prueba, Cierre de Prueba, Cierre de Prueba"},
			Asignado: sql.NullInt32{
				Int32: 5},
		},
		MockListaAvance: []service.Avance{
			service.Avance{
				Id:             1,
				Fecha_Avance:   sql.NullTime{Time: time.Date(2025, time.Month(11), 20, 0, 0, 0, 0, time.UTC)},
				Descripcion:    "Avance de prueba 1",
				Ubicacion_mail: sql.NullString{String: "D:\\milesi\\Archivos\\01 - Personal\\Proyectos\\GO-ollama\\mailTest\\prueba.msg"},
			},
			service.Avance{
				Id:             8,
				Fecha_Avance:   sql.NullTime{Time: time.Date(2025, time.Month(11), 15, 0, 0, 0, 0, time.UTC)},
				Descripcion:    "Registro de la activad, Avance de prueba 1, ",
				Ubicacion_mail: sql.NullString{String: "hipervinculo de error"},
			},
			service.Avance{
				Id:           10,
				Fecha_Avance: sql.NullTime{Time: time.Date(2025, time.Month(11), 3, 0, 0, 0, 0, time.UTC)},
				Descripcion:  "Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, Descripción larga, ",
			},
		},
		MockListaAdjunto: []service.Adjunto{
			service.Adjunto{
				Id:                2,
				Descripcion:       "Esquemas",
				Ubicacion_archivo: sql.NullString{String: "D:\\milesi\\Archivos\\01 - Personal\\Proyectos\\GO-ollama\\mailTest\\foto.jpeg"},
			},
			service.Adjunto{
				Id:                7,
				Descripcion:       "lista en PDF",
				Ubicacion_archivo: sql.NullString{String: "D:\\milesi\\Archivos\\01 - Personal\\Proyectos\\GO-ollama\\mailTest\\informe.pdf"},
			},
		},
		MockListaUsuarios: []string{
			"Mateo García",
			"Valentina López",
			"Lucas Fernández",
			"Isabella Martínez",
			"Thiago Rodríguez",
			"Sofía Torres",
		},
	}

	vlp := NuevoDetallePendiente(msd, 4)

	t.Run("Prueba de ventana", func(t *testing.T) {

		gui := g.NewMasterWindow("Buscar expediente", 1000, 600, g.MasterWindowFlagsMaximized)

		gui.Run(func() {

			swmb := g.SingleWindowWithMenuBar()
			swmb.Layout(
				vlp,
			)
		})
	})

}
