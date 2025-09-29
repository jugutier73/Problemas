/*
   Crear un programa para aplicar un descuento a la tarifa del
   transporte público según el perfil del usuario, como estudiantes o
   adultos mayores. Considerando que los domingos y festivos la tarifa
   es mayor.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"descuento/modulo"
	"fmt"
)

const (
	TARIFA_NORMAL          = 2900
	TARIFA_DOMINGO_FESTIVO = 3000

	PORCENTAJE_DESCUENTO_ESTUDIANTE   = 10.0
	PORCENTAJE_DESCUENTO_ADULTO_MAYOR = 15.0
)

func main() {
	esDomingoFestivo := modulo.IngresarLogico("Es domingo o festivo (s/n): ")
	esEstudiante := modulo.IngresarLogico("Es estudiante (s/n): ")
	esAdultoMayor := modulo.IngresarLogico("Es adulto mayor (s/n): ")

	tarifaDia := obtenerTarifaDia(esDomingoFestivo)

	porcentajeDescuento := obtenerPorcentajeDescuento(
		esEstudiante, esAdultoMayor)

	valorDescuento := calcularValorDescuento(
		tarifaDia, porcentajeDescuento)

	valorTarifa := calcularValorTarifa(
		tarifaDia, valorDescuento)

	reciboTarifa := generarReciboTarifa(
		tarifaDia, valorTarifa, porcentajeDescuento, valorDescuento)

	modulo.MostrarMensaje(reciboTarifa)
}

func obtenerTarifaDia(esDomingoFestivo bool) int {
	tarifaDia := TARIFA_NORMAL

	if esDomingoFestivo {
		tarifaDia = TARIFA_DOMINGO_FESTIVO
	}

	return tarifaDia
}

func obtenerPorcentajeDescuento(esEstudiante bool, esAdultoMayor bool) float64 {
	porcentajeDescuento := 0.0

	if esEstudiante {
		porcentajeDescuento = PORCENTAJE_DESCUENTO_ESTUDIANTE
	}

	if esAdultoMayor {
		porcentajeDescuento = PORCENTAJE_DESCUENTO_ADULTO_MAYOR
	}
	return porcentajeDescuento
}

func calcularValorDescuento(tarifaDia int, porcentajeDescuento float64) float64 {
	return float64(tarifaDia) * porcentajeDescuento / 100.0
}

func calcularValorTarifa(tarifaDia int, valorDescuento float64) float64 {
	return float64(tarifaDia) - valorDescuento
}

func generarReciboTarifa(
	tarifaDia int, valorTarifa float64, porcentajeDescuento float64,
	valorDescuento float64) string {

	mensaje := fmt.Sprintf("\nLa tarifa es de %d\n", tarifaDia)

	if valorDescuento > 0 {
		mensaje = fmt.Sprintf(
			"\n%s"+
				"\nsu tarifa a pagar es de $%.0f por tener"+
				"\nun descuento del %.1f%% "+
				"equivalente  a $%.0f\n",
			mensaje, valorTarifa, porcentajeDescuento, valorDescuento)
	}

	return mensaje
}
