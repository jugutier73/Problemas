/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
*/

package modulo

import (
	"bufio"
	"fmt"
	"os"
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
