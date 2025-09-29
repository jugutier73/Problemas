/*
   Programa para verificar la elegibilidad de los donantes
   con base en la edad, el peso y la madurez fisiológica
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"donar/modulo"
)

const (
	EDAD_MINIMA              = 18
	EDAD_MAXIMA              = 65
	EDAD_MINIMA_AUTORIZACION = 16
	EDAD_MAXIMA_AUTORIZACION = 18
	PESO_MINIMO              = 50.0
)

func main() {
	edadDonante := modulo.IngresarEntero("Edad del donante: ")
	autorizacionEdad := modulo.IngresarLogico(
		"Tiene autorización por edad (s/n): ")
	pesoDonante := modulo.IngresarReal("Peso del donante: ")
	suficienteMadurez := modulo.IngresarLogico(
		"Suficiente madurez fisiológica(s/n): ")

	reporteLegibilidad := generarReporteEligibilidad(
		edadDonante, autorizacionEdad, pesoDonante, suficienteMadurez)

	modulo.MostrarMensaje(reporteLegibilidad)
}

func generarReporteEligibilidad(
	edadDonante int, autorizacionEdad bool, pesoDonante float64, suficienteMadurez bool) string {
	mensaje := "\nEl donante no cumple las condiciones para donar\n"

	if (edadDonante >= EDAD_MINIMA && edadDonante <= EDAD_MAXIMA ||
		edadDonante >= EDAD_MINIMA_AUTORIZACION &&
			edadDonante <= EDAD_MAXIMA_AUTORIZACION && autorizacionEdad) &&
		pesoDonante > PESO_MINIMO && suficienteMadurez {
		mensaje = "\nEl donante es elegible para donar\n"
	}

	return mensaje
}
