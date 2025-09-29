/*
   Crear un programa para determinar si hay problemas con el uso
   de los espacios en un frase que el usuario ingrese

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"alfabetizacion/modulo"
	"fmt"
)

func main() {
	frase := modulo.IngresarTexto("Ingrese una frase a analizar: ")

	fraseCorregida := corregirUsoEspacios(frase)

	reporte := generarReporteAlfabetizacion(frase, fraseCorregida)

	modulo.MostrarMensaje(reporte)
}

func corregirUsoEspacios(frase string) string {
	inicioFrase := false
	espacioPendiente := false
	fraseCorregida := ""

	for _, simbolo := range frase {
		if simbolo == ' ' {
			if inicioFrase {
				espacioPendiente = true
			}
		} else {
			if espacioPendiente {
				fraseCorregida += " "
				espacioPendiente = false
			}
			fraseCorregida += string(simbolo)
			inicioFrase = true
		}
	}

	return fraseCorregida
}

func generarReporteAlfabetizacion(frase string, fraseCorregida string) string {
	reporteAlfabetizacion := fmt.Sprintf("\nLa frase \"%s\" hace un uso ", frase)

	if frase == fraseCorregida {
		reporteAlfabetizacion += "CORRECTO de espacios."
	} else {
		reporteAlfabetizacion += fmt.Sprintf("INCORRECTO de espacios,\n"+
			"lo correcto es \"%s\"", fraseCorregida)
	}

	return reporteAlfabetizacion + "\n"
}
