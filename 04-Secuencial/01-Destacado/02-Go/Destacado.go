/*
   Programa para la empresa de seguridad Aros S.A.
   que permita resaltar el nombre del empleado destacado
   del mes ofreciendo unas felicitaciones públicas.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"destacado/modulo"
	"fmt"
)

func main() {
	nombreEmpleado := ingresarTexto("Nombre del empleado destacado: ")
	nombreMes := ingresarTexto("Nombre del mes: ")
	felicitaciones := generarFelicitaciones(nombreEmpleado, nombreMes)
	modulo.MostrarMensaje(felicitaciones)
}

func ingresarTexto(pregunta string) string {
	var texto string
	fmt.Print(pregunta)
	fmt.Scanln(&texto)
	return texto
}

func generarFelicitaciones(nombreEmpleado, nombreMes string) string {
	return fmt.Sprintf(
		"\nLa empresa de seguridad Aros S.A. quiere felicitar"+
			"\npúblicamente a %s como nuestro"+
			"\nempleado destacado del mes de %s,"+
			"\nmuchas gracias por su excelencia.", nombreEmpleado, nombreMes)
}
