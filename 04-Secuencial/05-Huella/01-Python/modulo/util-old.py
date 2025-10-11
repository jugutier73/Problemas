"""
    Módulo de utilidades (funciones de apoyo)
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
"""


def mostrar_mensaje(mensaje):
    """Método para mostrar un mensaje (cadena) en la salida estandar.

    Args:
        mensaje: Mensaje que se desea mostrar.
    """
    print(mensaje, end="")


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
    return int(ingresar_texto(pregunta))


def ingresar_real(pregunta):
    """Devuelve el real ingresado por el usuario como respuesta a una 
       pregunta.

    Args:
        pregunta: Texto que se le presenta al usuario como pregunta.

    Returns:
        Valor ingresado por el usuario.
    """
    return float(ingresar_texto(pregunta))