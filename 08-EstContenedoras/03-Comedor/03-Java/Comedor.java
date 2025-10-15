/*
   Crear un programa para organizar la entrada de los beneficiarios
   y garantice la atención en orden de llegada. Se debe registrar
   datos clave (nombre, edad y presencia de necesidades especiales).
   Además de informar la cantidad de personas con necesidades y el
   promedio de edades.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

record Reserva(String nombre, int edad, boolean necesidadEspecial) {}

void main() {
	var reservasComedor = Util.ingresarColeccion(() -> ingresarReserva());
	
	var cantidadConNecesidades = Util.contarSegunCriterio(
		reservasComedor, (reserva, necesidadEspecial) -> 
	    tenerNecesidadesEspeciales(reserva, necesidadEspecial), true );

	var promedioEdades = calcularPromedioEdades(reservasComedor);

	var reporteReservas = generarReporteReservasComedor(
		reservasComedor, cantidadConNecesidades, promedioEdades);

	Util.mostrarMensaje(reporteReservas);
}

Reserva ingresarReserva(){
	Util.mostrarMensaje("\nIngrese los datos de la reserva:\n");
	var nombre = Util.ingresarTexto("\tIngrese el nombre de la persona   : ");
	var edad  = Util.ingresarEntero("\tIngrese la edad de la persona     : ");
	var necesidadEspecial  = Util.ingresarLogico("\tTiene necesidades especiales (s/n): ");

	return new Reserva(nombre, edad, necesidadEspecial);
}

boolean tenerNecesidadesEspeciales(Reserva reservas, boolean necesidadEspecial){
	return reservas.necesidadEspecial() == necesidadEspecial;
}

double calcularPromedioEdades(List<Reserva> reservasComedor){
	return reservasComedor.stream()
		.mapToInt(Reserva::edad)
        .average()
        .orElse(0.0);
}

String generarReporteReservasComedor(
		List<Reserva> reservasComedor,
		int cantidadConNecesidades,
		double promedioEdades){
	var listadoReservas = Util.convertirColeccionCadena(
			"LISTADO DE RESERVAS", 
			reservasComedor,
			(reserva) -> convertirReservaCadena(reserva));

	return String.format("%s\n"
			+ "Cantidad con necesidades : %d\n" 
			+ "Promedio de edades       : %.1f\n",
			listadoReservas, cantidadConNecesidades, promedioEdades);
}

String convertirReservaCadena(Reserva reserva){
	var mensaje = String.format("%s, %d años", reserva.nombre(), reserva.edad());

	if (reserva.necesidadEspecial()) {
		mensaje +=  ", con necesidad especial";
	}

	return mensaje + "\n";
}
