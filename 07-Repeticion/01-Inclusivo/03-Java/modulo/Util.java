package modulo;

/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

import java.util.Scanner;

public class Util {
  
  /**
   * Muestra un mensaje (cadena) en la salida estandar.
   * 
   * @param mensaje Mensaje que se desea mostrar.
   */
  public static void mostrarMensaje(String mensaje) {
    System.out.print(mensaje);
  }


  /**
   * Muestra un mensaje (cadena) en el error estandar.
   * 
   * @param mensaje Mensaje que se desea mostrar como error.
   */
  public static void mostrarError(String mensaje) {
    System.err.print(mensaje);
  }


  /**
   * Devuelve el texto ingresado por el usuario como respuesta a una pregunta.
   * 
   * @param pregunta Texto que se le presenta al usuario como pregunta.
   * @return Texto ingresado por el usuario
   */
  public static String ingresarTexto(String pregunta) {
    var scanner = new Scanner(System.in);
    
    mostrarMensaje(pregunta);

    return scanner.nextLine();
  }


  /**
   * Devuelve el entero ingresado por el usuario como respuesta a una pregunta.
   * 
   * @param pregunta Texto que se le presenta al usuario como pregunta.
   * @return Valor ingresado por el usuario.
   */
  public static int ingresarEntero(String pregunta) {
    int valor;
    try {
      valor = Integer.parseInt(ingresarTexto(pregunta));
    } catch (NumberFormatException e) {
      mostrarError("El valor ingresado es inválido, intente nuevamente.\n\n");
      valor = ingresarEntero(pregunta);
    }
    return valor;
  }


  /**
   * Devuelve el real ingresado por el usuario como respuesta a una pregunta.
   * 
   * @param pregunta Texto que se le presenta al usuario como pregunta.
   * @return Valor ingresado por el usuario.
   */
  public static double ingresarReal(String pregunta) {
    double valor;
    try {
      valor = Double.parseDouble(ingresarTexto(pregunta));
    } catch (NumberFormatException e) {
      mostrarError("El valor ingresado es inválido, intente nuevamente.\n\n");
      valor = ingresarReal(pregunta);
    }
    return valor;
  }


  /**
   * Devuelve el booleano ingresado por el usuario como respuesta a una pregunta.
   * 
   * @param pregunta Texto que se le presenta al usuario como pregunta.
   * @return Valor ingresado por el usuario.
   */
  public static boolean ingresarLogico(String pregunta) {
    String opcion;

    while (true) {
      opcion = ingresarTexto(pregunta).toLowerCase();

      if ( opcion.equals("s") || opcion.equals("n")) {
            return opcion.equals("s");  
        }
        mostrarError("No es una opción válida, intente nuevamente.\n\n");
    }
    
  }


  /**
   * Devuelve el entero ingresado por el usuario como respuesta a una pregunta con múltiples opciones del 1 a un máximo especificado.
   * 
   * @param pregunta Texto que se le presenta al usuario como pregunta.
   * @param maximaOpcion Entero que indica la cantidad de opciones.
   * @return Valor ingresado por el usuario.
   */
  public static int ingresarOpcion(String pregunta, int maximaOpcion) {
      var opcion = ingresarEntero(pregunta);

      if (opcion < 1 || opcion > maximaOpcion ){
        mostrarError("La opción no es válida, intente nuevamente.\n\n");
        opcion = ingresarOpcion(pregunta, maximaOpcion);
      }

      return opcion;
  }

}