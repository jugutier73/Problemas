/*
   Programa para calcular el índice de masa corporal (IMC), usando la
   altura y el peso.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"imc/modulo"
	"math"
	"strconv"
)

func main() {
	peso := modulo.IngresarEntero("Peso  (kg): ")
	altura := ingresarReal("Altura (m): ")
	imc := calcularIMC(peso, altura)
	informeIMC := generarInformeIMC(peso, altura, imc)
	modulo.MostrarMensaje(informeIMC)
}

func ingresarReal(pregunta string) float64 {
	real, _ := strconv.ParseFloat(modulo.IngresarTexto(pregunta), 64)
	return real
}

func calcularIMC(peso int, altura float64) float64 {
	return float64(peso) / math.Pow(altura, 2.0)
}

func generarInformeIMC(peso int, altura float64, imc float64) string {
	return fmt.Sprintf(
		"\nCon su peso de %d kg y su altura de %.1f metros"+
			"\nsu índice de masa corporal (IMC) es de %.1f", peso, altura, imc)
}
