package models

import (
	"time"
)

// Estructura para representar un préstamo
type Prestamo struct {
	ID        int       `json:"id"`
	UsuarioID int       `json:"usuario_id"`
	LibroID   int       `json:"libro_id"`
	Fecha     time.Time `json:"fecha"`
	Devuelto  bool      `json:"devuelto"`
}
