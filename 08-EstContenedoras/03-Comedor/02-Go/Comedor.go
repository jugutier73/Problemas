/*
   Crear un programa para organizar la entrada de los beneficiarios
   y garantice la atención en orden de llegada. Se debe registrar
   datos clave (nombre, edad y presencia de necesidades especiales).
   Además de informar la cantidad de personas con necesidades y el
   promedio de edades.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"comedor/modulo"
	"fmt"
)

type Reserva struct {
	Nombre            string
	Edad              int
	NecesidadEspecial bool
}

func main() {
	reservasComedor := modulo.IngresarColeccion(ingresarReverva)

	cantidadConNecesidades := modulo.ContarSegunCriterio(reservasComedor, tenerNecesidadEspecial, true)

	promedioEdades := calcularPromedioEdades(reservasComedor)

	reporteReservas := generarReporteReservasComedor(
		reservasComedor,
		cantidadConNecesidades,
		promedioEdades)

	modulo.MostrarMensaje(reporteReservas)
}

func ingresarReverva() Reserva {
	modulo.MostrarMensaje("\nIngrese los datos de la reserva:\n")
	nombre := modulo.IngresarTexto("\tIngrese el nombre de la persona   : ")
	edad := modulo.IngresarEntero("\tIngrese la edad de la persona     : ")
	necesidadEspecial := modulo.IngresarLogico("\tTiene necesidades especiales (s/n): ")

	return Reserva{Nombre: nombre, Edad: edad, NecesidadEspecial: necesidadEspecial}
}

func tenerNecesidadEspecial(reserva Reserva, necesidadEspecialInteres bool) bool {
	return reserva.NecesidadEspecial == necesidadEspecialInteres
}

func calcularPromedioEdades(reservasComedor []Reserva) float64 {
	totalReservas := len(reservasComedor)
	promedioEdades := 0.0
	sumaEdades := 0

	for _, reserva := range reservasComedor {
		sumaEdades += reserva.Edad
	}

	if totalReservas > 0 {
		promedioEdades = float64(sumaEdades) / float64(totalReservas)
	}

	return promedioEdades
}

func generarReporteReservasComedor(
	reservasComedor []Reserva,
	cantidadConNecesidades int,
	promedioEdades float64) string {

	listadoReservas := modulo.ConvertirColeccionCadena(
		"LISTADO DE RESERVAS",
		reservasComedor,
		convertirReservaCadena)

	return fmt.Sprintf("%s\n\n"+
		"Cantidad con necesidades : %d\n"+
		"Promedio de edades       : %0.1f\n",
		listadoReservas,
		cantidadConNecesidades,
		promedioEdades)
}

func convertirReservaCadena(reserva Reserva) string {
	mensaje := fmt.Sprintf("%s, %d años", reserva.Nombre, reserva.Edad)

	if reserva.NecesidadEspecial {
		mensaje += ", con necesidad especial"
	}

	return mensaje + "\n"
}
