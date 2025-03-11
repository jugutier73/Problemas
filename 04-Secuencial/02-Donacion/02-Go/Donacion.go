/*
   Programa para la iniciativa Amigo Social que permita generar
   el documento de recibido que sirva como soporte contable de
   la donación.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
*/

package main

import (
	"destacado/modulo"
	"fmt"
	"strconv"
)

func main() {
	nombreCasa := modulo.IngresarTexto("Nombre casa adulto mayor: ")
	cantidadRecolectada := ingresarEntero("Cantidad recolectada: ")
	reciboDonacion := generarRecibo(nombreCasa, cantidadRecolectada)
	modulo.MostrarMensaje(reciboDonacion)
}

func ingresarEntero(pregunta string) int {
	entero, _ := strconv.Atoi(modulo.IngresarTexto(pregunta))
	return entero
}

func generarRecibo(nombreCasa string,
	cantidadRecolectada int) string {
	return fmt.Sprintf(
		"\nLa iniciativa Amigo Social tiene el gusto de entregar"+
			"\nuna donación de $%d pesos colombianos"+
			"\na la casa del adulto mayor %s.", cantidadRecolectada, nombreCasa+
			"\n\n_______________________"+
			"\nFirma representante legal")
}
