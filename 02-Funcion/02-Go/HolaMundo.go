/*
	Programa que muestra el mensaje "Hola Mundo" en la pantalla
	Autor: Julián Esteban Gutiérrez Posada
	Fecha: Febrero 2025
	Licencia: GNU GLP v3
*/

package main

import "fmt"

func main() {
	mostrarMensaje("Hola Mundo")
}

func mostrarMensaje(mensaje string) {
	fmt.Println(mensaje)
}
