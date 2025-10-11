/*
    Crear un programa para recomendar un medio de transporte
    según el tipo de distancia a la universidad, si está lloviendo,
    y si hay o no transporte público.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Julio 2025
    Licencia: GNU GPL v3
*/

const CERQUITA = 1;
const LEJOS    = 3;

function main() {
	let tipoDistancia = ingresarOpcion('tipoDistancia');
	let estaLloviendo = ingresarLogico('estaLloviendo');
	let hayTransporte = ingresarLogico('hayTransporte');

	let medioTransporte = recomendarMedioTransporte(
		tipoDistancia, estaLloviendo, hayTransporte);
		
	let reporteTransporte = generarReporteTransporte(medioTransporte);

	mostrarMensaje(reporteTransporte);
}

function recomendarMedioTransporte(tipoDistancia,
	estaLloviendo, hayTransporte) {
	let medioTransporte;

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

function generarReporteTransporte(medioTransporte){
	return `\nMedio de transporte recomendado: ${medioTransporte}.\n`;
}