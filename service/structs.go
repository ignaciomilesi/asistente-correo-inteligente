package service

import (
	"database/sql"
)

type PendienteResumen struct {
	ID                int    `db:"id"`
	Titulo            string `db:"Titulo"`
	Descripcion       string `db:"Descripcion"`
	Estado            string `db:"Estado"`
	FechaUltimoAvance sql.NullTime
}

type PendienteCompleto struct {
	ID               int            `db:"id"`
	Titulo           string         `db:"Titulo"`
	Descripcion      string         `db:"Descripcion"`
	Estado           string         `db:"Estado"`
	Finalizada       bool           `db:"Finalizada"`
	Fecha_iniciada   sql.NullTime   `db:"Fecha_iniciada"`
	Fecha_finalizada sql.NullTime   `db:"Fecha_finalizada"`
	Cierre           sql.NullString `db:"Cierre"`
	Asignado         sql.NullInt32  `db:"asignado"`
}

type Avance struct {
	Id             int            `db:"id"`
	Fecha_Avance   sql.NullTime   `db:"Fecha_Avance"`
	Descripcion    string         `db:"Descripcion"`
	Ubicacion_mail sql.NullString `db:"ubicacion_mail"`
	Pendiente_id   int            `db:"Pendientes_id"`
}

type Adjunto struct {
	Id                int            `db:"id"`
	Descripcion       string         `db:"Descripcion"`
	Ubicacion_archivo sql.NullString `db:"ubicacion_archivo"`
	Pendiente_id      int            `db:"Pendientes_id"`
}
