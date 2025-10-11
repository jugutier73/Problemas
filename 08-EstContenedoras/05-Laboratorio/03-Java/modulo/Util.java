package modulo;

/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.Scanner;
import java.util.function.BiPredicate;
import java.util.function.Function;
import java.util.function.Supplier;

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
   * Devuelve el texto ingresado por el usuario como respuesta a una 
   * pregunta.
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
   * Devuelve el entero ingresado por el usuario como respuesta a una 
   * pregunta.
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
   * Devuelve el real ingresado por el usuario como respuesta a una 
   * pregunta.
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
   * Devuelve el booleano ingresado por el usuario como respuesta a una 
   * pregunta.
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
   * Devuelve el entero ingresado por el usuario como respuesta a una 
   * pregunta con múltiples opciones del 1 a un máximo especificado.
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


  /**
   * Devuelve una colección, con los elementos obtenidos de un función
   * especificada, hasta que el usuario indique que ya no sea ingresar más 
   * datos.
   * 
   * @param <Tipo> Tipo de dato de un elemento.
   * @param ingresarElemento Función que solicita un dato al usuario.
   * @return Coleccion con los valores ingresados por medio de la función 
   *         indicada.
   */
  public static <Tipo> List<Tipo> ingresarColeccion(Supplier<Tipo> ingresarElemento){
          var coleccion = new ArrayList<Tipo>();
          var hayMas = true;

          while (hayMas){
              coleccion.add(ingresarElemento.get());

              hayMas = ingresarLogico( "Hay más datos (s/n): ");
          }

          return coleccion;
  }


  /**
   * Ordena una coleccion de elementos, según el campo y orden indicado.
   * 
   * @param <Tipo> Tipo de dato de un elemento.
   * @param coleccion Colección de elementos que se desean ordenar.
   * @param comparador Función utilizada como criterio de ordenación.
   * @param descendente Valor booleano para indicar si se ordena
   *        descendentemente.
   * @return Nueva colección con todos los datos ordenados por el criterio
   *         y sentido indicado.
   */
  public static <Tipo> List<Tipo>  ordenarColeccion(List<Tipo> coleccion, 
                                      Comparator<Tipo> comparador, 
                                      boolean descendente){
    return coleccion.stream()
      .sorted(descendente ? comparador.reversed() : comparador)
      .toList();
  }


  /**
   * Convierte una coleccion de elementos a una cadena, según una función 
   * dada como argumento que convierte un elemento.
   * 
   * @param <Tipo> Tipo de dato de un elemento.
   * @param titulo Título que se utiliza para los elementos de la colección
   * @param coleccion Colección de elementos a mostrar.
   * @param convertirElementoCadena Función que convierte un elemento a una
   *        cadena.
   * @return Un texto con un título y todos los elementos de la colección.
   */
  public static <Tipo> String convertirColeccionCadena(String titulo, 
                                        List<Tipo> coleccion, 
                                        Function<Tipo,String> convertirElementoCadena){
          var mensaje = "\n" + titulo + "\n";

          for (var elemento : coleccion) {
                  mensaje += "\t" + convertirElementoCadena.apply(elemento);
          }

          return mensaje;
  }


  /**
   * Cuenta la cantidad de elementos de una colección que cumplen un 
   * criterio dado como argumento.
   * 
   * @param <Tipo1> Tipo de dato de un elemento.
   * @param <Tipo2> Tipo de dato del valor de referencia.
   * @param coleccion Colección de elementos a procesar.
   * @param aplicarCriterio Función que toma un elemento e indica si cumple 
   *        un cierto criterio.
   * @param valorCriterio  Valor de refencia para el criterio.
   * @return Cantidad de elementos de la colección que cumplemen el 
   *         el criterio indicado.
   */
  public static <Tipo1, Tipo2> int contarSegunCriterio(List<Tipo1> coleccion, BiPredicate<Tipo1, Tipo2> aplicarCriterio, Tipo2 valorCriterio){
    return (int) coleccion.stream()
      .filter(elemento -> aplicarCriterio.test(elemento, valorCriterio))
      .count();
  }

}