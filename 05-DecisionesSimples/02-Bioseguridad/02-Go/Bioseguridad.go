/*
   Crear un programa para programa para la verificación de medidas
   de bioseguridad para el ingreso a un evento público.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"bioseguridad/modulo"
	"strings"
)

func main() {
	tieneVacunas := ingresarLogico("Tiene vacunas (s/n): ")
	restadosNegativos := ingresarLogico("Pruebas negativas (s/n): ")
	tieneSintomas := ingresarLogico("Tiene síntomas (s/n): ")

	reporteIngreso := generarReporteIngreso(
		tieneVacunas, restadosNegativos, tieneSintomas)

	modulo.MostrarMensaje(reporteIngreso)
}

func ingresarLogico(pregunta string) bool {
	opcion := strings.ToLower(modulo.IngresarTexto(pregunta))

	if opcion != "s" && opcion != "n" {
		modulo.MostrarError("No es una opción válida, se asume \"n\"\n\n")
	}

	return opcion == "s"
}

func generarReporteIngreso(
	tieneVacunas bool, restadosNegativos bool, tieneSintomas bool) string {
	mensaje := "\nLa persona no puede ingresar al evento\n"

	if tieneVacunas && restadosNegativos && !tieneSintomas {
		mensaje = "\nLa persona puede ingresar al evento\n"
	}

	return mensaje
}
