/*
   Crear un programa para garantizar que los pacientes sean atendidos
   de manera equitativa y oportuna. El orden de llegada debe
   equilibrarse con la prioridad médica de cada caso

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"citas/modulo"
	"fmt"
)

const (
	NIVELES = 3

	NIVEL_RIESGO      = 1
	NIVEL_URGENCIA    = 2
	NIVEL_PRIORITARIO = 3
)

type Cita struct {
	Nombre string
	Nivel  int
}

func main() {
	citas := modulo.IngresarColeccion(ingresarCita)

	citasPorNivel := modulo.OrdenarColeccion(
		citas, compararCampoNivel, false)

	cantidadRiesgo := contarSegunCriterio(
		citas, tenerNivelRiesgo, NIVEL_RIESGO)

	cantidadUrgencia := contarSegunCriterio(
		citas, tenerNivelRiesgo, NIVEL_URGENCIA)

	cantidadPrioritario := contarSegunCriterio(
		citas, tenerNivelRiesgo, NIVEL_PRIORITARIO)

	reporteCitas := generarReporteCitas(
		citasPorNivel,
		cantidadRiesgo,
		cantidadUrgencia,
		cantidadPrioritario)

	modulo.MostrarMensaje(reporteCitas)
}

func ingresarCita() Cita {
	modulo.MostrarMensaje("\nIngrese los datos de la cita:\n")
	nombre := modulo.IngresarTexto("\tIngrese el nombre paciente: ")
	nivel := modulo.IngresarOpcion("\tNIVEL DE URGENCIA\n"+
		"\t\t1: Riesgo\n"+
		"\t\t2: Urgencia\n"+
		"\t\t3: Prioritario\n"+
		"\tIngrese tipo de persona: ",
		NIVELES)

	return Cita{Nombre: nombre, Nivel: nivel}
}

func compararCampoNivel(cita1 Cita, cita2 Cita) bool {
	return cita1.Nivel > cita2.Nivel
}

func contarSegunCriterio[Tipo1 any, Tipo2 any](coleccion []Tipo1, aplicarCriterio func(Tipo1, Tipo2) bool, valorCriterio Tipo2) int {
	cantidad := 0

	for _, elemento := range coleccion {
		if aplicarCriterio(elemento, valorCriterio) {
			cantidad += 1
		}
	}

	return cantidad
}

func tenerNivelRiesgo(cita Cita, nivelRiesgo int) bool {
	return cita.Nivel == nivelRiesgo
}

func generarReporteCitas(
	citasPorNivel []Cita,
	cantidadRiesgo int,
	cantidadUrgencia int,
	cantidadPrioritario int) string {
	listadoPorNivel := modulo.ConvertirColeccionCadena(
		"LISTADO POR NIVEL DE URGENCIA",
		citasPorNivel,
		convertirCitaCadena)

	return fmt.Sprintf("%s\n\n"+
		"Nivel 1 (Riesgo)     : %d\n"+
		"Nivel 2 (Urdencia)   : %d\n"+
		"Nivel 3 (Prioritario): %d\n",
		listadoPorNivel,
		cantidadRiesgo,
		cantidadUrgencia,
		cantidadPrioritario)
}

func convertirCitaCadena(cita Cita) string {
	return fmt.Sprintf("%d : %s\n", cita.Nivel, cita.Nombre)
}
