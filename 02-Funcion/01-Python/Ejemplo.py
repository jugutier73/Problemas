"""
    Programa que muestra varios mensaje en pantalla
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Enero 2025
    Licencia: GNU GPL v3
"""


def main():
    mostrar_mensaje("Hola ")
    mostrar_mensaje("Mundo\n")
    mostrar_mensaje("Bienvenidos\t Al lenguaje Python\n")


def mostrar_mensaje(mensaje):
    print(mensaje, end="")


main()