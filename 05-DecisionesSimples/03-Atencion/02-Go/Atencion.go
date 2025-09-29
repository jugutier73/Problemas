/*
   Programa para establecer la prioridad de atención
   en los centros de salud para identificar a los pacientes con
   mayor riesgo o vulnerabilidad
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"atencion/modulo"
)

const (
	EDAD_RECIEN_NACIDO = 1
	EDAD_ADULTO_MAYOR  = 60
)

func main() {
	edadPaciente := modulo.IngresarEntero("Edad del paciente: ")
	enfermedadCronica := modulo.IngresarLogico("Enfermedad crónica (s/n): ")
	estadoInmunosupresion := modulo.IngresarLogico("Estado de inmunosupresión (s/n): ")

	reporteAtencion := generarReporteAtencion(edadPaciente,
		enfermedadCronica, estadoInmunosupresion)

	modulo.MostrarMensaje(reporteAtencion)
}

func generarReporteAtencion(
	edadPaciente int, enfermedadCronica bool, estadoInmunosupresion bool) string {
	mensaje := "\nEl paciente es de atención general\n"

	if edadPaciente < EDAD_RECIEN_NACIDO ||
		edadPaciente > EDAD_ADULTO_MAYOR ||
		enfermedadCronica ||
		estadoInmunosupresion {
		mensaje = "\nEl paciente es de atención prioritaria\n"
	}

	return mensaje
}
