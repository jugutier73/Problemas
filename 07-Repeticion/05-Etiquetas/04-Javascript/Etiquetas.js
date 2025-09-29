/*
   Programa para verificar si una etiqueta de reciclaje es válida
   y reportar los campos con error.

   Formato: T-aaaa-mm
   - T: tipo de residuo
	 (P: plástico, V: vidrio, M: metal, C: cartón/papel, O: orgánico)
   - aaaa: año (> 2025)
   - mm: mes (01-12)

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

const LONGITUD_ETIQUETA = 9;

const SEPARADOR = '-';
const POSICION_SEPARADOR1 = 1;
const POSICION_SEPARADOR2 = 6;
const POSICION_TIPO = 0;
const POSICION_ANIO = 2;
const LONGITUD_ANIO = 4;
const POSICION_MES = 7;
const LONGITUD_MES = 2;

const MINIMO_ANIO = 2026;
const MINIMO_MES = 1;
const MAXIMO_MES = 12;

const TIPOS_VALIDOS = "PVMCO";

const ETIQUETA_OK = 0;
const ERROR_LONGITUD = 1;
const ERROR_FORMATO = 2;
const ERROR_TIPO = 4;
const ERROR_ANIO = 8;
const ERROR_MES = 16;

function main() {
	let etiqueta = ingresarTexto('etiqueta');

	let codigo = verificarEtiquetaReciclaje(etiqueta);

	let reporteEtiqueta = generarReporteEtiqueta(etiqueta, codigo);

	mostrarMensaje(reporteEtiqueta);
}

function verificarEtiquetaReciclaje(etiqueta) {
	let codigo = verificarLongitud(etiqueta);

	if (codigo != ERROR_LONGITUD) {
		codigo  = verificarFormato(etiqueta) |
		          verificarTipo(etiqueta)    |
				  		verificarAnio(etiqueta)    |
				  		verificarMes(etiqueta);
	}
	
	return codigo;
}

function verificarLongitud(etiqueta) {
	let codigoError = ETIQUETA_OK;

	if (etiqueta.length != LONGITUD_ETIQUETA) {
		codigoError = ERROR_LONGITUD;
	}

	return codigoError;
}

function verificarFormato(etiqueta) {
	let codigoError = ETIQUETA_OK;

	if (etiqueta[POSICION_SEPARADOR1] != SEPARADOR || 
	    etiqueta[POSICION_SEPARADOR2] != SEPARADOR) {
		codigoError = ERROR_FORMATO;
	}

	return codigoError;
}

function verificarTipo(etiqueta) {
	let codigoError = ETIQUETA_OK;

	if ( !TIPOS_VALIDOS.includes(etiqueta[POSICION_TIPO]) ) {
		codigoError = ERROR_TIPO;
	}

	return codigoError;
}

function verificarAnio(etiqueta) {
	let codigoError = ETIQUETA_OK;

	let anioEtiqueta = etiqueta.substring(POSICION_ANIO, POSICION_ANIO + LONGITUD_ANIO);
	let anio = Number(anioEtiqueta);

	if (Number.isInteger(anio)) {
		
		if (anio < MINIMO_ANIO) {
			codigoError =  ERROR_ANIO;
		}
	} else {
		codigoError =  ERROR_ANIO;
	}

	return codigoError;
}

function verificarMes(etiqueta) {
	let codigoError = ETIQUETA_OK;

	let mesEtiqueta = etiqueta.substring(POSICION_MES, POSICION_MES + LONGITUD_MES);
	let mes = Number(mesEtiqueta);

	if (Number.isInteger(mes)) {	
		if (mes < MINIMO_MES || mes > MAXIMO_MES) {
			codigoError =  ERROR_MES;
		}
	} else {
		codigoError =  ERROR_MES;
	}

	return codigoError;
}

function generarReporteEtiqueta(etiqueta, codigo) {
	let reporte = `\nLa  etiqueta "${etiqueta}" `;

	if (codigo === ETIQUETA_OK) {
		reporte += "es válida.\n";
	} else {
		reporte += "es inválida por tener:";

		if (codigo & ERROR_LONGITUD) {
			reporte += "\n  - Error en la longitud";
		}

		if (codigo & ERROR_FORMATO) {
			reporte += "\n  - Error en el formato";
		} 

		if (codigo & ERROR_TIPO) {
			reporte += "\n  - Error en el tipo";
		}

		if (codigo & ERROR_ANIO) {
			reporte += "\n  - Error en el año";
		}

		if (codigo & ERROR_MES) {
			reporte += "\n  - Error en el mes";
		}     
	}

	return reporte + "\n";
}
