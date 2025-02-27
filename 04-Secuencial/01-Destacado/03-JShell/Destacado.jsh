/*
   Programa para la empresa de seguridad Aros S.A.
   que permita resaltar el nombre del empleado destacado
   del mes ofreciendo unas felicitaciones públicas.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
*/

void main() {
    var nombreEmpleado = ingresarTexto("Nombre del empleado destacado: ");
	  var nombreMes = ingresarTexto("Nombre del mes: ");
	  var felicitaciones = generarFelicitaciones(nombreEmpleado, nombreMes);
	  mostrarMensaje(felicitaciones);

    System.exit(0);
}

String ingresarTexto(String pregunta) {
  var scanner = new Scanner(System.in);

  mostrarMensaje(pregunta);
  return scanner.nextLine();
}

String generarFelicitaciones(String nombreEmpleado, String nombreMes) {
	return String.format(
		"\nLa empresa de seguridad Aros S.A. quiere felicitar"+
		"\npúblicamente a %s como nuestro"+
		"\nempleado destacado del mes de %s,"+
		"\nmuchas gracias por su excelencia.", nombreEmpleado, nombreMes);
}

main();