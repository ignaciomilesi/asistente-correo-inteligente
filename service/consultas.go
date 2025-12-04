package service

// devuelve la lista de todos los pendientes no finalizado, con su ultimo avance
func (s service) ObtenerListaPendientes() (pendientes []Pendiente, err error) {

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

// devuelve la lista de avances del pendiente, ordenados decendientes
func (s service) ObtenerListaAvance(idPendiente int) (avances []Avance, err error) {

	query := `SELECT * FROM avances
	WHERE Pendientes_id = ?
	ORDER BY Fecha_Avance ASC;`

	err = s.db.ConsultaListaConMapeo(&avances, query, idPendiente)

	return avances, err
}

// devuelve la lista de adjuntos del pendiente, ordenados decendientes
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
