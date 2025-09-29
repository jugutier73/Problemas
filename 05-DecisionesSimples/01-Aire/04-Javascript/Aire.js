/*
   Crear un programa para mostrar una alerta cuando el índice de la
   calidad del aire (medido con algún instrumento) indique que el aire
   puede resultar perjudicial para la población.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

function main() {
	let indiceCalidadAire = ingresarReal("indiceCalidadAire");

	let alertaCalidadAire = generarAlertaCalidadAire(indiceCalidadAire);

	mostrarMensaje(alertaCalidadAire);
}

function generarAlertaCalidadAire(indiceCalidadAire ) {
	let mensaje = "\nEl aire supone un riesgo bajo para la salud\n";

	if (indiceCalidadAire > 100.0) {
		mensaje = "\nEl aire puede presentar efectos sobre la salud\n";
	}
	
	return mensaje;
}
