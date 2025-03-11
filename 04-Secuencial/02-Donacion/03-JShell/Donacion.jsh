/*
   Programa para la iniciativa Amigo Social que permita generar
   el documento de recibido que sirva como soporte contable de
   la donación.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
*/

void main() {
    var nombreCasa = ingresarTexto("Nombre casa adulto mayor: ");
	  var cantidadRecolectada = ingresarEntero("Cantidad recolectada: ");
	  var reciboDonacion = generarRecibo(nombreCasa, cantidadRecolectada);
	  mostrarMensaje(reciboDonacion);

    System.exit(0);
}

int ingresarEntero(String pregunta) {
  return Integer.parseInt(ingresarTexto(pregunta));
}

String generarRecibo(String nombreCasa, int cantidadRecolectada) {
	return String.format(
    "\nLa iniciativa Amigo Social tiene el gusto de entregar"+
		"\nuna donación de $%d pesos colombianos"+
		"\na la casa del adulto mayor %s.", 
    cantidadRecolectada, nombreCasa);
}

main();