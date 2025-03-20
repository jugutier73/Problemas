"""
    Módulo de utilidades (funciones de apoyo)
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Febrero 2025
    Licencia: GNU GPL v3
"""

def mostrar_mensaje(mensaje):
    print(mensaje, end="")

def ingresar_texto(pregunta):
    return input(pregunta)

def ingresar_entero(pregunta):
    return int(ingresar_texto(pregunta))