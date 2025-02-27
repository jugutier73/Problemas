"""
    Programa para la empresa de seguridad Aros S.A.
    que permita resaltar el nombre del empleado destacado 
    del mes ofreciendo unas felicitaciones públicas.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Febrero 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje

def main():
    nombre_empleado = ingresar_texto("Nombre del empleado destacado: ")
    nombre_mes = ingresar_texto("Nombre del mes: ")
    felicitaciones = generar_felicitaciones(nombre_empleado, nombre_mes)
    mostrar_mensaje(felicitaciones)

def ingresar_texto(pregunta):
    return input(pregunta)

def generar_felicitaciones(nombre_empleado, nombre_mes):
    return (
        f"\nLa empresa de seguridad Aros S.A. quiere felicitar"
        f"\npúblicamente a {nombre_empleado} como nuestro"
        f"\nempleado destacado del mes de {nombre_mes},"
        f"\nmuchas gracias por su excelencia."
    )

main()