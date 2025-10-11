/*
   Crear un programa para registrar y consultar de manera ágil
   todas las donaciones recibidas. Esta debería ofrecer el total
   recaudado, listados alfabéticos y descendentes por valor,
   identificar a los aportantes que superen un umbral específico y
   señalar al mayor donante.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"donaciones/modulo"
	"fmt"
	"sort"
)

const (
	UMBRAL = 10000
)

type Donante struct {
	Nombre   string
	Donacion int
}

func main() {
	donantes := ingresarColeccion(ingresarDonante)

	donantesPorNombre := ordenarColeccion(donantes, compararNombre, false)
	donantesPorDonacion := ordenarColeccion(donantes, compararDonacion, true)

	mayoresDonantes := obtenerMayoresDonaciones(donantes, UMBRAL)
	mayorDonante := obtenerMayorDonante(donantes)
	sumaDonaciones := obtenerSumaDonaciones(donantes)

	reporteDonaciones := generarReporteDonaciones(
		donantesPorNombre,
		donantesPorDonacion,
		mayoresDonantes,
		mayorDonante,
		sumaDonaciones)

	modulo.MostrarMensaje(reporteDonaciones)
}

func ingresarColeccion[Tipo any](ingresarElemento func() Tipo) []Tipo {
	var coleccion []Tipo
	hayMas := true

	for hayMas {
		coleccion = append(coleccion, ingresarElemento())

		hayMas = modulo.IngresarLogico("Hay más datos (s/n): ")
	}

	return coleccion
}

func ingresarDonante() Donante {
	modulo.MostrarMensaje("\nIngrese los datos de un donante:\n")
	nombre := modulo.IngresarTexto("\tIngrese el nombre: ")
	donacion := modulo.IngresarEntero("\tIngrese donación : ")

	return Donante{Nombre: nombre, Donacion: donacion}
}

func ordenarColeccion[Tipo any](
	coleccion []Tipo,
	comparador func(elemento1 Tipo, elemento2 Tipo) bool,
	descendente bool) []Tipo {
	copiaColeccion := make([]Tipo, len(coleccion))
	copy(copiaColeccion, coleccion)

	sort.SliceStable(copiaColeccion, func(i, j int) bool {
		if descendente {
			return comparador(copiaColeccion[i], copiaColeccion[j])
		} else {
			return comparador(copiaColeccion[j], copiaColeccion[i])
		}
	})

	return copiaColeccion
}

func compararNombre(donante1 Donante, donante2 Donante) bool {
	return donante1.Nombre > donante2.Nombre
}

func compararDonacion(donante1 Donante, donante2 Donante) bool {
	return donante1.Donacion > donante2.Donacion
}

func obtenerMayoresDonaciones(donantes []Donante, valorLimite int) []Donante {
	mayoresDonaciones := []Donante{}

	for _, donante := range donantes {
		if donante.Donacion > valorLimite {
			mayoresDonaciones = append(mayoresDonaciones, donante)
		}
	}

	return mayoresDonaciones
}

func obtenerMayorDonante(donantes []Donante) Donante {
	var mayorDonante Donante

	primerDonante := true
	for _, donacion := range donantes {
		if primerDonante || donacion.Donacion > mayorDonante.Donacion {
			mayorDonante = donacion
			primerDonante = false
		}
	}

	return mayorDonante
}

func obtenerSumaDonaciones(coleccion []Donante) int {
	sumaDonaciones := 0

	for _, donacion := range coleccion {
		sumaDonaciones += donacion.Donacion
	}

	return sumaDonaciones
}

func generarReporteDonaciones(donantesOrdenNombre,
	donantesOrdenadas,
	donantesMayores []Donante,
	mayorDonante Donante,
	sumaDonaciones int) string {
	listadoPorNombre := convertirColeccionCadena(
		"LISTADO EN ORDEN ALFABÉTICO",
		donantesOrdenNombre,
		convertirDonanteCadena)

	listadoPorDonacion := convertirColeccionCadena(
		"LISTADO ORDENADO POR DONACIÓN",
		donantesOrdenadas,
		convertirDonanteCadena)

	listadoMayores := convertirColeccionCadena(
		fmt.Sprintf("LISTADO DONACIONES MAYORES A $%d", UMBRAL),
		donantesMayores,
		convertirDonanteCadena)

	return fmt.Sprintf("%s\n%s\n%s\n"+
		"El mayor donante: %s\n"+
		"Total de donaciones $%d\n",
		listadoPorNombre,
		listadoPorDonacion,
		listadoMayores,
		mayorDonante.Nombre,
		sumaDonaciones)
}

func convertirColeccionCadena[Tipo any](titulo string, coleccion []Tipo,
	convertirElementoCadena func(elemento Tipo) string) string {
	mensaje := "\n" + titulo + "\n"

	for _, elemento := range coleccion {
		mensaje += "\t" + convertirElementoCadena(elemento)
	}

	return mensaje
}

func convertirDonanteCadena(donante Donante) string {
	return fmt.Sprintf("$%10d \t %s\n", donante.Donacion, donante.Nombre)
}
