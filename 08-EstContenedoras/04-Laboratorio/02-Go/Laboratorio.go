/*
   Crear un programa para registrar el ingreso y salida de los
   usuarios de un laboratorio para asegurar que solo personas
   autorizadas -estudiantes, docentes o personal técnico- accedan
   al espacio. El programa debe permitir identificar quiénes estaban
   presentes y la cantidad por cada tipo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"laboratorio/modulo"
)

const (
	ANCHO = 20

	TOTAL_TIPO_AUTORIZADOS = 3

	ESTUDIANTE = 1
	DOCENTE    = 2
	TECNICO    = 3
)

var ETIQUETAS = map[int]string{
	ESTUDIANTE: "Estudiante",
	DOCENTE:    "Docente",
	TECNICO:    "Técnico",
}

type Ingreso struct {
	Documento string
	Nombre    string
	Tipo      int
}

func main() {
	ingresos := modulo.IngresarColeccion(ingresarIngreso)
	salidas  := modulo.IngresarColeccion(ingresarSalida)

	personasLaboratorio := obtenerPersonasLaboratorio(ingresos, salidas)

	cantidadEstudiantes := modulo.ContarSegunCriterio(
		personasLaboratorio, tenerTipo, ESTUDIANTE)

	cantidadDocentes := modulo.ContarSegunCriterio(
		personasLaboratorio, tenerTipo, DOCENTE)

	cantidadTecnicos := modulo.ContarSegunCriterio(
		personasLaboratorio, tenerTipo, TECNICO)

	reporteIngresos  := generarReporteIngresos(
		ingresos,
		salidas,
		personasLaboratorio,
		cantidadEstudiantes,
		cantidadDocentes,
		cantidadTecnicos)

	modulo.MostrarMensaje(reporteIngresos)
}

func ingresarIngreso() Ingreso {
	modulo.MostrarMensaje("\nIngrese los datos del ingreso:\n")
	documento := modulo.IngresarTexto("\tIngrese el documento   : ")
	nombre    := modulo.IngresarTexto("\tIngrese el nombre      : ")
	tipo      := modulo.IngresarOpcion(
		"\tTIPO DE PERSONAS AUTORIZADAS\n"+
			"\t\t1: Estudiante\n"+
			"\t\t2: Docente\n"+
			"\t\t3: Técnico\n"+
			"\tIngrese tipo de persona: ",
		TOTAL_TIPO_AUTORIZADOS)

	return Ingreso{Documento: documento, Nombre: nombre, Tipo: tipo}
}

func ingresarSalida() string {
	modulo.MostrarMensaje("\nIngrese los datos de la salida:\n")
	documento := modulo.IngresarTexto("\tIngrese el documento   : ")

	return documento
}

func set[Tipo comparable](coleccion []Tipo) map[Tipo]struct{} {
	conjunto := make(map[Tipo]struct{}, len(coleccion))

	for _, elemento := range coleccion {
		conjunto[elemento] = struct{}{}
	}

	return conjunto
}

func obtenerPersonasLaboratorio(ingresos []Ingreso, salidas []string) []Ingreso {
	var personasLaboratorio []Ingreso
	salidasSet := set(salidas)

	for _, ingreso := range ingresos {
		if _, ok := salidasSet[ingreso.Documento]; !ok {
			personasLaboratorio = append(personasLaboratorio, ingreso)
		}
	}

	return personasLaboratorio
}

func tenerTipo(ingreso Ingreso, tipo int) bool {
	return ingreso.Tipo == tipo
}

func generarReporteIngresos(
	ingresos []Ingreso,
	salidas []string,
	personasLaboratorio []Ingreso,
	cantidadEstudiantes int,
	cantidadDocentes int,
	cantidadTecnicos int) string {
	listadoIngresos := modulo.ConvertirColeccionCadena(
		"LISTADO DE INGRESOS", ingresos, convertirIngresoCadena)

	listadoSalidas := modulo.ConvertirColeccionCadena(
		"LISTADO DE SALIDAS", salidas, convertirSalidaCadena)

	listadoPersonasLaboratorio := modulo.ConvertirColeccionCadena(
		"LISTADO DE PERSONAS EN EL LABORATORIO",
		personasLaboratorio, convertirIngresoCadena)

	return fmt.Sprintf("%s\n\n"+
		"%s\n\n"+
		"%s\n\n"+
		"CANTIDAD DE PERSONAS AÚN EN EL LABORATORIO\n"+
		"Cantidad de Estudiantes  : %d\n"+
		"Cantidad de Docentes     : %d\n"+
		"Cantidad de Técnicos     : %d\n",
		listadoIngresos,
		listadoSalidas,
		listadoPersonasLaboratorio,
		cantidadEstudiantes,
		cantidadDocentes,
		cantidadTecnicos)
}

func convertirIngresoCadena(ingreso Ingreso) string {
	etiqueta := ETIQUETAS[ingreso.Tipo]

	return fmt.Sprintf("%-*s %-*s %s\n",
		ANCHO, ingreso.Documento,
		ANCHO, ingreso.Nombre,
		etiqueta)
}

func convertirSalidaCadena(salida string) string {
	return salida
}
