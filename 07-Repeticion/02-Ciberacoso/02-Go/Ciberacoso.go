/*
   Crear un programa para contar cuantas palabras comienzan con
    la letra "p" en un mensaje.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"ciberacoso/modulo"
	"fmt"
	"strings"
)

const (
	LETRA_INICIO = 'p'
)

func main() {
	texto := modulo.IngresarTexto("Ingrese un texto con los mensajes a analizar: ")

	cantidadPalabras := contarPalabrasInician(texto, LETRA_INICIO)

	reporteCiberacoso := generarReporteAcoso(cantidadPalabras)

	modulo.MostrarMensaje(reporteCiberacoso)
}

func contarPalabrasInician(texto string, letraInicio byte) int {
	palabras := strings.Split(strings.ToLower(texto), " ")

	cantidadPalabrasInteres := 0

	for _, palabra := range palabras {
		if strings.HasPrefix(palabra, string(letraInicio)) {
			cantidadPalabrasInteres++
		}
	}

	return cantidadPalabrasInteres
}

func generarReporteAcoso(cantidadPalabrasInteres int) string {
	return fmt.Sprintf("\nHay %d palabras "+
		"que inician con la letra \"%c\".\n",
		cantidadPalabrasInteres, LETRA_INICIO)
}
