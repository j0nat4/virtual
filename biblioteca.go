package main

import (
	"fmt"
	"time"
)

// Definimos las estructuras para usuario, libro y préstamo
type Usuario struct {
	ID       int
	Nombre   string
	Apellido string
}

type Libro struct {
	ID     int
	Titulo string
	Autor  string
}

type Prestamo struct {
	ID              int
	Usuario         Usuario
	Libro           Libro
	FechaPrestamo   time.Time
	FechaDevolucion time.Time
}

var usuarios []Usuario
var libros []Libro
var prestamos []Prestamo

// Función para agregar un usuario
func agregarUsuario(id int, nombre, apellido string) {
	usuario := Usuario{ID: id, Nombre: nombre, Apellido: apellido}
	usuarios = append(usuarios, usuario)
	fmt.Println("Usuario agregado:", usuario)
}

// Función para agregar un libro
func agregarLibro(id int, titulo, autor string) {
	libro := Libro{ID: id, Titulo: titulo, Autor: autor}
	libros = append(libros, libro)
	fmt.Println("Libro agregado:", libro)
}

// Función para registrar un préstamo
func registrarPrestamo(id int, usuarioID int, libroID int, fechaPrestamo time.Time, fechaDevolucion time.Time) {
	var usuario Usuario
	var libro Libro

	// Buscar el usuario y libro en sus respectivas listas
	for _, u := range usuarios {
		if u.ID == usuarioID {
			usuario = u
			break
		}
	}

	for _, l := range libros {
		if l.ID == libroID {
			libro = l
			break
		}
	}

	// Crear el préstamo
	prestamo := Prestamo{
		ID:              id,
		Usuario:         usuario,
		Libro:           libro,
		FechaPrestamo:   fechaPrestamo,
		FechaDevolucion: fechaDevolucion,
	}
	prestamos = append(prestamos, prestamo)
	fmt.Println("Préstamo registrado:", prestamo)
}

func mostrarMenu() {
	fmt.Println("\n--- Menú de Biblioteca ---")
	fmt.Println("1. Agregar usuario")
	fmt.Println("2. Agregar libro")
	fmt.Println("3. Registrar préstamo")
	fmt.Println("4. buscar libro")
	fmt.Println("5. Ver préstamos")
	fmt.Println("6. Salir")
}

func main() {
	var opcion int

	for {
		mostrarMenu()
		fmt.Print("Selecciona una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var id int
			var nombre, apellido string
			fmt.Print("ID del usuario: ")
			fmt.Scan(&id)
			fmt.Print("Nombre: ")
			fmt.Scan(&nombre)
			fmt.Print("Apellido: ")
			fmt.Scan(&apellido)
			agregarUsuario(id, nombre, apellido)
		case 2:
			var id int
			var titulo, autor string
			fmt.Print("ID del libro: ")
			fmt.Scan(&id)
			fmt.Print("Título: ")
			fmt.Scan(&titulo)
			fmt.Print("Autor: ")
			fmt.Scan(&autor)
			agregarLibro(id, titulo, autor)
		case 3:
			var id, usuarioID, libroID int
			var fechaPrestamo, fechaDevolucion string
			fmt.Print("ID del préstamo: ")
			fmt.Scan(&id)
			fmt.Print("ID del usuario: ")
			fmt.Scan(&usuarioID)
			fmt.Print("ID del libro: ")
			fmt.Scan(&libroID)
			fmt.Print("Fecha de préstamo (YYYY-MM-DD): ")
			fmt.Scan(&fechaPrestamo)
			fmt.Print("Fecha de devolución (YYYY-MM-DD): ")
			fmt.Scan(&fechaDevolucion)

			// Convertir las fechas de string a time.Time
			fechaP, _ := time.Parse("2006-01-02", fechaPrestamo)
			fechaD, _ := time.Parse("2006-01-02", fechaDevolucion)

			registrarPrestamo(id, usuarioID, libroID, fechaP, fechaD)

		case 4:
			fmt.Println("\n--- Lista de préstamos ---")
			for _, p := range prestamos {
				fmt.Printf("Préstamo ID: %d\nUsuario: %s %s\nLibro: %s\nFecha de préstamo: %s\nFecha de devolución: %s\n\n",
					p.ID, p.Usuario.Nombre, p.Usuario.Apellido, p.Libro.Titulo,
					p.FechaPrestamo.Format("2006-01-02"), p.FechaDevolucion.Format("2006-01-02"))
			}

		case 6:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida. Intenta nuevamente.")

		}

	}
}
