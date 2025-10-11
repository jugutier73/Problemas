"""
    Módulo de utilidades (funciones de apoyo)
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Septiembre 2025
    Licencia: GNU GPL v3
"""

import sys


def mostrar_mensaje(mensaje):
    """Método para mostrar un mensaje (cadena) en la salida estandar.

    Args:
        mensaje: Mensaje que se desea mostrar.
    """
    print(mensaje, end="")


def mostrar_error(mensaje):
    """Método para mostrar un mensaje (cadena) en el error estandar.

    Args:
        mensaje: Mensaje que se desea mostrar como error.
    """
    print(mensaje, end="", file=sys.stderr)


def ingresar_texto(pregunta):
    """Devuelve el texto ingresado por el usuario como respuesta a una
       pregunta.

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