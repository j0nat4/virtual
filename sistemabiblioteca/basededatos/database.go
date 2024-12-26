// database/database.go
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" // Paquete para SQL Server
)

// Conectar se encarga de establecer la conexión con la base de datos
func Conectar() (*sql.DB, error) {
	// La cadena de conexión a la base de datos (ajusta según tu configuración)
	connStr := "sqlserver://localhost:1433?database=BibliotecaDB&trusted_connection=true"
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
	}

	// Verificar si la conexión es exitosa
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al hacer ping a la base de datos: %v", err)
	}

	return db, nil
}
