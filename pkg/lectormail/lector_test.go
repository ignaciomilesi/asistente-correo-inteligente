package lectormail

import (
	"testing"
)

func Test_extraerFecha(t *testing.T) {

	tests := []struct {
		name         string
		cabeceraTest string
		want         string
		wantErr      bool
	}{
		{
			name:         "Prueba cabecera ok",
			cabeceraTest: "dummy;Mon, 12 Jun 2023 15:04:05",
			want:         "12/06/2023",
			wantErr:      false,
		},
		{
			name:         "Prueba cabecera con - al final",
			cabeceraTest: "dummy;Mon, 22 Dec 2016 15:04:05 -",
			want:         "22/12/2016",
			wantErr:      false,
		},

		{
			name:         "Prueba cabecera con doble ;",
			cabeceraTest: "dummy;dummy;Mon, 02 Jan 2006 15:04:05",
			want:         "02/01/2006",
			wantErr:      false,
		},
		{
			name:         "Prueba cabecera sin ;",
			cabeceraTest: "Mon, 02 Jan 2006 15:04:05",
			want:         "01/01/0001",
			wantErr:      true,
		},
		{
			name:         "Prueba cabecera sin mon",
			cabeceraTest: "Mon, 02 Jan 2006 15:04:05",
			want:         "01/01/0001",
			wantErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			resultado, err := extraerFecha(test.cabeceraTest)

			if (err != nil) != test.wantErr {
				t.Errorf("Error! -----> se esperaba error %t, devolvió %t, \nmsgError: %v",
					test.wantErr, (err != nil), err)
			}

			if resultado.Format("02/01/2006") != test.want {
				t.Errorf("Error! -----> se esperaba %s, devolvió %s",
					test.want, resultado)
			}
		})
	}

}
