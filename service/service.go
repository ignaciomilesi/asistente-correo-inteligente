package service

import (
	"database/sql"
	"fmt"
)

type dbInterface interface {
	Conectar(string)
	Salir()
	ConsultaSimple(string, ...any) (*sql.Rows, error)
	ConsultaConMapeo(interface{}, string, ...any) error
	ConsultaListaConMapeo(interface{}, string, ...any) error

	NuevoRegistro(string, ...any) (int64, error)
	ActualizarRegistro(string, ...any) (int64, error)
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

// devuelve la lista de todos los pendientes no finalizado, con su ultimo avance
func (s service) ObtenerListaPendientes() (pendientes []PendienteResumen, err error) {

	query := `SELECT * FROM Lista_de_Pendientes`

	err = s.db.ConsultaListaConMapeo(&pendientes, query)

	return pendientes, err

}

// devuelve el detalle del pendiente
func (s service) ObtenerDetallePendiente(idPendiente int) (pendiente PendienteCompleto, err error) {

	query := `SELECT *
				FROM pendientes AS p
				WHERE p.id=?;`

	err = s.db.ConsultaConMapeo(&pendiente, query, idPendiente)

	return pendiente, nil
}

// devuelve la lista de avances del pendiente, ordenados descendientes
func (s service) ObtenerListaAvance(idPendiente int) (avances []Avance, err error) {

	query := `SELECT * FROM avances
	WHERE Pendientes_id = ?
	ORDER BY Fecha_Avance ASC;`

	err = s.db.ConsultaListaConMapeo(&avances, query, idPendiente)

	return avances, err
}

// devuelve la lista de adjuntos del pendiente, ordenados descendientes
func (s service) ObtenerListaAdjunto(idPendiente int) (adjunto []Adjunto, err error) {

	query := `SELECT * FROM adjuntos
	WHERE Pendientes_id = ?;`

	err = s.db.ConsultaListaConMapeo(&adjunto, query, idPendiente)

	return adjunto, err
}

// devuelve la lista de solo los nombres de los usuarios, ordenados por ID
func (s service) ObtenerListaUsuarios() (lista []string, err error) {

	query := `SELECT nombre FROM usuarios 
				ORDER BY id ASC;`

	rows, err := s.db.ConsultaSimple(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var nombre string
		if err := rows.Scan(&nombre); err != nil {
			return nil, err
		}
		lista = append(lista, nombre)
	}

	return lista, err
}

// devuelve el id del registro generado
func (s service) NuevoRegistroPediente(nuevoPendiente PendienteCompleto) (idGenerado int64, err error) {

	if nuevoPendiente.Titulo == "" {
		return 0, fmt.Errorf("El campo título no puede estar en blanco")
	}

	if nuevoPendiente.Descripcion == "" {
		return 0, fmt.Errorf("El campo Descripción no puede estar en blanco")
	}

	if nuevoPendiente.Estado == "" {
		return 0, fmt.Errorf("El campo Estado no puede estar en blanco")
	}
	if !nuevoPendiente.Fecha_iniciada.Valid {
		return 0, fmt.Errorf("La fecha no es una fecha valida")
	}

	query := `INSERT INTO Pendientes 
	(Titulo, Descripcion, Estado, Fecha_iniciada, asignado)
		VALUES (?, ?, ?, ?, ?)`

	return s.db.NuevoRegistro(query, nuevoPendiente.Titulo, nuevoPendiente.Descripcion, nuevoPendiente.Estado, nuevoPendiente.Fecha_iniciada, nuevoPendiente.Asignado)

}

// devuelve el id del registro generado
func (s service) NuevoRegistroAvence(nuevoAvance Avance) (idGenerado int64, err error) {

	if !nuevoAvance.Fecha_Avance.Valid {
		return 0, fmt.Errorf("La fecha no es una fecha valida")
	}
	if nuevoAvance.Descripcion == "" {
		return 0, fmt.Errorf("El campo Descripción no puede estar en blanco")
	}
	if nuevoAvance.Pendiente_id == 0 {
		return 0, fmt.Errorf("El campo Pendiente_id debe ser distinto a cero")
	}
	query := `INSERT INTO Avances 
	(Fecha_Avance, Descripcion, ubicacion_mail, Pendientes_id)
		VALUES (?, ?, ?, ?)`

	return s.db.NuevoRegistro(query, nuevoAvance.Fecha_Avance, nuevoAvance.Descripcion, nuevoAvance.Ubicacion_mail, nuevoAvance.Pendiente_id)
}

/*
	Id             int          `db:"id"`
	Fecha_Avance   sql.NullTime `db:"Fecha_Avance"`
	Descripcion    string       `db:"Descripcion"`
	Ubicacion_mail string       `db:"ubicacion_mail"`
*/
