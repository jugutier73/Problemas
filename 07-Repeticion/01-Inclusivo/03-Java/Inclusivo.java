/*
   Crear un programa para contar cuantas veces se emplean símbolos 
   "x" o la "@" en un mensaje.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

final char SIMBOLO_INCLUSIVO_1 = 'x';
final char SIMBOLO_INCLUSIVO_2 = '@';

void main() {
	var texto = Util.ingresarTexto("Ingrese un texto con los mensajes: ");

	var cantidadSimbolos = contarEmpleoSimbolos(texto);

	var reporteSimbolos = generarReporteInclusivo(cantidadSimbolos);

	Util.mostrarMensaje(reporteSimbolos);
}

int contarEmpleoSimbolos(String texto) {
	var cantidadSimbolosInclusivos = 0;

	for (char caracter : texto.toCharArray() ){
		if (caracter == SIMBOLO_INCLUSIVO_1 || caracter ==SIMBOLO_INCLUSIVO_2){
			cantidadSimbolosInclusivos = cantidadSimbolosInclusivos + 1;
		}
	}

	return cantidadSimbolosInclusivos;
}

String generarReporteInclusivo(int cantidadSimbolosInclusivos) {
	return String.format("\nSe emplearon %d veces "+
		"los símbolos inclusivos \"%c\" y \"%c\".\n",
		cantidadSimbolosInclusivos, SIMBOLO_INCLUSIVO_1, SIMBOLO_INCLUSIVO_2);
}
