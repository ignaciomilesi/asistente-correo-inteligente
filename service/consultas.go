package service

// devuelve la lista de todos los pendientes no finalizado, con su ultimo avance
func (s Service) ObtenerListaPendientes() (pendientes []Pendiente, err error) {

	query := `SELECT * FROM Lista_de_Pendientes`

	err = s.db.ConsultaListaConMapeo(&pendientes, query)

	return pendientes, err

}

// devuelve el detalle del pendiente
func (s Service) ObtenerDetallePendiente(idPendiente int) (pendiente PendienteCompleto, err error) {

	query := `SELECT p.*, u.nombre AS asignado
				FROM pendientes AS p
				LEFT JOIN usuarios AS u
				ON p.asignado = u.id
				WHERE p.id=?;`

	err = s.db.ConsultaConMapeo(&pendiente, query, idPendiente)

	return pendiente, nil
}

// devuelve la lista de avances del pendiente, ordenados decendientes
func (s Service) ObtenerListaAvance(idPendiente int) (avances []Avance, err error) {

	query := `SELECT * FROM avances
	WHERE Pendientes_id = ?
	ORDER BY Fecha_Avance ASC;`

	err = s.db.ConsultaListaConMapeo(&avances, query, idPendiente)

	return avances, err
}

// devuelve la lista de adjuntos del pendiente, ordenados decendientes
func (s Service) ObtenerListaAdjunto(idPendiente int) (adjunto []Adjunto, err error) {

	query := `SELECT * FROM adjuntos
	WHERE Pendientes_id = ?;`

	err = s.db.ConsultaListaConMapeo(&adjunto, query, idPendiente)

	return adjunto, err
}
