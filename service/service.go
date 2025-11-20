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

type Service struct {
	db dbInterface
}

func New(nuevoGestorBaseDatos dbInterface) *Service {
	return &Service{db: nuevoGestorBaseDatos}
}

func (s *Service) Cerrar() {
	s.db.Salir()
}
