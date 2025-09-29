/*
   Crear un programa para registrar síntomas y recibir recomendaciones 
   simples y seguras como reposar, hidratarse o consultar a un profesional 
   de la salud, según la gravedad o combinación de síntomas reportados.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

const FIEBRE_ALTA = 39.5;
const FIEBRE      = 37.5;

const MENSAJE_ESPECIALISTA = `
 - Consultar un especialista.
 - Anotar los síntomas y cuándo comenzaron.
 - Evitar esfuerzos físicos y actividades intensas.`;

const MENSAJE_FIEBRE_ALTA = `
 - Solicitar una cita medica con urgencia.
 - Usar paños húmedos y fríos en la frente.
 - Permanecer en un lugar fresco y ventilado.`;

const MENSAJE_FIEBRE = `
 - Descansar lo suficiente.
 - Hidratarse bebiendo agua u otros líquidos.`;

const MENSAJE_DOLOR_CABEZA = `
 - Realizar ejercicio cervical isométrico.
 - Descansar en un lugar oscuro y silencioso.
 - Beber agua, ya que la deshidratación puede empeorar el dolor.`;

const MENSAJE_CONGESTION = `
 - Realizar lavados nasales con solución salina.
 - Usar almohadas extras para dormir con la cabeza elevada.`;
	
function main() {
	let temperatura = ingresarReal('temperatura');
	let sintomasVariosDias = ingresarLogico('sintomasVariosDias');
	let malestarIntenso = ingresarLogico('malestarIntenso');
	let dolorCabeza = ingresarLogico('dolorCabeza');
	let congestionNasal = ingresarLogico('congestionNasal');

	let recomendacionesEspecialista = recibirRecomendaciones(
		sintomasVariosDias || malestarIntenso, MENSAJE_ESPECIALISTA);

	let recomendacionesFiebreAlta = recibirRecomendaciones(
		temperatura >= FIEBRE_ALTA, MENSAJE_FIEBRE_ALTA);

	let recomendacionesFiebre = recibirRecomendaciones(
		temperatura >= FIEBRE, MENSAJE_FIEBRE);
		
	let recomendacionesDolorCabeza = recibirRecomendaciones(
		dolorCabeza, MENSAJE_DOLOR_CABEZA);
		
	let recomendacionesCongestion = recibirRecomendaciones(
		congestionNasal, MENSAJE_CONGESTION);

	let recomendaciones = generarReporteRecomendaciones(
		recomendacionesEspecialista,
		recomendacionesFiebreAlta,
		recomendacionesFiebre,
		recomendacionesDolorCabeza,
		recomendacionesCongestion);

	mostrarMensaje(recomendaciones);
}


function recibirRecomendaciones(condicion, recomendaciones) {
	return condicion ? recomendaciones : "";
}

function generarReporteRecomendaciones(
	recomendacionesEspecialista,
	recomendacionesFiebreFuerte,
	recomendacionesFiebre,
	recomendacionesDolorCabeza,
	recomendacionesCongestion) {
	let reporteRecomendaciones = "\nSe recomienda:";

	if (recomendacionesEspecialista != "") {
		reporteRecomendaciones += recomendacionesEspecialista;
	} else {
		reporteRecomendaciones += recomendacionesFiebreFuerte +
			recomendacionesFiebre + recomendacionesDolorCabeza +
			recomendacionesCongestion;
	}

	return reporteRecomendaciones + "\n";
}