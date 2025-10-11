/*
    Crear un programa para recomendar un medio de transporte
    según el tipo de distancia a la universidad, si está lloviendo,
    y si hay o no transporte público.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Julio 2025
    Licencia: GNU GPL v3
*/

import modulo.Util;

final int CERQUITA = 1;
final int LEJOS    = 3;

void main() {
	var tipoDistancia = Util.ingresarOpcion("Vive\n"+
		"\t(1) cerquita,\n"+
		"\t(2) cerca,\n"+
		"\t(3) lejos: \n\n"+
		"Cuál es el tipo de distancia: ", 3);

	var estaLloviendo = Util.ingresarLogico("Está lloviendo (s/n): ");

	var hayTransporte = Util.ingresarLogico("Hay transporte público (s/n):");

	var medioTransporte = recomendarMedioTransporte(
		tipoDistancia, estaLloviendo, hayTransporte);

	var reporteTransporte = generarReporteTransporte(medioTransporte);

	Util.mostrarMensaje(reporteTransporte);
}

String recomendarMedioTransporte(int tipoDistancia,
	boolean estaLloviendo, boolean hayTransporte) {
	String medioTransporte;

	if (tipoDistancia == CERQUITA && !estaLloviendo) {
		medioTransporte = "caminar o usar bicicleta";
	} else {
		if (tipoDistancia == LEJOS || !hayTransporte) {
			medioTransporte = "carro compartido";
		} else {
			medioTransporte = "transporte público";
		}
	}

	return medioTransporte;
}

String generarReporteTransporte(String medioTransporte){
	return String.format(
		"\nMedio de transporte recomendado: %s.\n", medioTransporte);
}
