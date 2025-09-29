/*
   Crear un programa para mostrar una alerta cuando el índice de la
   calidad del aire (medido con algún instrumento) indique que el aire
   puede resultar perjudicial para la población.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

void main() {
	var	indiceCalidadAire = Util.ingresarReal("Índice calidad del aire: ");

	var alertaCalidadAire = generarAlertaCalidadAire(indiceCalidadAire);

	Util.mostrarMensaje(alertaCalidadAire);
}

String generarAlertaCalidadAire(double indiceCalidadAire) {
	var mensaje = "\nEl aire supone un riesgo bajo para la salud\n";
	
	if (indiceCalidadAire > 100.0) {
		mensaje = "\nEl aire puede presentar efectos sobre la salud\n";
	}
	
	return mensaje;
}
