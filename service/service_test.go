package service

import (
	"database/sql"
	"fmt"
	"go-ollama/pkg/database"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestNuevoRegistroPediente(t *testing.T) {

	pendienteParaCargar := PendienteCompleto{
		Titulo:         fmt.Sprint("Pendiente Test ", time.Now().Format("02-01-2006")),
		Descripcion:    fmt.Sprint("Descripción Test ", time.Now().Format("02-01-2006")),
		Estado:         fmt.Sprint("Estado Test ", time.Now().Format("02-01-2006")),
		Fecha_iniciada: sql.NullTime{Time: time.Now()},
		Asignado:       sql.NullInt32{Int32: 4},
	}

	if err := godotenv.Load("..\\.env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	gdb := &database.GestorDb{}
	gdb.Conectar("test")

	ser := New(gdb)

	idCreado, err := ser.NuevoRegistroPediente(pendienteParaCargar)
	if err != nil {
		log.Fatal("Error al crear el registro. ", err)
	}
	fmt.Printf("Registro creado con el id: %d\n", idCreado)

	pendienteGenerado, err := ser.ObtenerDetallePendiente(int(idCreado))
	if err != nil {
		fmt.Println("Error obtener el pendiente", err)
	}

	fmt.Println(pendienteGenerado)
}

func TestNuevoRegistroAvance(t *testing.T) {

	tests := []struct {
		name       string
		avanceTest Avance
	}{
		{
			name: "Todo OK",
			avanceTest: Avance{
				Fecha_Avance:   sql.NullTime{Time: time.Now(), Valid: true},
				Descripcion:    "Descripción de prueba. Todo OK",
				Ubicacion_mail: sql.NullString{String: "ubicación/cualquiera", Valid: true},
				Pendiente_id:   7,
			},
		},
		{
			name: "Descripción en blanco",
			avanceTest: Avance{
				Fecha_Avance:   sql.NullTime{Time: time.Now(), Valid: true},
				Descripcion:    "",
				Ubicacion_mail: sql.NullString{String: "ubicación/cualquiera", Valid: true},
				Pendiente_id:   7,
			},
		},
		{
			name: "Sin mail",
			avanceTest: Avance{
				Fecha_Avance: sql.NullTime{Time: time.Now(), Valid: true},
				Descripcion:  "Descripción de prueba. Sin mail",
				Pendiente_id: 7,
			},
		},
		{
			name: "Sin id pendiente",
			avanceTest: Avance{
				Fecha_Avance:   sql.NullTime{Time: time.Now(), Valid: true},
				Descripcion:    "Descripción de prueba. Sin id pendiente",
				Ubicacion_mail: sql.NullString{String: "ubicación/cualquiera", Valid: true},
			},
		},
	}

	if err := godotenv.Load("..\\.env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	gdb := &database.GestorDb{}
	gdb.Conectar("test")
	ser := New(gdb)

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			idCreado, err := ser.NuevoRegistroAvence(test.avanceTest)
			if err != nil {
				t.Errorf("Error al crear el registro. %v\n\n", err)
			} else {
				fmt.Printf("Test %s: Registro creado con el id: %d\n", test.name, idCreado)
			}
		})
	}

}
