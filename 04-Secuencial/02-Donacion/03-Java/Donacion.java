/*
   Programa para la iniciativa Amigo Social que permita generar
   el documento de recibido que sirva como soporte contable de
   la donación.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
*/
import modulo.Util;

void main() {
    var nombreCasa = Util.ingresarTexto("Nombre casa adulto mayor: ");
	  var cantidadRecolectada = ingresarEntero("Cantidad recolectada: ");

	  var reciboDonacion = generarRecibo(nombreCasa, cantidadRecolectada);
    
	  Util.mostrarMensaje(reciboDonacion);
}

int ingresarEntero(String pregunta) {
  return Integer.parseInt(Util.ingresarTexto(pregunta));
}

String generarRecibo(String nombreCasa, int cantidadRecolectada) {
	return String.format(
    "\nLa iniciativa Amigo Social tiene el gusto de entregar"+
		"\nuna donación de $%d pesos colombianos"+
		"\na la casa del adulto mayor %s.\n", 
    cantidadRecolectada, nombreCasa);
}
