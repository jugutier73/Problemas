/*
   Crear un programa para a clasificación de residuos según su
   material. Residuos aprovechable (limpios y secos), usar bolsa
   blanca; residuos orgánicos, usar bolsa verde; lo anterior si
   existe ruta de recolección selectiva; en otro caso, usar bolsa
   de color negro.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

final int APROVECHABLE = 1;
final int OTRO         = 3;

final int MAX_OPCIONES = 3;

void main() {
	var tipoResiduo = ingresarOpcion("Residuo \n"+
		"\t(1) aprovechable (limpia y seca),\n"+
		"\t(2) orgánico,\n"+
		"\t(3) otro\n\n"+
		"Cuál es el tipo de residuo: ", MAX_OPCIONES);

	var hayRecoleccionSelectiva = Util.ingresarLogico(
		"Hay ruta de recolección selectiva (s/n): ");

	var colorBolsa = recomendarColorBolsa(
		tipoResiduo, hayRecoleccionSelectiva);
		
	var reporteBolsaRecoleccion = generarReporteBolsa(colorBolsa);

	Util.mostrarMensaje(reporteBolsaRecoleccion);
}

int ingresarOpcion(String pregunta, int maximaOpcion) {
	var opcion = Util.ingresarEntero(pregunta);

	if (opcion < 1 || opcion > maximaOpcion ){
		Util.mostrarError("La opción no es válida, se asume 1\n\n");
		opcion = 1;
	}

	return opcion;
}

String recomendarColorBolsa(int tipoResiduo,
	boolean hayRecoleccionSelectiva) {
	String colorBolsa;

	if (tipoResiduo == OTRO || !hayRecoleccionSelectiva) {
		colorBolsa = "NEGRA";
	} else {
		if (tipoResiduo == APROVECHABLE) {
			colorBolsa = "BLANCA";
		} else {
			colorBolsa = "VERDE";
		}
	}

	return colorBolsa;
}

String generarReporteBolsa(String colorBolsa){
	return String.format(
		"\nLa bolsa recomendada es de color %s.\n", colorBolsa);
}