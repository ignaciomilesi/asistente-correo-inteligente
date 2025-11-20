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
	MockPediente []service.Pendiente
}

func (ms mockservice) ObtenerListaPendientes() ([]service.Pendiente, error) {
	return ms.MockPediente, nil
}

func TestEsquema(t *testing.T) {

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
	}

	vlp := VentanaListaPendiente{
		Service: ms,
	}

	vlp.Actualizar()

	t.Run("Prueba de ventana", func(t *testing.T) {

		gui := g.NewMasterWindow("Buscar expediente", 1000, 600, g.MasterWindowFlagsTransparent)
		gui.SetPos(100, 100)

		gui.Run(vlp.Esquema)
	})

}
