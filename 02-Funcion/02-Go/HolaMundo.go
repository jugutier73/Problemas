/*
	Programa que muestra el mensaje "Hola Mundo" en la pantalla
	Autor: Julián Esteban Gutiérrez Posada
	Fecha: Enero 2025
	Licencia: GNU GPL v3
*/

package main

import "fmt"

func main() {
	mostrarMensaje("Hola Mundo\n")
}

func mostrarMensaje(mensaje string) {
	fmt.Print(mensaje)
}
