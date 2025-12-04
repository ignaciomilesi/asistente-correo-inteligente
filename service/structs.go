package service

import (
	"database/sql"
)

type Pendiente struct {
	ID                int    `db:"id"`
	Titulo            string `db:"Titulo"`
	Descripcion       string `db:"Descripcion"`
	Estado            string `db:"Estado"`
	FechaUltimoAvance sql.NullTime
}

type PendienteCompleto struct {
	ID               int          `db:"id"`
	Titulo           string       `db:"Titulo"`
	Descripcion      string       `db:"Descripcion"`
	Estado           string       `db:"Estado"`
	Finalizada       bool         `db:"Finalizada"`
	Fecha_iniciada   sql.NullTime `db:"Fecha_iniciada"`
	Fecha_finalizada sql.NullTime `db:"Fecha_finalizada"`
	Cierre           string       `db:"Cierre"`
	Asignado         int32        `db:"asignado"`
}

type Avance struct {
	Id             int          `db:"id"`
	Fecha_Avance   sql.NullTime `db:"Fecha_Avance"`
	Descripcion    string       `db:"Descripcion"`
	Ubicacion_mail string       `db:"ubicacion_mail"`
}

type Adjunto struct {
	Id                int    `db:"id"`
	Descripcion       string `db:"Descripcion"`
	Ubicacion_archivo string `db:"ubicacion_archivo"`
}
