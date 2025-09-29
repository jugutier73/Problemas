/*
   Programa para la empresa de seguridad Aros S.A.
   que permita resaltar el nombre del empleado destacado
   del mes ofreciendo unas felicitaciones públicas.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"bufio"
	"destacado/modulo"
	"fmt"
	"os"
	"strings"
)

func main() {
	nombreEmpleado := ingresarTexto("Nombre del empleado destacado: ")
	nombreMes := ingresarTexto("Nombre del mes: ")

	felicitaciones := generarFelicitaciones(nombreEmpleado, nombreMes)

	modulo.MostrarMensaje(felicitaciones)
}

func ingresarTexto(pregunta string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(pregunta)
	texto, _ := reader.ReadString('\n')

	return strings.TrimSpace(texto)
}

func generarFelicitaciones(nombreEmpleado, nombreMes string) string {
	return fmt.Sprintf(
		"\nLa empresa de seguridad Aros S.A. quiere felicitar"+
			"\npúblicamente a %s como nuestro"+
			"\nempleado destacado del mes de %s,"+
			"\nmuchas gracias por su excelencia.\n", nombreEmpleado, nombreMes)
}
