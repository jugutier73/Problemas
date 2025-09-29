package modulo;

/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
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
   * @return Valor ingresado por el usuario o cero si es un valor inválido.
   */
  public static int ingresarEntero(String pregunta) {
    int valor;
    try {
      valor = Integer.parseInt(ingresarTexto(pregunta));
    } catch (NumberFormatException e) {
      mostrarError("El valor ingresado es inválido, se asume 0\n\n");
      valor = 0;
    }
    return valor;
  }


  /**
   * Devuelve el real ingresado por el usuario como respuesta a una pregunta.
   * 
   * @param pregunta Texto que se le presenta al usuario como pregunta.
   * @return Valor ingresado por el usuario o cero si es un valor inválido.
   */
  public static double ingresarReal(String pregunta) {
    double valor;
    try {
      valor = Double.parseDouble(ingresarTexto(pregunta));
    } catch (NumberFormatException e) {
      mostrarError("El valor ingresado es inválido, se asume 0.0\n\n");
      valor = 0.0;
    }
    return valor;
  }

}