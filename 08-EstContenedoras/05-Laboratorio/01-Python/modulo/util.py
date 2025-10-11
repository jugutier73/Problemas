"""
    Módulo de utilidades (funciones de apoyo)
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Septiembre 2025
    Licencia: GNU GPL v3
"""

import sys


def mostrar_mensaje(mensaje):
    """Muestra un mensaje (cadena) en la salida estandar.

    Args:
        mensaje: Mensaje que se desea mostrar.
    """
    print(mensaje, end="")


def mostrar_error(mensaje):
    """Muestra un mensaje (cadena) en el error estandar.

    Args:
        mensaje: Mensaje que se desea mostrar como error.
    """
    print(mensaje, end="", file=sys.stderr)


def ingresar_texto(pregunta):
    """Devuelve el texto ingresado por el usuario como respuesta 
       a una pregunta.

    Args:
        pregunta: Texto que se le presenta al usuario como pregunta.

    Returns:
        Texto ingresado por el usuario.
    """
    return input(pregunta)


def ingresar_entero(pregunta):
    """Devuelve el entero ingresado por el usuario como respuesta a una
       pregunta.

    Args:
        pregunta: Texto que se le presenta al usuario como pregunta.

    Returns:
        Valor ingresado por el usuario.
    """
    try:
        valor = int(ingresar_texto(pregunta))
    except ValueError:
        mostrar_error("El valor ingresado es inválido, intente nuevamente.\n\n")
        valor = ingresar_entero(pregunta)
    return valor


def ingresar_real(pregunta):
    """Devuelve el real ingresado por el usuario como respuesta a una 
       pregunta.

    Args:
        pregunta: Texto que se le presenta al usuario como pregunta.

    Returns:
        Valor ingresado por el usuario.
    """
    try:
        valor = float(ingresar_texto(pregunta))
    except ValueError:
        mostrar_error("El valor ingresado es inválido, intente nuevamente.\n\n")
        valor = ingresar_real(pregunta)
    return valor


def ingresar_logico(pregunta):
    """Devuelve el booleano ingresado por el usuario como respuesta a una
       pregunta.

    Args:
        pregunta: Texto que se le presenta al usuario como pregunta.

    Returns:
        Valor ingresado por el usuario.
    """     
    while True:
        opcion = ingresar_texto(pregunta).strip().lower()
        
        if opcion in ("s", "n"):
            return opcion == "s"
        
        mostrar_error("No es una opción válida, intente nuevamente.\n\n")


def ingresar_opcion(pregunta, maxima_opcion):
    """Devuelve el entero ingresado por el usuario como respuesta a una
       pregunta con múltiples opciones del 1 a un máximo especificado.

    Args:
        pregunta: Texto que se le presenta al usuario como pregunta.
        maxima_opcion: Entero que indica la cantidad de opciones.

    Returns:
        Valor ingresado por el usuario.
    """
    opcion = ingresar_entero(pregunta)

    if opcion < 1 or opcion > maxima_opcion:
        mostrar_error(f"La opción no es válida, intente nuevamente.\n\n")
        opcion = ingresar_opcion(pregunta, maxima_opcion)

    return opcion 


def ingresar_coleccion(ingresar_elemento):
    """Devuelve una colección, con los elementos obtenidos de un función
       especificada, hasta que el usuario indique que ya no sea ingresar 
       más datos.

    Args:
        ingresar_elemento: Función que solicita un dato al usuario.

    Returns:
        Coleccion con los valores ingresados por medio de la función indicada.
    """    
    coleccion = []
    hay_mas = True

    while hay_mas:
        coleccion.append(ingresar_elemento())

        hay_mas = ingresar_logico("Hay más datos (s/n): ")
        
    return coleccion


def ordenar_coleccion(coleccion, campo_seleccionado, descendente):
    """Ordena una coleccion de elementos, según el campo y orden indicado.

    Args:
        coleccion: Colección de elementos que se desean ordenar.
        campo_seleccionado: Campo empleado para ordenar.
        descendente: Valor booleano para indicar si se ordena descendentemente.

    Returns:
        Nueva colección con todos los datos ordenados por el campo y sentido indicado.
    """       
    return sorted(coleccion, key=campo_seleccionado, reverse=descendente)


def convertir_coleccion_cadena(titulo, coleccion, convertir_elemento_cadena):
    """Convierte una coleccion de elementos a una cadena, según una función
       dada como argumento que convierte un elemento.

    Args:
        titulo: Título que se utiliza para los elementos de la colección.
        coleccion: Colección de elementos a mostrar.
        convertir_elemento_cadena: Función que convierte un elemento a una cadena.

    Returns:
        Un texto con un título y todos los elementos de la colección.
    """ 
    mensaje = f"\n{titulo}\n"

    for elemento in coleccion:
        mensaje += "\t" + convertir_elemento_cadena(elemento)

    return mensaje


def contar_segun_criterio(coleccion, aplicar_criterio, valor_criterio):
    """Cuenta la cantidad de elementos de una colección que cumplen un
       criterio dado como argumento.

    Args:
        coleccion: Colección de elementos a procesar.
        aplicar_criterio: Función que toma un elemento e indica si cumple 
                          un cierto criterio.
        valor_criterio: Valor de refencia para el criterio.

    Returns:
        Cantidad de elementos de la colección que cumplemen el 
        el criterio indicado.
    """     
    cantidad = 0

    for elemento in coleccion:
        if aplicar_criterio(elemento, valor_criterio):
            cantidad += 1
    
    return cantidad