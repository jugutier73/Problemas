/*
   Crear un programa para a clasificación de residuos según su
   material. Residuos aprovechable (limpios y secos), usar bolsa
   blanca; residuos orgánicos, usar bolsa verde; lo anterior si
   existe ruta de recolección selectiva; en otro caso, usar bolsa
   de color negro.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"residuo/modulo"
)

const (
	APROVECHABLE = 1
	OTRO         = 3

	MAX_OPCIONES = 3
)

func main() {
	tipoResiduo := ingresarOpcion("Residuo \n"+
		"\t(1) aprovechable (limpia y seca),\n"+
		"\t(2) orgánico,\n"+
		"\t(3) otro\n\n"+
		"Cuál es el tipo de residuo: ", MAX_OPCIONES)

	hayRecoleccionSelectiva := modulo.IngresarLogico(
		"Hay ruta de recolección selectiva (s/n): ")

	colorBolsa := recomendarColorBolsa(tipoResiduo, hayRecoleccionSelectiva)
	
	reporteBolsaRecoleccion := generarReporteBolsa(colorBolsa)

	modulo.MostrarMensaje(reporteBolsaRecoleccion)
}

func ingresarOpcion(pregunta string, maximaOpcion int) int {
	opcion := modulo.IngresarEntero(pregunta)

	if opcion < 1 || opcion > maximaOpcion {
		modulo.MostrarError("La opción no es válida, se asume 1\n\n")
		opcion = 1
	}

	return opcion
}

func recomendarColorBolsa(tipoResiduo int,
	hayRecoleccionSelectiva bool) string {
	var colorBolsa string

	if tipoResiduo == OTRO || !hayRecoleccionSelectiva {
		colorBolsa = "NEGRA"
	} else {
		if tipoResiduo == APROVECHABLE {
			colorBolsa = "BLANCA"
		} else {
			colorBolsa = "VERDE"
		}
	}

	return colorBolsa
}

func generarReporteBolsa(colorBolsa string) string {
	return fmt.Sprintf(
		"\nLa bolsa recomendada es de color %s.\n", colorBolsa)
}
