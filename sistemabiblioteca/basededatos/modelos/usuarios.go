// models/usuarios.go
package models

import (
	"database/sql"
	"fmt"
)

// Usuario representa la estructura de un usuario
type Usuario struct {
	UsuarioID int    `json:"usuario_id"`
	Nombre    string `json:"nombre"`
	Email     string `json:"email"`
	Activo    bool   `json:"activo"`
}

// ObtenerUsuarios obtiene todos los usuarios de la base de datos
func ObtenerUsuarios(db *sql.DB) ([]Usuario, error) {
	// Ejecutar la consulta para obtener los usuarios
	rows, err := db.Query("SELECT UsuarioID, Nombre, Email, Activo FROM Usuarios")
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	// Crear una lista de usuarios
	var usuarios []Usuario

	// Iterar sobre los resultados y agregarlos a la lista
	for rows.Next() {
		var usuario Usuario
		if err := rows.Scan(&usuario.UsuarioID, &usuario.Nombre, &usuario.Email, &usuario.Activo); err != nil {
			return nil, fmt.Errorf("error al leer los resultados: %v", err)
		}
		usuarios = append(usuarios, usuario)
	}

	// Verificar si hubo errores al iterar sobre los resultados
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre los resultados: %v", err)
	}

	return usuarios, nil
}
