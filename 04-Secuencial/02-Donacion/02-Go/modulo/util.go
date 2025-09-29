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
	"strings"
)

// MostrarMensaje mostrar un mensaje (cadena) en la salida estandar.
func MostrarMensaje(mensaje string) {
	fmt.Print(mensaje)
}

// IngresarTexto devuelve el texto ingresado por el usuario como respuesta a una pregunta.
func IngresarTexto(pregunta string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(pregunta)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}
