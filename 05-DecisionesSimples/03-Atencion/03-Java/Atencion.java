/*
   Programa para establecer la prioridad de atención
   en los centros de salud para identificar a los pacientes con
   mayor riesgo o vulnerabilidad
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/
import modulo.Util;

final int EDAD_RECIEN_NACIDO = 1;
final int EDAD_ADULTO_MAYOR  = 60;

void main() {
	var edadPaciente = Util.ingresarEntero("Edad del paciente: ");
	var enfermedadCronica = Util.ingresarLogico("Enfermedad crónica (s/n):");
	var estadoInmunosupresion = Util.ingresarLogico("Estado de inmunosupresión (s/n): ");

	var reporteAtencion = generarReporteAtencion(edadPaciente,
		enfermedadCronica, estadoInmunosupresion);

	Util.mostrarMensaje(reporteAtencion);
}

String generarReporteAtencion(
	int edadPaciente, 
	boolean enfermedadCronica, 
	boolean estadoInmunosupresion) {

	var mensaje = "\nEl paciente es de atención general\n";

	if (edadPaciente < EDAD_RECIEN_NACIDO ||
		edadPaciente > EDAD_ADULTO_MAYOR ||
		enfermedadCronica || 
		estadoInmunosupresion) {
		mensaje = "\nEl paciente es de atención prioritaria\n";
	}
	
	return mensaje;
}
