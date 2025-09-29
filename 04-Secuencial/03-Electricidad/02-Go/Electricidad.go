/*
   Programa que permita tener un control del consumo de la energía
   eléctrica con relación al consumo del mes anterior, para que
   así el usuario pueda adoptar hábitos de consumo más conscientes
   y responsables.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"consumo/modulo"
	"fmt"
)

func main() {
	consumoActual := modulo.IngresarEntero(
		"Consumo mes actual   (kilovatios): ")
	consumoAnterior := modulo.IngresarEntero(
		"Consumo mes anterior (kilovatios): ")

	relacionConsumo := calcularRelacionConsumo(consumoActual,
		consumoAnterior)

	reporteRelacion := generarReporeRelacion(consumoActual, consumoAnterior,
		relacionConsumo)

	modulo.MostrarMensaje(reporteRelacion)
}

func calcularRelacionConsumo(consumoActual int, consumoAnterior int) float64 {
	return float64(consumoActual) / float64(consumoAnterior) * 100.0
}

func generarReporeRelacion(consumoActual int, consumoAnterior int, relacionConsumo float64) string {
	return fmt.Sprintf(
		"\nEl consumo actual de %d kilovatios representa"+
			"\nun %.1f%% con relación al consumo del mes"+
			"\nanterior de %d kilovatios.\n", consumoActual, relacionConsumo, consumoAnterior)
}
