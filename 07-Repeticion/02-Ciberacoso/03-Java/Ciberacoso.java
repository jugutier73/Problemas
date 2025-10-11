/*
   Crear un programa para contar cuantas palabras comienzan con
   la letra "p" en un mensaje.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

final char LETRA_INICIO = 'p';

void main() {
	var texto = Util.ingresarTexto("Ingrese un texto con mensajes a analizar:");

	var cantidadPalabras = contarPalabrasInician(texto, LETRA_INICIO);

	var reporteCiberacoso = generarReporteAcoso(cantidadPalabras);

	Util.mostrarMensaje(reporteCiberacoso);
}

int contarPalabrasInician(String texto, char letraInicio) {
	var palabras = texto.toLowerCase().split(" ");

	var cantidadPalabrasInteres = 0;

	for (var palabra : palabras) {
		  if (palabra.startsWith(Character.toString(letraInicio))) {
			  cantidadPalabrasInteres++;
		  }
    }

	return cantidadPalabrasInteres;
}

String generarReporteAcoso(int cantidadPalabrasInteres ) {
	return String.format("\nHay %d palabras "+
		"que inician con la letra \"%c\".\n",
		cantidadPalabrasInteres, LETRA_INICIO);
}
