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

import modulo.Util;

final int LONGITUD_ETIQUETA = 9;

final char SEPARADOR = '-';
final int POSICION_SEPARADOR1 = 1;
final int POSICION_SEPARADOR2 = 6;

final int POSICION_TIPO = 0;

final int POSICION_ANIO = 2;
final int LONGITUD_ANIO = 4;

final int POSICION_MES = 7;
final int LONGITUD_MES = 2;

final int MINIMO_ANIO = 2026;
final int MINIMO_MES = 1;
final int MAXIMO_MES = 12;

final String TIPOS_VALIDOS = "PVMCO";

final int ETIQUETA_OK    = 0;
final int ERROR_LONGITUD = 1;
final int ERROR_FORMATO  = 2;
final int ERROR_TIPO     = 4;
final int ERROR_ANIO     = 8;
final int ERROR_MES      = 16;

void main() {
	var etiqueta = Util.ingresarTexto("Ingrese la etiqueta a analizar: ");  

	var codigo = verificarEtiquetaReciclaje(etiqueta);

	var reporteEtiqueta = generarReporteEtiqueta(etiqueta, codigo);

	Util.mostrarMensaje(reporteEtiqueta);
}

int verificarEtiquetaReciclaje(String etiqueta) {
	int codigo = verificarLongitud(etiqueta);
	if (codigo != ERROR_LONGITUD) {
		codigo  = verificarFormato(etiqueta)  |
		          verificarTipo(etiqueta)     |
				  		verificarAnio(etiqueta)     |
				  		verificarMes(etiqueta);
	}

	return codigo;
}

int verificarLongitud(String etiqueta) {
	var codigoError = ETIQUETA_OK;

	if (etiqueta.length() != LONGITUD_ETIQUETA ) {
		codigoError = ERROR_LONGITUD;
	}

	return codigoError;
}

int verificarFormato(String etiqueta) {
	var codigoError = ETIQUETA_OK;

	if (etiqueta.charAt(POSICION_SEPARADOR1) != SEPARADOR || 
	    etiqueta.charAt(POSICION_SEPARADOR2) != SEPARADOR) {
		codigoError = ERROR_FORMATO;
	}

	return codigoError;
}

int verificarTipo(String etiqueta) {
	var codigoError = ETIQUETA_OK;

	var tipo = etiqueta.charAt(POSICION_TIPO);

	if (!TIPOS_VALIDOS.contains(Character.toString(tipo)) ) {
		codigoError = ERROR_TIPO;
	}

	return codigoError;
}

int verificarAnio(String etiqueta) {
	var codigoError = ETIQUETA_OK;

	var anioEtiqueta = etiqueta.substring(POSICION_ANIO, POSICION_ANIO + LONGITUD_ANIO);

	if (anioEtiqueta.chars().allMatch(Character::isDigit)) {
		int anio = Integer.parseInt(anioEtiqueta);
		if (anio < MINIMO_ANIO) {
			codigoError = ERROR_ANIO;
		}
	} else {
		codigoError = ERROR_ANIO;
	}

	return codigoError;
}

int verificarMes(String etiqueta) {
	var codigoError = ETIQUETA_OK;

	var mesEtiqueta = etiqueta.substring(POSICION_MES, POSICION_MES + LONGITUD_MES);

	if (mesEtiqueta.chars().allMatch(Character::isDigit)) {
		var mes = Integer.parseInt(mesEtiqueta);
		if (mes < MINIMO_MES || mes > MAXIMO_MES) {
			codigoError =  ERROR_MES;
		}
	} else {
		codigoError =  ERROR_MES;
	}

	return codigoError;
}

String generarReporteEtiqueta(String etiqueta, int codigo) {
	var reporte = String.format("\nLa  etiqueta \"%s\" ", etiqueta );

	if (codigo == ETIQUETA_OK) {
		reporte += "es válida.\n";
	} else {
		reporte += "es inválida por tener:";

		if ((codigo & ERROR_LONGITUD) != 0) { 
			reporte += "\n  - Error en la longitud"; 
		}
		
		if ((codigo & ERROR_FORMATO) != 0)  { 
			reporte += "\n  - Error en el formato";  
		}

		if ((codigo & ERROR_TIPO) != 0)     { 
			reporte += "\n  - Error en el tipo";     
		}

		if ((codigo & ERROR_ANIO) != 0)     { 
			reporte += "\n  - Error en el año";      
		}

		if ((codigo & ERROR_MES) != 0)      { 
			reporte += "\n  - Error en el mes";      
		}
	}

	return reporte + "\n";
}
