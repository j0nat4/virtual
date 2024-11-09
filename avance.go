/*
@Autor: Jhonatan velez
@Fecha: 08/11/2024
@Descripcion: Sistema de gestion de libros electronicos
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("bienvenidos a mi bibliotca electronica")

	var libro string

	//libros = "Blanca nieves, lo tres chanchitos, pinocho, movidic, etc"

	fmt.Println("Ingrese su cuento: ")
	fmt.Scanln(&libro)
	fmt.Printf("Emos encontrado su cuento de, %s!\n", libro)

}
