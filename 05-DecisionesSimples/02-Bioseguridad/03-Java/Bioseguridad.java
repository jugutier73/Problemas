/*
   Crear un programa para programa para la verificación de medidas
   de bioseguridad para el ingreso a un evento público.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

void main() {
	var tieneVacunas = ingresarLogico("Tiene vacunas (s/n): ");
	var restadosNegativos = ingresarLogico("Pruebas negativas (s/n): ");
	var tieneSintomas = ingresarLogico("Tiene síntomas (s/n): ");

	var reporteIngreso = generarReporteIngreso(
		tieneVacunas, restadosNegativos, tieneSintomas);

	Util.mostrarMensaje(reporteIngreso);
}

boolean ingresarLogico(String pregunta) {
	var opcion = Util.ingresarTexto(pregunta).toLowerCase();

	if ( !opcion.equals("s") && !opcion.equals("n")) {
		Util.mostrarError("No es una opción válida, se asume \"n\"\n\n");
	}

	return opcion.equals("s");
}

String generarReporteIngreso(
	boolean tieneVacunas, boolean restadosNegativos, boolean tieneSintomas) {
	var mensaje = "\nLa persona no puede ingresar al evento\n";

	if (tieneVacunas && restadosNegativos && !tieneSintomas) {
		mensaje = "\nLa persona puede ingresar al evento\n";
	}
	
	return mensaje;
}
