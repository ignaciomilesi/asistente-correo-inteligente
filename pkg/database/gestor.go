package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type GestorDb struct {
	db *sql.DB
}

func (g *GestorDb) Conectar(nombreBaseDatos string) {

	fmt.Print("Conectando a ", nombreBaseDatos, "... ")

	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/%s?parseTime=True&loc=Local",
		os.Getenv("DB_PASS"),
		nombreBaseDatos)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error al conectar:", err)
	}
	g.db = db

	fmt.Println("Conectado!")
}

func (g *GestorDb) Salir() {
	if g.db != nil {
		g.db.Close()
	}
}

// Solo realiza un query. Es necesario realizar un rows.close() posterior
func (g GestorDb) ConsultaSimple(query string, args ...any) (*sql.Rows, error) {
	return g.db.Query(query, args...)
}

// Realiza el query y el mapeo de las columnas al struct pasado (debe ser un puntero).
// Se debe utilizar el tag "db" para marcar las columnas.
// En caso de que la consulta devuelva varias filas, la función solo mapea la primera
func (g GestorDb) ConsultaConMapeo(dest interface{}, query string, args ...any) error {

	// Verifico que sea un puntero
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("dest debe ser un puntero no nulo a un struct")
	}

	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("dest debe ser un puntero a un struct")
	}

	// realizo la consulta
	rows, err := g.db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	// si no obtuve nada devuelvo un error
	if !rows.Next() {
		return sql.ErrNoRows
	}

	return g.mapearStructDesdeRows(dest, rows)

}

// Realiza el query y el mapeo (completo, todas las listas) de las columnas al struct pasado (debe ser un puntero).
// Se debe utilizar el tag "db" para marcar las columnas
func (g GestorDb) ConsultaListaConMapeo(dest interface{}, query string, args ...any) error {

	// Verifico que sea un puntero y que sea un slice
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("dest debe ser un puntero no nulo a un slice")
	}

	sliceValue := v.Elem()
	if sliceValue.Kind() != reflect.Slice {
		return fmt.Errorf("dest debe ser un puntero a un slice")
	}

	elemType := sliceValue.Type().Elem()

	// realizo la consulta
	rows, err := g.db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		// Crear un nuevo elemento del tipo del slice
		elemPtr := reflect.New(elemType)

		if err := g.mapearStructDesdeRows(elemPtr.Interface(), rows); err != nil {
			return err
		}

		// Agregarlo al slice
		sliceValue.Set(reflect.Append(sliceValue, elemPtr.Elem()))
	}

	return rows.Err()
}

// Realiza el mapeo de las columnas al struct pasado (dest).
// Se debe utilizar el tag "db" para marcar las columnas
func (g GestorDb) mapearStructDesdeRows(dest interface{}, rows *sql.Rows) error {
	// Verifico que dest sea puntero a struct
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("dest debe ser un puntero no nulo a un struct")
	}

	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("dest debe ser un puntero a un struct")
	}

	// Obtener columnas de la consulta
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Crear mapa columna → puntero a campo
	fieldMap := map[string]interface{}{}

	t := elem.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag != "" {
			fieldMap[tag] = elem.Field(i).Addr().Interface()
		}
	}

	// Preparar slice para Scan
	scanArgs := make([]interface{}, len(columns))
	for i, col := range columns {
		if ptr, ok := fieldMap[col]; ok {
			scanArgs[i] = ptr
		} else {
			// Ignorar columnas no mapeadas
			var dummy interface{}
			scanArgs[i] = &dummy
		}
	}

	// Escanear campos
	return rows.Scan(scanArgs...)
}

func (g GestorDb) MostrarTablas() {

	rows, err := g.db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Tablas disponibles:")
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		fmt.Println("-", tableName)
	}
}

// Genera un nuevo registro, devuelve el id del registro generado
func (g *GestorDb) NuevoRegistro(query string, args ...any) (int64, error) {

	result, err := g.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// actualiza un registro, devuelve la cantidad de registros modificados
func (g *GestorDb) ActualizarRegistro(query string, args ...any) (int64, error) {

	result, err := g.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
