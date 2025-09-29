/*
    Programa para verificar la elegibilidad de los donantes 
    con base en la edad, el peso y la madurez fisiológica
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
*/

import modulo.Util;

final int EDAD_MINIMA              = 18;
final int EDAD_MAXIMA              = 65;
final int EDAD_MINIMA_AUTORIZACION = 16;
final int EDAD_MAXIMA_AUTORIZACION = 18;
final double PESO_MINIMO           = 50.0;

void main() {
	var edadDonante = Util.ingresarEntero("Edad del donante: ");
	var autorizacionEdad = Util.ingresarLogico("Tiene autorización edad (s/n): ");
	var pesoDonante = Util.ingresarReal("Peso del donante: ");
	var suficienteMadurez = Util.ingresarLogico(
		"Suficiente madurez fisiológica(s/n): ");

	var reporteLegibilidad = generarReporteEligibilidad(
		edadDonante, autorizacionEdad, pesoDonante, suficienteMadurez);

	Util.mostrarMensaje(reporteLegibilidad);
}

String generarReporteEligibilidad(
	int edadDonante,    
	boolean autorizacionEdad,  
	double pesoDonante,	
	boolean suficienteMadurez) {
	var mensaje = "\nEl donante no cumple las condiciones para donar\n";

	if ((edadDonante >= EDAD_MINIMA && edadDonante <= EDAD_MAXIMA ||
		edadDonante >= EDAD_MINIMA_AUTORIZACION &&
			edadDonante <= EDAD_MAXIMA_AUTORIZACION && autorizacionEdad) &&
		pesoDonante > PESO_MINIMO && suficienteMadurez) {
		mensaje = "\nEl donante es elegible para donar\n";
	}
	
	return mensaje;
}
