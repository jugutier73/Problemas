/*
   Programa que registra los residuos recolectados durante una
   jornada, especificando su código, tipo y condición de
   reparabilidad. El sistema genera un informe con el
   listado de los elementos recolectados, ordenados por tipo y
   reparabilidad, además de presentar un conteo total por cada
   tipo de residuo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"recoleccion/modulo"
)

const (
	ANCHO = 15

	TOTAL_TIPOS = 5

	BATERIA          = 1
	TELEFONO         = 2
	COMPUTADOR       = 3
	CARGADOR         = 4
	OTRO_DISPOSITIVO = 5
)

var ETIQUETAS = map[int]string{
	BATERIA:          "Batería",
	TELEFONO:         "Teléfono",
	COMPUTADOR:       "Computador",
	CARGADOR:         "Cargador",
	OTRO_DISPOSITIVO: "Otro",
}

type Residuo struct {
	Codigo    string
	Tipo      int
	Reparable bool
}

func main() {
	residuos := modulo.IngresarColeccion(ingresarResiduo)

	residuosClasificados := modulo.OrdenarColeccion(residuos,
		compararCampoTipo, false)

	cantidadBaterias := modulo.ContarSegunCriterio(residuos,
		tenerTipo, BATERIA)

	cantidadTelefonos := modulo.ContarSegunCriterio(residuos,
		tenerTipo, TELEFONO)

	cantidadComputadores := modulo.ContarSegunCriterio(residuos,
		tenerTipo, COMPUTADOR)

	cantidadCargadores := modulo.ContarSegunCriterio(residuos,
		tenerTipo, CARGADOR)

	cantidadOtros := modulo.ContarSegunCriterio(residuos,
		tenerTipo, OTRO_DISPOSITIVO)

	reporteRecoleccion := generarReporteRecoleccion(
		residuosClasificados,
		cantidadBaterias,
		cantidadTelefonos,
		cantidadComputadores,
		cantidadCargadores,
		cantidadOtros)

	modulo.MostrarMensaje(reporteRecoleccion)
}

func ingresarResiduo() Residuo {
	modulo.MostrarMensaje("\nIngrese los datos del residuo electrónico:\n")
	codigo := modulo.IngresarTexto("\tIngrese el código del residuo: ")
	tipo := modulo.IngresarOpcion("\tTIPOS DE RESIDUOS\n"+
		"\t\t1: Batería\n"+
		"\t\t2: Teléfono\n"+
		"\t\t3: Computador\n"+
		"\t\t4: Cargador\n"+
		"\t\t5: Otro dispositivo)\n"+
		"\tIngrese tipo de residuo      : ",
		TOTAL_TIPOS)
	reparable := modulo.IngresarLogico("\tPuede ser reparado (s/n): ")

	return Residuo{Codigo: codigo, Tipo: tipo, Reparable: reparable}
}

func compararCampoTipo(residuo1 Residuo, residuo2 Residuo) bool {
	return (residuo1.Tipo > residuo2.Tipo) ||
		(residuo1.Tipo == residuo2.Tipo && residuo1.Reparable != residuo2.Reparable)
}

func tenerTipo(residuo Residuo, tipoInteres int) bool {
	return residuo.Tipo == tipoInteres
}

func generarReporteRecoleccion(
	residuosClasificados []Residuo,
	cantidadBaterias int,
	cantidadTelefonos int,
	cantidadComputadores int,
	cantidadCargadores int,
	cantidadOtros int) string {
	listadoResiduosClasificados := modulo.ConvertirColeccionCadena(
		"LISTADO DE RESIDUOS CLASIFICADOS",
		residuosClasificados,
		convertirResiduoCadena)

	return fmt.Sprintf("%s\n\n"+
		"Cantidad de Baterías     : %d\n"+
		"Cantidad de Teléfonos    : %d\n"+
		"Cantidad de Computadores : %d\n"+
		"Cantidad de Cargadores   : %d\n"+
		"Cantidad de Otros        : %d\n",
		listadoResiduosClasificados,
		cantidadBaterias,
		cantidadTelefonos,
		cantidadComputadores,
		cantidadCargadores,
		cantidadOtros)
}

func convertirResiduoCadena(residuo Residuo) string {
	etiqueta := ETIQUETAS[residuo.Tipo]
	reparable := ""

	if residuo.Reparable {
		reparable = "(REPARABLE)"
	}

	return fmt.Sprintf("%-*s %-*s %s\n",
		ANCHO, residuo.Codigo,
		ANCHO, etiqueta, reparable)
}
