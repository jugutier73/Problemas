/*
   Programa para calcular la cantidad de CO2 emitido por el uso de
   transporte particular (carro y moto), el empleo del transporte
   público (buses)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"huella/modulo"
)

const (
	HUELLA_CARRO = 121.0
	HUELLA_MOTO  = 53.0
	HUELLA_BUS   = 49.0
)

func main() {
	kmCarro :=modulo.IngresarReal("Total de kilómetros recorridos en carro:")
	kmMoto := modulo.IngresarReal("Total de kilómetros recorridos en moto :")
	kmBuses := modulo.IngresarReal("Total de kilómetros recorridos en bus :")

	huella := calcularHuellaCarbono(kmCarro, kmMoto, kmBuses)
	informeHuella := generarHuella(kmCarro, kmMoto, kmBuses, huella)

	modulo.MostrarMensaje(informeHuella)
}

func calcularHuellaCarbono(kmCarro float64, kmMoto float64, kmBuses float64) float64 {
	return kmCarro*HUELLA_CARRO + kmMoto*HUELLA_MOTO + kmBuses*HUELLA_BUS
}

func generarHuella(kmCarro float64, kmMoto float64, kmBuses float64, huella float64) string {
	return fmt.Sprintf(
		"\nCon %.1f, %.1f, %.1f km de recorrido"+
			"\nen carro, moto y bus respectivamente,"+
			"\nsu huella de carbono por el uso de"+
			"\ntransporte es de %.1f kg de CO2.\n",
		kmCarro, kmMoto, kmBuses, huella)
}
