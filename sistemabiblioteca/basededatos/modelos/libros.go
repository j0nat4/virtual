// models/libros.go
package models

import (
	"fmt"
	"sistema-biblioteca/database"
)

// Libro representa la estructura de un libro
type Libro struct {
	LibroID   int              `json:"libro_id"`
	Titulo    string           `json:"titulo"`
	Autor     string           `json:"autor"`
	Categoria string           `json:"categoria"`
	Año       int              `json:"año"`
	Prestado  bool             `json:"prestado"`
	Precio    sql.NullFloat644 `json:"precio"`
}

// ObtenerLibros obtiene todos los libros de la base de datos
func ObtenerLibros() ([]Libro, error) {
	// Establecer la conexión a la base de datos
	db, err := database.Conectar()
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Ejecutar la consulta para obtener los libros
	rows, err := db.Query("SELECT LibroID, Titulo, Autor, Categoria, Año, Prestado, Precio FROM Libros")
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	// Crear una lista de libros
	var libros []Libro

	// Iterar sobre los resultados y agregarlos a la lista
	for rows.Next() {
		var libro Libro
		if err := rows.Scan(&libro.LibroID, &libro.Titulo, &libro.Autor, &libro.Categoria, &libro.Año, &libro.Prestado, &libro.Precio); err != nil {
			return nil, fmt.Errorf("error al leer los resultados: %v", err)
		}
		libros = append(libros, libro)
	}

	// Verificar si hubo errores al iterar sobre los resultados
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre los resultados: %v", err)
	}

	return libros, nil
}
