/*
   Crear un programa para garantizar que los pacientes sean atendidos
   de manera equitativa y oportuna. El orden de llegada debe
   equilibrarse con la prioridad médica de cada caso

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/
import modulo.Util;

final int NIVELES = 3;

final int NIVEL_RIESGO      = 1;
final int NIVEL_URGENCIA    = 2;
final int NIVEL_PRIORITARIO = 3;

record Cita(String nombre, int nivel) {}

void main() {
	var citas = Util.ingresarColeccion(() -> ingresarCita());
	
	var citasPorNivel = Util.ordenarColeccion(citas, 
	    seleccionarCampoNivel(),   false);

	var cantidadRiesgo      = contarSegunCriterio(
		citas, (c, nivel) -> tenerNivelRiesgo(c, nivel), NIVEL_RIESGO);

	var cantidadUrgencia    = contarSegunCriterio(
		citas, (c, nivel) -> tenerNivelRiesgo(c, nivel), NIVEL_URGENCIA);

	var cantidadPrioritario = contarSegunCriterio(
		citas, (c, nivel) -> tenerNivelRiesgo(c, nivel), NIVEL_PRIORITARIO);

	var reporteCitas = generarReporteCitas(citasPorNivel,
										   cantidadRiesgo,
										   cantidadUrgencia,
										   cantidadPrioritario);

	Util.mostrarMensaje(reporteCitas);

}

Cita ingresarCita(){
	Util.mostrarMensaje("\nIngrese los datos de la cita:\n");
	var nombre = Util.ingresarTexto("\tIngrese el nombre paciente: ");
	var nivel  = Util.ingresarOpcion("\tNIVEL DE URGENCIA\n"+
		"\t\t1: Riesgo\n"+
		"\t\t2: Urgencia\n"+
		"\t\t3: Prioritario\n"+
		"\tIngrese tipo de persona: ",
		NIVELES);

	return new Cita(nombre, nivel);
}

Comparator<Cita> seleccionarCampoNivel() {
    return Comparator.comparingInt(Cita::nivel);
}

<Tipo1, Tipo2> int contarSegunCriterio(List<Tipo1> coleccion,
 BiPredicate<Tipo1, Tipo2> aplicarCriterio, Tipo2 valorCriterio){
	return (int) coleccion.stream()
		.filter(elemento -> aplicarCriterio.test(elemento, valorCriterio))
		.count();
}

boolean tenerNivelRiesgo(Cita cita, int nivelRiesgo )  {
	return cita.nivel() == nivelRiesgo;
}

String generarReporteCitas(List<Cita> citasPorNivel,
						   int cantidadRiesgo,
						   int cantidadUrgencia,
						   int cantidadPrioritario){
	var listadoPorNivel = Util.convertirColeccionCadena(
			"LISTADO POR NIVEL DE URGENCIA", 
			citasPorNivel,
			(cita) -> convertirCitaCadena(cita));

	return listadoPorNivel + "\n" 
			+ "Nivel 1 (Riesgo)     : " + cantidadRiesgo + "\n" 
			+ "Nivel 2 (Urdencia)   : " + cantidadUrgencia + "\n"
			+ "Nivel 3 (Prioritario): " + cantidadPrioritario + "\n";
}

String convertirCitaCadena(Cita cita){
	return String.format("%d : %s%n", cita.nivel(), cita.nombre());
}
