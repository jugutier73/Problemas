/*
   Crear un programa para registrar síntomas y recibir recomendaciones
   simples y seguras como reposar, hidratarse o consultar a un profesional
   de la salud, según la gravedad o combinación de síntomas reportados.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"diagnostico/modulo"
)

const (
	FIEBRE_ALTA = 39.5
	FIEBRE      = 37.5

	MENSAJE_ESPECIALISTA = `
	- Consultar un especialista.
	- Anotar los síntomas y cuándo comenzaron.
	- Evitar esfuerzos físicos y actividades intensas.`

	MENSAJE_FIEBRE_ALTA = `
	- Solicitar una cita medica con urgencia.
	- Usar paños húmedos y fríos en la frente.
	- Permanecer en un lugar fresco y ventilado.`

	MENSAJE_FIEBRE = `
	- Descansar lo suficiente.
	- Hidratarse bebiendo agua u otros líquidos.`

	MENSAJE_DOLOR_CABEZA = `
	- Realizar ejercicio cervical isométrico.
	- Descansar en un lugar oscuro y silencioso.
	- Beber agua, ya que la deshidratación puede empeorar el dolor.`

	MENSAJE_CONGESTION = `
	- Realizar lavados nasales con solución salina.
	- Usar almohadas extras para dormir con la cabeza elevada.`
)

func main() {
	temperatura := modulo.IngresarReal("Temperatura corporal: ")
	sintomasVariosDias := modulo.IngresarLogico("Síntomas por 2 o más días (s/n): ")
	malestarIntenso := modulo.IngresarLogico("Tiene malestar intenso (s/n):")
	dolorCabeza := modulo.IngresarLogico("Tiene dolor de cabeza (s/n): ")
	congestionNasal := modulo.IngresarLogico("Tiene congestion nasal (s/n):")

	recomendacionesEspecialista := recibirRecomendaciones(
		sintomasVariosDias || malestarIntenso, MENSAJE_ESPECIALISTA)

	recomendacionesFiebre := recibirRecomendaciones(
		temperatura >= FIEBRE, MENSAJE_FIEBRE)

	recomendacionesFiebreAlta := recibirRecomendaciones(
		temperatura >= FIEBRE_ALTA, MENSAJE_FIEBRE_ALTA)

	recomendacionesDolorCabeza := recibirRecomendaciones(
		dolorCabeza, MENSAJE_DOLOR_CABEZA)

	recomendacionesCongestion := recibirRecomendaciones(
		congestionNasal, MENSAJE_CONGESTION)

	recomendaciones := generarReporteRecomendaciones(
		recomendacionesEspecialista, recomendacionesFiebreAlta,
		recomendacionesFiebre, recomendacionesDolorCabeza,
		recomendacionesCongestion)

	modulo.MostrarMensaje(recomendaciones)
}

func recibirRecomendaciones(condicion bool, recomendaciones string) string {
	mensajeRecomendaciones := ""

	if condicion {
		mensajeRecomendaciones = recomendaciones
	}

	return mensajeRecomendaciones
}

func generarReporteRecomendaciones(recomendacionesEspecialista string,
	recomendacionesFiebreFuerte string, recomendacionesFiebre string,
	recomendacionesDolorCabeza string, recomendacionesCongestion string) string {
	reporteRecomendaciones := "\nSe recomienda:"

	if recomendacionesEspecialista != "" {
		reporteRecomendaciones += recomendacionesEspecialista
	} else {
		reporteRecomendaciones += recomendacionesFiebreFuerte +
			recomendacionesFiebre + recomendacionesDolorCabeza +
			recomendacionesCongestion
	}

	return reporteRecomendaciones + "\n"
}
