/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
*/

package modulo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MostrarMensaje mostrar un mensaje (cadena) en la salida estandar.
func MostrarMensaje(mensaje string) {
	fmt.Print(mensaje)
}

// IngresarTexto muestra una pregunta y devuelve el texto ingresado por el usuario.
func IngresarTexto(pregunta string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(pregunta)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

// IngresarEntero muestra una pregunta y devuelve el entero ingresado por el usuario o cero si es un valor inválido.
func IngresarEntero(pregunta string) int {
	entero, _ := strconv.Atoi(IngresarTexto(pregunta))
	return entero
}

// IngresarReal muestra una pregunta y devuelve el real ingresado por el usuario o cero si es un valor inválido.
func IngresarReal(pregunta string) float64 {
	real, _ := strconv.ParseFloat(IngresarTexto(pregunta), 64)
	return real
}