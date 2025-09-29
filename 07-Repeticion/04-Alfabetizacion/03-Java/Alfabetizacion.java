/*
   Crear un programa para determinar si hay problemas con el uso
   de los espacios en un frase que el usuario ingrese

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

void main() {
	var frase = Util.ingresarTexto("Ingrese una frase a analizar: ");

	var fraseCorregida = corregirUsoEspacios(frase);

	var reporte = generarReporteAlfabetizacion(frase, fraseCorregida);

	Util.mostrarMensaje(reporte);
}

String corregirUsoEspacios(String frase) {
	var inicioFrase = false;
	var espacioPendiente = false;
	var fraseCorregida = "";

	for (var simbolo : frase.toCharArray()) {
		if (simbolo == ' ' ){
			if (inicioFrase) {
				espacioPendiente = true;
			}
		} else {
			if (espacioPendiente) {
				fraseCorregida += " ";
				espacioPendiente = false;
			}
			fraseCorregida += simbolo;
			inicioFrase = true;
		}
	}

	return fraseCorregida;
}

String generarReporteAlfabetizacion(String frase, String fraseCorregida ) {
	var reporteAlfabetizacion = String.format("\nLa frase \"%s\" hace un uso ", frase);

	if (frase.equals(fraseCorregida)) {
		reporteAlfabetizacion += "CORRECTO de espacios.";
	} else {
		reporteAlfabetizacion += String.format("INCORRECTO de espacios,\n"+
			"lo correcto es \"%s\"", fraseCorregida);
	}

	return reporteAlfabetizacion + "\n";
}
