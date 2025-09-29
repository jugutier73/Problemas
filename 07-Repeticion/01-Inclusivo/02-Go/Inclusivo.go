/*
   Crear un programa para contar cuantas veces se emplean símbolos
   "x" o la "@" en un mensaje.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"inclusivo/modulo"
)

const (
	SIMBOLO_INCLUSIVO_1 = 'x'
	SIMBOLO_INCLUSIVO_2 = '@'
)

func main() {
	texto := modulo.IngresarTexto("Ingrese un texto con los mensajes: ")

	cantidadSimbolos := contarEmpleoSimbolos(texto)

	reporteSimbolos := generarReporteInclusivo(cantidadSimbolos)

	modulo.MostrarMensaje(reporteSimbolos)
}

func contarEmpleoSimbolos(texto string) int {
	cantidadSimbolosInclusivos := 0

	for _, caracter := range texto {
		if caracter == SIMBOLO_INCLUSIVO_1 ||
			caracter == SIMBOLO_INCLUSIVO_2 {
			cantidadSimbolosInclusivos = cantidadSimbolosInclusivos + 1
		}
	}

	return cantidadSimbolosInclusivos
}

func generarReporteInclusivo(cantidadSimbolosInclusivos int) string {
	return fmt.Sprintf("\nSe emplearon %d veces "+
		"los símbolos inclusivos \"%c\" y \"%c\".\n",
		cantidadSimbolosInclusivos, SIMBOLO_INCLUSIVO_1, SIMBOLO_INCLUSIVO_2)
}
