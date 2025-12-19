package ventanas

import (
	"database/sql"
	"strconv"
	"testing"
	"time"

	g "github.com/AllenDang/giu"

	"go-ollama/service"
)

type mockservice struct {
	MockPediente         []service.Pendiente
	MockPedienteCompleto service.PendienteCompleto
	MockListaAdjunto     []service.Adjunto
	MockListaAvance      []service.Avance
	MockListaUsuarios    []string
}

func (self mockservice) ObtenerListaPendientes() ([]service.Pendiente, error) {
	return self.MockPediente, nil
}
func (self mockservice) ObtenerDetallePendiente(id int) (service.PendienteCompleto, error) {
	return self.MockPedienteCompleto, nil
}

func (self mockservice) ObtenerListaAdjunto(id int) ([]service.Adjunto, error) {
	return self.MockListaAdjunto, nil
}

func (self mockservice) ObtenerListaAvance(id int) ([]service.Avance, error) {
	return self.MockListaAvance, nil
}

func (self mockservice) ObtenerListaUsuarios() ([]string, error) {
	return self.MockListaUsuarios, nil
}

func TestListaPendiente(t *testing.T) {

	var pedientes []service.Pendiente

	for i := range 10 {
		newPendiente := service.Pendiente{
			ID:          i,
			Titulo:      "Titulo" + strconv.Itoa(i),
			Descripcion: "Descripción" + strconv.Itoa(i),
			Estado:      "Estado" + strconv.Itoa(i),
			FechaUltimoAvance: sql.NullTime{
				Time: time.Date(2025, time.Month(i+1), i+3, 0, 0, 0, 0, time.UTC),
			},
		}

		pedientes = append(pedientes, newPendiente)

	}

	ms := mockservice{
		MockPediente: pedientes,
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
			Cierre:           "Cierre de Prueba, Cierre de Prueba, Cierre de Prueba",
			Asignado:         5,
		},
		MockListaAvance: []service.Avance{
			service.Avance{
				Id:             1,
				Fecha_Avance:   sql.NullTime{Time: time.Date(2025, time.Month(11), 20, 0, 0, 0, 0, time.UTC)},
				Descripcion:    "Avance de prueba 1",
				Ubicacion_mail: "D:\\milesi\\Archivos\\01 - Personal\\Proyectos\\GO-ollama\\mailTest\\prueba.msg",
			},
			service.Avance{
				Id:             8,
				Fecha_Avance:   sql.NullTime{Time: time.Date(2025, time.Month(11), 15, 0, 0, 0, 0, time.UTC)},
				Descripcion:    "Registro de la activad, Avance de prueba 1, ",
				Ubicacion_mail: "hipervinculo de error",
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
				Ubicacion_archivo: "D:\\milesi\\Archivos\\01 - Personal\\Proyectos\\GO-ollama\\mailTest\\foto.jpeg",
			},
			service.Adjunto{
				Id:                7,
				Descripcion:       "lista en PDF",
				Ubicacion_archivo: "D:\\milesi\\Archivos\\01 - Personal\\Proyectos\\GO-ollama\\mailTest\\informe.pdf",
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

	vlp := NuevaVentanaListaPendiente(ms)

	gui := g.NewMasterWindow("Buscar expediente", 1000, 600, g.MasterWindowFlagsMaximized)

	gui.SetDropCallback(func(pathFiles []string) {

		vlp.DropFileCallback(*gui, pathFiles)

	})

	gui.Run(vlp.Esquema)

}
