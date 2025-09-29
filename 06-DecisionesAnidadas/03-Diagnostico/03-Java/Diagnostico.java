/*
   Crear un programa para registrar síntomas y recibir recomendaciones 
   simples y seguras como reposar, hidratarse o consultar a un profesional 
   de la salud, según la gravedad o combinación de síntomas reportados.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

final double FIEBRE_ALTA = 39.5;
final double FIEBRE      = 37.5;

final String MENSAJE_ESPECIALISTA = """

 - Consultar un especialista.
 - Anotar los síntomas y cuándo comenzaron.
 - Evitar esfuerzos físicos y actividades intensas.""";

final String MENSAJE_FIEBRE_ALTA = """

 - Solicitar una cita medica con urgencia.
 - Usar paños húmedos y fríos en la frente.
 - Permanecer en un lugar fresco y ventilado.""";

final String MENSAJE_FIEBRE = """

 - Descansar lo suficiente.
 - Hidratarse bebiendo agua u otros líquidos.""";

final String MENSAJE_DOLOR_CABEZA = """

 - Realizar ejercicio cervical isométrico.
 - Descansar en un lugar oscuro y silencioso.
 - Beber agua, ya que la deshidratación puede empeorar el dolor.""";

final String MENSAJE_CONGESTION = """

 - Realizar lavados nasales con solución salina.
 - Usar almohadas extras para dormir con la cabeza elevada.""";

void main() {
	var temperatura = Util.ingresarReal("Temperatura corporal: ");
	var sintomasVariosDias = Util.ingresarLogico("Síntomas por 2 o más días (s/n): ");
	var malestarIntenso = Util.ingresarLogico("Tiene malestar intenso (s/n): ");
	var dolorCabeza = Util.ingresarLogico("Tiene dolor de cabeza (s/n): ");
	var congestionNasal = Util.ingresarLogico("Tiene congestion nasal (s/n): ");

	var recomendacionesEspecialista = recibirRecomendaciones
		(sintomasVariosDias || malestarIntenso,	MENSAJE_ESPECIALISTA);

	var recomendacionesFiebreAlta = recibirRecomendaciones
		(temperatura >= FIEBRE_ALTA, MENSAJE_FIEBRE_ALTA);

	var recomendacionesFiebre = recibirRecomendaciones
		(temperatura >= FIEBRE, MENSAJE_FIEBRE);

	var recomendacionesDolorCabeza = recibirRecomendaciones
		(dolorCabeza, MENSAJE_DOLOR_CABEZA);
		
	var recomendacionesCongestion = recibirRecomendaciones
		(congestionNasal, MENSAJE_CONGESTION);

	var recomendaciones = generarReporteRecomendaciones(
							recomendacionesEspecialista, recomendacionesFiebreAlta,  
							recomendacionesFiebre,   	 recomendacionesDolorCabeza, 
							recomendacionesCongestion);
	
	Util.mostrarMensaje(recomendaciones);
}

String recibirRecomendaciones(boolean condicion, String recomendaciones) {
	return condicion ? recomendaciones : "";
}

String generarReporteRecomendaciones(String recomendacionesEspecialista,
									 String recomendacionesFiebreFuerte, 
									 String recomendacionesFiebre,
									 String recomendacionesDolorCabeza, 
									 String recomendacionesCongestion) {
	var reporteRecomendaciones = "\nSe recomienda:";

	if (recomendacionesEspecialista != "") {
		reporteRecomendaciones += recomendacionesEspecialista;
	} else {
		reporteRecomendaciones += recomendacionesFiebreFuerte +
			recomendacionesFiebre + recomendacionesDolorCabeza +
			recomendacionesCongestion;
	}

	return reporteRecomendaciones + "\n";
}