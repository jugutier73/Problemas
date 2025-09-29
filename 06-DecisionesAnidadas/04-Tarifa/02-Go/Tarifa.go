/*
   Crear un programa para calcular, de forma clara y transparente,
   el valor de un trayecto en taxi considerando variables como
   el horario diurno/nocturno, los recargos de domingos y festivos,
   y el destino (ya sea dentro del perímetro urbano o fuera de él)
   para determinar si la tarifa cobrada es justa o no.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"tarifa/modulo"
)

const (
	URBANO         = 1
	PERIFERIA      = 2
	EXTRA_URBANO   = 3
	VIA_AEROPUERTO = 4

	SIN_COSTO      = 0
	VALOR_MINIMO   = 6000
	VALOR_ESPECIAL = 1300

	VALOR_PERIFERIA      = 3400
	VALOR_EXTRA_URBANO   = 4900
	VALOR_VIA_AEROPUERTO = 1600
)

func main() {
	valorCobrado := modulo.IngresarEntero("Valor cobrado: ")
	valorTaximetro := modulo.IngresarEntero("Valor del taxímetro: ")
	esEspecial := modulo.IngresarLogico("Es domingo, festivo o nocturno (s/n): ")
	destinoTrayecto := modulo.IngresarOpcion("Destino \n"+
		"\t(1) urbano,\n"+
		"\t(2) lugar periférico,\n"+
		"\t(3) extra urbano\n"+
		"\t(4) vía al aeropuerto\n\n"+
		"Cuál es su destino: ", 4)

	tarifaMinima := determinarTarifaMinima(valorTaximetro)
	tarifaEspecial := determinarTarifaEspecial(esEspecial)
	tarifaDestino := determinarTarifaDestino(destinoTrayecto)

	tarifaReal := calcularTarifa(
		tarifaMinima, tarifaEspecial, tarifaDestino)

	mensajeCobro := determinarMensajeCobro(valorCobrado, tarifaReal)

	informeCobro := generarInformeCobro(
		valorCobrado, tarifaReal, mensajeCobro)

	modulo.MostrarMensaje(informeCobro)
}

func determinarTarifaMinima(valorTaximetro int) int {
	return max(valorTaximetro, VALOR_MINIMO)
}

func determinarTarifaEspecial(esEspecial bool) int {
	tarifaEspecial := SIN_COSTO

	if esEspecial {
		tarifaEspecial = VALOR_ESPECIAL
	}

	return tarifaEspecial
}

func determinarTarifaDestino(destinoTrayecto int) int {
	tarifaDestino := SIN_COSTO

	if destinoTrayecto == PERIFERIA {
		tarifaDestino = VALOR_PERIFERIA
	} else if destinoTrayecto == EXTRA_URBANO {
		tarifaDestino = VALOR_EXTRA_URBANO
	} else if destinoTrayecto == VIA_AEROPUERTO {
		tarifaDestino = VALOR_VIA_AEROPUERTO
	}

	return tarifaDestino
}

func calcularTarifa(tarifaMinima int, tarifaEspecial int, tarifaDestino int) int {
	return tarifaMinima + tarifaEspecial + tarifaDestino
}

func determinarMensajeCobro(valorCobrado int, tarifaReal int) string {
	mensajeCobro := "JUSTO"

	if valorCobrado != tarifaReal {
		mensajeCobro = "INJUSTO"
	}

	return mensajeCobro
}

func generarInformeCobro(valorCobrado int, tarifaReal int, mensajeCobro string) string {
	return fmt.Sprintf(
		"\nValor cobrado \t$%7d"+
			"\nValor real \t$%7d\n"+
			"\nPor lo anterior el cobro es %s\n",
		valorCobrado, tarifaReal, mensajeCobro)
}
