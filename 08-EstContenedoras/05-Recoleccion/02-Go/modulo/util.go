/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

package modulo

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// MostrarMensaje mostrar un mensaje (cadena) en la salida estandar.
func MostrarMensaje(mensaje string) {
	fmt.Print(mensaje)
}

// MostrarError mostrar un mensaje (cadena) en el error estandar.
func MostrarError(mensaje string) {
	fmt.Fprint(os.Stderr, mensaje)
}

// IngresarTexto devuelve el texto ingresado por el usuario como respuesta a una pregunta.
func IngresarTexto(pregunta string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(pregunta)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

// IngresarEntero devuelve el entero ingresado por el usuario como respuesta a una pregunta.
func IngresarEntero(pregunta string) int {
	entero, err := strconv.Atoi(IngresarTexto(pregunta))

	if err != nil {
		MostrarError("El valor ingresado es inválido, intente nuevamente.\n\n")
		entero = IngresarEntero(pregunta)
	}

	return entero
}

// IngresarReal devuelve el real ingresado por el usuario como respuesta a una pregunta.
func IngresarReal(pregunta string) float64 {
	real, err := strconv.ParseFloat(IngresarTexto(pregunta), 64)

	if err != nil {
		MostrarError("El valor ingresado es inválido, intente nuevamente.\n\n")
		real = IngresarReal(pregunta)
	}

	return real
}

// IngresarLogico devuelve el booleano ingresado por el usuario como respuesta a una pregunta.
func IngresarLogico(pregunta string) bool {
	for {
		opcion := strings.ToLower(IngresarTexto(pregunta))

		if opcion == "s" || opcion == "n" {
			return opcion == "s"
		}

		MostrarError("No es una opción válida, intente nuevamente.\n\n")
	}
}

// IngresarOpcion devuelve el entero ingresado por el usuario como respuesta a una pregunta con múltiples opciones del 1 a un máximo especificado.
func IngresarOpcion(pregunta string, maximaOpcion int) int {
	opcion := IngresarEntero(pregunta)

	if opcion < 1 || opcion > maximaOpcion {
		MostrarError("La opción no es válida, intente nuevamente.\n\n")
		opcion = IngresarOpcion(pregunta, maximaOpcion)
	}

	return opcion
}

// IngresarColeccion devuelve una colección, con los elementos obtenidos de un función especificada, hasta que el usuario indique que ya no sea ingresar más datos.
func IngresarColeccion[Tipo any](ingresarElemento func() Tipo) []Tipo {
	var coleccion []Tipo
	hayMas := true

	for hayMas {
		coleccion = append(coleccion, ingresarElemento())

		hayMas = IngresarLogico("Hay más datos (s/n): ")
	}

	return coleccion
}

// OrdenarColeccion ordena una coleccion de elementos, según una función que hace las veces de comparador y orden indicado.
func OrdenarColeccion[Tipo any](
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

// ConvertirColeccionCadena convierte una coleccion de elementos a una cadena, según una función dada como argumento que convierte un elemento.
func ConvertirColeccionCadena[Tipo any](titulo string, coleccion []Tipo,
	convertirElementoCadena func(elemento Tipo) string) string {
	mensaje := "\n" + titulo + "\n"

	for _, elemento := range coleccion {
		mensaje += "\t" + convertirElementoCadena(elemento)
	}

	return mensaje
}

// ContarSegunCriterio cuenta la cantidad de elementos de una colección que cumplen un criterio dado como argumento.
func ContarSegunCriterio[Tipo1 any, Tipo2 any](coleccion []Tipo1, aplicarCriterio func(Tipo1, Tipo2) bool, valorCriterio Tipo2) int {
	cantidad := 0

	for _, elemento := range coleccion {
		if aplicarCriterio(elemento, valorCriterio) {
			cantidad += 1
		}
	}

	return cantidad
}
