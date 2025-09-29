/*
   Crear un programa para contar cuántas letras mayúsculas,
   cuántas minúsculas y cuántos dígitos contiene una contraseña

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"fmt"
	"seguridad/modulo"
	"unicode"
)

func main() {
	contrasenia := modulo.IngresarTexto("Ingrese la contraseña a analizar: ")

	cantidadMayusculas := contarMayusculas(contrasenia)
	cantidadMinusculas := contarMinusculas(contrasenia)
	cantidadDigitos := contarDigitos(contrasenia)

	reporteContrasenia := generarReporteContrasenia(contrasenia,
		cantidadMayusculas, cantidadMinusculas, cantidadDigitos)

	modulo.MostrarMensaje(reporteContrasenia)
}

func contarMayusculas(texto string) int {
	cantidadMayusculas := 0

	for _, caracter := range texto {
		if unicode.IsUpper(caracter) {
			cantidadMayusculas++
		}
	}

	return cantidadMayusculas
}

func contarMinusculas(texto string) int {
	cantidadMinusculas := 0

	for _, caracter := range texto {
		if unicode.IsLower(caracter) {
			cantidadMinusculas++
		}
	}

	return cantidadMinusculas
}

func contarDigitos(texto string) int {
	cantidadDigitas := 0

	for _, caracter := range texto {
		if unicode.IsDigit(caracter) {
			cantidadDigitas++
		}
	}

	return cantidadDigitas
}

func generarReporteContrasenia(contrasenia string,
	cantidadMayusculas int, cantidadMinusculas int, cantidadDigitos int) string {
	return fmt.Sprintf("\nEn la constraseña \"%s\" hay:\n"+
		"%d mayúscula(s), "+
		"%d minúscula(s) y "+
		"%d dígito(s)\n",
		contrasenia, cantidadMayusculas, cantidadMinusculas, cantidadDigitos)
}
