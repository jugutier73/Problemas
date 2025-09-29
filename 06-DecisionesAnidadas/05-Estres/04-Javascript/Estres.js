/*
   Crear un programa para evaluar el nivel de estrés de los estudiantes

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Agosto 2025
   Licencia: GNU GPL v3
*/

const OPCIONES = 5;

const NIVEL_BAJO     = 19;
const NIVEL_MODERADO = 25;

function main() {
	let respuesta01 = ingresarRespuesta('respuesta01');
	let respuesta02 = ingresarRespuesta('respuesta02');
	let respuesta03 = ingresarRespuesta('respuesta03');
	let respuesta04 = ingresarRespuestaInvertida('respuesta04');
	let respuesta05 = ingresarRespuestaInvertida('respuesta05');
	let respuesta06 = ingresarRespuestaInvertida('respuesta06');
	let respuesta07 = ingresarRespuestaInvertida('respuesta07');
	let respuesta08 = ingresarRespuesta('respuesta08');
	let respuesta09 = ingresarRespuestaInvertida('respuesta09');
	let respuesta10 = ingresarRespuestaInvertida('respuesta10');
	let respuesta11 = ingresarRespuesta('respuesta11');
	let respuesta12 = ingresarRespuesta('respuesta12');
	let respuesta13 = ingresarRespuestaInvertida('respuesta13');
	let respuesta14 = ingresarRespuesta('respuesta14');

	let puntajeTotal = calcularPuntajeTotal(respuesta01, respuesta02,
		respuesta03, respuesta04, respuesta05, respuesta06, respuesta07,
		respuesta08, respuesta09, respuesta10, respuesta11, respuesta12,
		respuesta13, respuesta14)
	let nivelEstres = obtenerNivelEstres(puntajeTotal);

	let reporteEstres = generarReporteEstres(nivelEstres, puntajeTotal);

	mostrarMensaje(reporteEstres);
}

function ingresarRespuesta(componente) {
	let opcion;

	if (ingresarLogico(componente + '_1')) {
		opcion = 0;
	} else if (ingresarLogico(componente + '_2')) {
		opcion = 1;
	} else if (ingresarLogico(componente + '_3')) {
		opcion = 2;
	}else if (ingresarLogico(componente + '_4')) {
		opcion = 3;
	}else if (ingresarLogico(componente + '_5')) {
		opcion = 4;
	}

	return opcion;
}

function ingresarRespuestaInvertida(pregunta) {
	// Se usa valor absoluto(abs) al restar 4 para tener la escala 4 a 0
	return Math.abs(ingresarRespuesta(pregunta) - OPCIONES + 1);
}

function calcularPuntajeTotal(respuesta01, respuesta02, respuesta03, 
	respuesta04, respuesta05, respuesta06, respuesta07, respuesta08, 
	respuesta09, respuesta10, respuesta11, respuesta12, respuesta13, 
	respuesta14) {
	return respuesta01 + respuesta02 + respuesta03 + respuesta04 +
		respuesta05 + respuesta06 + respuesta07 + respuesta08 +
		respuesta09 + respuesta10 + respuesta11 + respuesta12 +
		respuesta13 + respuesta14;
}

function obtenerNivelEstres(puntajeTotal) {
	let nivelEstres;

	if (puntajeTotal < NIVEL_BAJO) {
		nivelEstres = "BAJO";
	} else {
		if (puntajeTotal < NIVEL_MODERADO) {
			nivelEstres = "MODERADO";
		} else {
			nivelEstres = "ALTO";
		}
	}

	return nivelEstres;
}

function generarReporteEstres(nivelEstres, puntajeTotal) {
	return `\nSu nivel de estrés es ${nivelEstres}, ` +
		   `con un puntaje de ${puntajeTotal}.\n`;
}
