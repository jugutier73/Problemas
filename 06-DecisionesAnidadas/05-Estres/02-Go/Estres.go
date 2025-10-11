/*
   Crear un programa para evaluar el nivel de estrés de los estudiantes

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"estres/modulo"
	"fmt"
	"math"
)

const (
	OPCIONES = 5

	NIVEL_BAJO     = 19
	NIVEL_MODERADO = 25
)

func main() {
	respuesta01 := ingresarRespuesta(
		"¿con qué frecuencia te has sentido afectado por algo que ocurrió inesperadamente?")
	respuesta02 := ingresarRespuesta(
		"¿con qué frecuencia te has sentido incapaz de controlar las  cosas importantes en tu vida?")
	respuesta03 := ingresarRespuesta(
		"¿con qué frecuencia te has sentido nervioso o estresado?")
	respuesta04 := ingresarRespuestaInvertida(
		"¿con qué frecuencia has manejado con éxito los pequeños problemas irritantes de la vida?")
	respuesta05 := ingresarRespuestaInvertida(
		"¿con qué frecuencia has sentido que has afrontado efectivamente los cambios importantes que han estado ocurriendo en tu vida?")
	respuesta06 := ingresarRespuestaInvertida(
		"¿con qué frecuencia has estado seguro sobre tu capacidad para manejar tus problemas personales?")
	respuesta07 := ingresarRespuestaInvertida(
		"¿con qué frecuencia has sentido que las cosas van bien?")
	respuesta08 := ingresarRespuesta(
		"¿con qué frecuencia has sentido que no podías afrontar todas las cosas que tenías que hacer?")
	respuesta09 := ingresarRespuestaInvertida(
		"¿con qué frecuencia has podido controlar las dificultades de tu vida?")
	respuesta10 := ingresarRespuestaInvertida(
		"¿con qué frecuencia has sentido que tenías todo bajo control?")
	respuesta11 := ingresarRespuesta(
		"¿con qué  frecuencia has estado enfadado porque las cosas que te han ocurrido estaban fuera de tu control?")
	respuesta12 := ingresarRespuesta(
		"¿con qué frecuencia has pensado sobre las cosas que te faltan por hacer?")
	respuesta13 := ingresarRespuestaInvertida(
		"¿con qué frecuencia has podido controlar la forma de pasar el tiempo?")
	respuesta14 := ingresarRespuesta(
		"¿con qué frecuencia has sentido que las dificultades se acumulan tanto que no puedes superarlas?")

	puntajeTotal := calcularPuntajeTotal(respuesta01, respuesta02,
		respuesta03, respuesta04, respuesta05, respuesta06,
		respuesta07, respuesta08, respuesta09, respuesta10,
		respuesta11, respuesta12, respuesta13, respuesta14)
		
	nivelEstres := obtenerNivelEstres(puntajeTotal)

	reporteEstres := generarReporteEstres(nivelEstres, puntajeTotal)

	modulo.MostrarMensaje(reporteEstres)
}

func ingresarRespuesta(pregunta string) int {
	// Se restar uno para obtener una escala de 0 a 4
	return modulo.IngresarOpcion(
		fmt.Sprintf("\nEn el último mes, %s\n\n"+
			"\t(1) Nunca,\n"+
			"\t(2) Casi nunca,\n"+
			"\t(3) De vez en cuando\n"+
			"\t(4) A menudo\n"+
			"\t(5) Muy a menudo\n\n"+
			"\tCuál es su frecuencia: ", pregunta),
		OPCIONES) - 1
}

func ingresarRespuestaInvertida(pregunta string) int {
	// Se usa valor absoluto(abs) al restar 4 para tener la escala 4 a 0
	return int(math.Abs(float64(ingresarRespuesta(pregunta) - OPCIONES + 1)))
}

func calcularPuntajeTotal(respuesta01 int, respuesta02 int,
	respuesta03 int, respuesta04 int, respuesta05 int, respuesta06 int,
	respuesta07 int, respuesta08 int, respuesta09 int, respuesta10 int,
	respuesta11 int, respuesta12 int, respuesta13 int, respuesta14 int) int {
	return respuesta01 + respuesta02 + respuesta03 + respuesta04 +
		respuesta05 + respuesta06 + respuesta07 + respuesta08 + respuesta09 +
		respuesta10 + respuesta11 + respuesta12 + respuesta13 + respuesta14
}

func obtenerNivelEstres(puntajeTotal int) string {
	var nivelEstres string

	if puntajeTotal < NIVEL_BAJO {
		nivelEstres = "BAJO"
	} else {
		if puntajeTotal < NIVEL_MODERADO {
			nivelEstres = "MODERADO"
		} else {
			nivelEstres = "ALTO"
		}
	}

	return nivelEstres
}

func generarReporteEstres(nivelEstres string, puntajeTotal int) string {
	return fmt.Sprintf("\nSu nivel de estrés es %s, "+
		"con un puntaje de %d.\n", nivelEstres, puntajeTotal)
}
