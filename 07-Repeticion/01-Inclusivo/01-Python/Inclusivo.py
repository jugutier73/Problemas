"""
    Crear un programa para contar cuantas veces se emplean símbolos 
    "x" o la "@" en un mensaje.

    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Septiembre 2025
    Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, mostrar_mensaje

SIMBOLO_INCLUSIVO_1 = "x"
SIMBOLO_INCLUSIVO_2 = "@"

def main():
    texto = ingresar_texto("Ingrese un texto con los mensajes: ")
    
    cantidad_simbolos = contar_empleo_simbolos(texto)
    
    reporte_simbolos = generar_reporte_inclusivo(cantidad_simbolos)

    mostrar_mensaje(reporte_simbolos)

def contar_empleo_simbolos(texto):
    cantidad_simbolos_inclusivos = 0

    for caracter in texto:
        if caracter == SIMBOLO_INCLUSIVO_1 or \
           caracter == SIMBOLO_INCLUSIVO_2:
            cantidad_simbolos_inclusivos = cantidad_simbolos_inclusivos + 1

    return cantidad_simbolos_inclusivos

def generar_reporte_inclusivo(cantidad_simbolos_inclusivos):
    return (f"\nSe emplearon {cantidad_simbolos_inclusivos} veces "
            f"los símbolos inclusivos \"{SIMBOLO_INCLUSIVO_1}\" y "
            f"\"{SIMBOLO_INCLUSIVO_2}\".\n"
    )

main()


