/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

package modulo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MostrarMensaje mostrar un mensaje (cadena) en la salida estandar.
func MostrarMensaje(mensaje string) {
	fmt.Print(mensaje)
}

// MostrarError mostrar un mensaje (cadena) en el error estandar.
func MostrarError(mensaje string) {
	fmt.Fprint(os.Stderr, mensaje)
}

// IngresarTexto devuelve el texto ingresado por el usuario como respuesta a una pregunta.
func IngresarTexto(pregunta string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(pregunta)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

// IngresarEntero devuelve el entero ingresado por el usuario como respuesta a una pregunta o cero si es un valor inválido.
func IngresarEntero(pregunta string) int {
	entero, err := strconv.Atoi(IngresarTexto(pregunta))
	if err != nil {
		MostrarError("El valor ingresado es inválido, se asume 0\n\n")
		entero = 0
	}
	return entero
}

// IngresarReal devuelve el real ingresado por el usuario como respuesta a una pregunta o cero si es un valor inválido.
func IngresarReal(pregunta string) float64 {
	real, err := strconv.ParseFloat(IngresarTexto(pregunta), 64)
	if err != nil {
		MostrarError("El valor ingresado es inválido, se asume 0.0\n\n")
		real = 0.0
	}
	return real
}

// IngresarLogico devuelve el booleano ingresado por el usuario como respuesta a una pregunta o falso si es un valor inválido.
func IngresarLogico(pregunta string) bool {
	opcion := strings.ToLower(IngresarTexto(pregunta))

	if opcion != "s" && opcion != "n" {
		MostrarError("No es una opción válida, se asume \"n\"\n\n")
	}

	return opcion == "s"
}
