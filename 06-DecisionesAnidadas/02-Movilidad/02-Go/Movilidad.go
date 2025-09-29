/*
   Crear un programa para recomendar un medio de transporte
   según el tipo de distancia a la universidad, si está lloviendo,
   y si hay o no transporte público.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"movilidad/modulo"
)

const (
	CERQUITA = 1
	LEJOS    = 3
)

func main() {
	tipoDistancia := modulo.IngresarOpcion("Vive\n"+
		"\t(1) cerquita,\n"+
		"\t(2) cerca,\n"+
		"\t(3) lejos: \n\n"+
		"Cuál es el tipo de distancia: ", 3)
	estaLloviendo := modulo.IngresarLogico("Está lloviendo (s/n): ")
	hayTransporte := modulo.IngresarLogico("Hay transporte público (s/n): ")

	medioTransporte := recomendarMedioTransporte(
		tipoDistancia, estaLloviendo, hayTransporte)
	reporteTransporte := generarReporteTransporte(medioTransporte)

	modulo.MostrarMensaje(reporteTransporte)
}

func recomendarMedioTransporte(tipoDistancia int,
	estaLloviendo bool, hayTransporte bool) string {
	var medioTransporte string

	if tipoDistancia == CERQUITA && !estaLloviendo {
		medioTransporte = "caminar o usar bicicleta"
	} else {
		if tipoDistancia == LEJOS || !hayTransporte {
			medioTransporte = "carro compartido"
		} else {
			medioTransporte = "transporte público"
		}
	}

	return medioTransporte
}

func generarReporteTransporte(medioTransporte string) string {
	return fmt.Sprintf(
		"\nMedio de transporte recomendado: %s.\n", medioTransporte)
}
