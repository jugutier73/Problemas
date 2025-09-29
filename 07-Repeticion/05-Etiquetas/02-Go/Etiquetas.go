/*
   Programa para verificar si una etiqueta de reciclaje es válida
   y reportar los campos con error.

   Formato: T-aaaa-mm
   - T: tipo de residuo
     (P: plástico, V: vidrio, M: metal, C: cartón/papel, O: orgánico)
   - aaaa: año (> 2025)
   - mm: mes (01-12)

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"etiquetas/modulo"
	"fmt"
	"strconv"
	"strings"
)

const (
	LONGITUD_ETIQUETA = 9 //T-aaaa-mm

	SEPARADOR           = '-'
	POSICION_SEPARADOR1 = 1
	POSICION_SEPARADOR2 = 6

	POSICION_TIPO = 0

	POSICION_ANIO = 2
	LONGITUD_ANIO = 4

	POSICION_MES = 7
	LONGITUD_MES = 2

	MINIMO_ANIO = 2026
	MINIMO_MES  = 1
	MAXIMO_MES  = 12

	TIPOS_VALIDOS = "PVMCO"

	ETIQUETA_OK    = 0
	ERROR_LONGITUD = 1
	ERROR_FORMATO  = 2
	ERROR_TIPO     = 4
	ERROR_ANIO     = 8
	ERROR_MES      = 16
)

func main() {
	etiqueta := modulo.IngresarTexto("Ingrese la etiqueta a analizar: ")

	codigo := verificarEtiquetaReciclaje(etiqueta)

	reporteEtiqueta := generarReporteEtiqueta(etiqueta, codigo)

	modulo.MostrarMensaje(reporteEtiqueta)
}

func verificarEtiquetaReciclaje(etiqueta string) int {
	codigo := verificarLongitud(etiqueta)

	if codigo != ERROR_LONGITUD {
		codigo = verificarFormato(etiqueta) |
			verificarTipo(etiqueta) |
			verificarAnio(etiqueta) |
			verificarMes(etiqueta)
	}

	return codigo
}

func verificarLongitud(etiqueta string) int {
	codigoError := ETIQUETA_OK

	if len(etiqueta) != LONGITUD_ETIQUETA {
		codigoError = ERROR_LONGITUD
	}

	return codigoError
}

func verificarFormato(etiqueta string) int {
	codigoError := ETIQUETA_OK

	if etiqueta[POSICION_SEPARADOR1] != SEPARADOR ||
		etiqueta[POSICION_SEPARADOR2] != SEPARADOR {
		codigoError = ERROR_FORMATO
	}

	return codigoError
}

func verificarTipo(etiqueta string) int {
	codigoError := ETIQUETA_OK

	tipo := rune(etiqueta[POSICION_TIPO])

	if !strings.ContainsRune(TIPOS_VALIDOS, tipo) {
		codigoError = ERROR_TIPO
	}

	return codigoError
}

func verificarAnio(etiqueta string) int {
	codigoError := ETIQUETA_OK

	anioEtiqueta := etiqueta[POSICION_ANIO : POSICION_ANIO+LONGITUD_ANIO]

	anio, error := strconv.Atoi(anioEtiqueta)
	if error != nil || anio < MINIMO_ANIO {
		codigoError = ERROR_ANIO
	}

	return codigoError
}

func verificarMes(etiqueta string) int {
	codigoError := ETIQUETA_OK

	mesEtiqueta := etiqueta[POSICION_MES : POSICION_MES+LONGITUD_MES]

	mes, error := strconv.Atoi(mesEtiqueta)
	if error != nil || mes < MINIMO_MES || mes > MAXIMO_MES {
		codigoError = ERROR_MES
	}

	return codigoError
}

func generarReporteEtiqueta(etiqueta string, codigo int) string {
	reporte := fmt.Sprintf("\nLa  etiqueta \"%s\" ", etiqueta)

	if codigo == ETIQUETA_OK {
		reporte += "es válida.\n"
	} else {
		reporte += "es inválida por tener:"

		if codigo & ERROR_LONGITUD != 0 {
			reporte += "\n  - Error en la longitud"
		}

		if codigo & ERROR_FORMATO != 0 {
			reporte += "\n  - Error en el formato"
		}

		if codigo & ERROR_TIPO != 0 {
			reporte += "\n  - Error en el tipo"
		}

		if codigo & ERROR_ANIO != 0 {
			reporte += "\n  - Error en el año"
		}

		if codigo & ERROR_MES != 0 {
			reporte += "\n  - Error en el mes"
		}
	}

	return reporte + "\n"
}
