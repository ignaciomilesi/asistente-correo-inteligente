package service

import (
	"database/sql"
)

type dbInterface interface {
	Conectar(string)
	Salir()
	ConsultaSimple(string, ...any) (*sql.Rows, error)
	ConsultaConMapeo(interface{}, string, ...any) error
	ConsultaListaConMapeo(interface{}, string, ...any) error
}

type service struct {
	db dbInterface
}

func New(nuevoGestorBaseDatos dbInterface) *service {
	return &service{db: nuevoGestorBaseDatos}
}

func (s *service) Cerrar() {
	s.db.Salir()
}
