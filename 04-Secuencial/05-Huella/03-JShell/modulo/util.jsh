/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
*/

void mostrarMensaje(String mensaje) {
  System.out.print(mensaje);
}

String ingresarTexto(String pregunta) {
  var scanner = new Scanner(System.in);

  mostrarMensaje(pregunta);
  return scanner.nextLine();
}

int ingresarEntero(String pregunta) {
  return Integer.parseInt(ingresarTexto(pregunta));
}

double ingresarReal(String pregunta) {
  return Double.parseDouble(ingresarTexto(pregunta));
}