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

func MostrarMensaje(mensaje string) {
	fmt.Print(mensaje)
}

func IngresarTexto(pregunta string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(pregunta)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func IngresarEntero(pregunta string) int {
	entero, _ := strconv.Atoi(IngresarTexto(pregunta))
	return entero
}

func IngresarReal(pregunta string) float64 {
	real, _ := strconv.ParseFloat(IngresarTexto(pregunta), 64)
	return real
}
