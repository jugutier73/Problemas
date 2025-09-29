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
    return Integer.parseInt(ingresarTexto(pregunta));
  }

}