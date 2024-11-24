/*
@Autor: Jhonatan velez
@Fecha: 24/11/2024
@Descripcion: Sistema de gestion de libros electronicos
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura para representar un libro
type Libro struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Año    string `json:"año"`
}

// crear metodos
func guardarLibro(libros []Libro) error {
	file, err := os.Create("libros.json")
	if err != nil {
		return err
	}

	// palabra de go
	defer file.Close()

	//guardar en la memorria
	enconder := json.NewEncoder(file)
	//recuperar errores
	err = enconder.Encode(libros)
	if err != nil {
		return err
	}
	return nil
}

//crear leer contctos

func leerLibrosArchivo(libros *[]Libro) error {
	file, err := os.Open("libros.json")
	if err != nil {
		return err
	}
	// palabra
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(libros)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// crear
	var libros []Libro

	// cargar

	err := leerLibrosArchivo(&libros)
	if err != nil {
		fmt.Println("Error al cargar el libro")
	}
	//crear una instancia

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("+++++++ GETOR DE LIBROS +++++++\n",
			"1. Crear nuevo libro\n",
			"2. Lista de libro\n",
			"3. Buscar libro\n",
			"4. Salir\n",
			"Ingresa una opcion:")
		var opcion int
		_, err = fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error al carga la entrada")
		}
		// el switch
		switch opcion {
		case 1:
			var c Libro
			fmt.Println("Titulos: ")
			c.Titulo, _ = reader.ReadString('\n')
			fmt.Println("Autor: ")
			c.Autor, _ = reader.ReadString('\n')
			fmt.Println("Año: ")
			c.Año, _ = reader.ReadString('\n')
			//agregar
			libros = append(libros, c)

			if err := guardarLibro(libros); err != nil {
				fmt.Println("Error al guardar el libro", err)
			}
		case 2:
			fmt.Println("============")
			fmt.Println("Lista de libros")
			fmt.Println("============")
			for index, libro := range libros {
				fmt.Printf("%d. Año: %s Autor: %s Titulo: %s\n",
					index+1, libro.Año, libro.Autor, libro.Titulo)

			}
			fmt.Println("============")
		case 3:
			fmt.Println("============")
			fmt.Print("\n Ingresa el titulo del libro a buscar: ")
			var tituloBuscar string
			tituloBuscar, _ = reader.ReadString('\n')

			encontrado := false
			for _, libro := range libros {
				if libro.Titulo == tituloBuscar {
					fmt.Println("\nTitulo encontrao")
					fmt.Printf("Año: %s Autor: %s Titulo: %s\n",
						libro.Año, libro.Autor, libro.Titulo)
					encontrado = true
					break

				}
			}

			if !encontrado {
				fmt.Println("============")
				fmt.Println("NO SE ENCONTRO EL LIBRO")
				fmt.Println("============")
			}

		case 4:
			fmt.Println("........Gracias por visitar nustra BIBLIOECA........")
			return
		default:
			fmt.Println("++OPCION INCORECTA++")

		}

	}
}
