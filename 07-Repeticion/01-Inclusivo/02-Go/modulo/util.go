/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
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

// IngresarEntero devuelve el entero ingresado por el usuario como respuesta a una pregunta.
func IngresarEntero(pregunta string) int {
	entero, err := strconv.Atoi(IngresarTexto(pregunta))

	if err != nil {
		MostrarError("El valor ingresado es inválido, intente nuevamente.\n\n")
		entero = IngresarEntero(pregunta)
	}

	return entero
}

// IngresarReal devuelve el real ingresado por el usuario como respuesta a una pregunta.
func IngresarReal(pregunta string) float64 {
	real, err := strconv.ParseFloat(IngresarTexto(pregunta), 64)

	if err != nil {
		MostrarError("El valor ingresado es inválido, intente nuevamente.\n\n")
		real = IngresarReal(pregunta)
	}

	return real
}

// IngresarLogico devuelve el booleano ingresado por el usuario como respuesta a una pregunta.
func IngresarLogico(pregunta string) bool {
	for {
		opcion := strings.ToLower(IngresarTexto(pregunta))

		if opcion == "s" || opcion == "n" {
			return opcion == "s"
		}

		MostrarError("No es una opción válida, intente nuevamente.\n\n")
	}
}

// IngresarOpcion devuelve el entero ingresado por el usuario como respuesta a una pregunta con múltiples opciones del 1 a un máximo especificado.
func IngresarOpcion(pregunta string, maximaOpcion int) int {
	opcion := IngresarEntero(pregunta)

	if opcion < 1 || opcion > maximaOpcion {
		MostrarError("La opción no es válida, intente nuevamente.\n\n")
		opcion = IngresarOpcion(pregunta, maximaOpcion)
	}

	return opcion
}
