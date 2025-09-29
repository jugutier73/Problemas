"""
    Crear un programa para contar cuantas palabras comienzan con 
    la letra "p" en un mensaje.

    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Septiembre 2025
    Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, mostrar_mensaje

LETRA_INICIO = "p"

def main():
    texto = ingresar_texto("Ingrese un texto con los mensajes a analizar:")
    
    cantidad_palabras = contar_palabras_inician(texto, LETRA_INICIO)
    
    reporte_simbolos = generar_reporte_acoso(cantidad_palabras)

    mostrar_mensaje(reporte_simbolos)

def contar_palabras_inician(texto, letra_inicio):
    palabras = texto.lower().split()
    cantidad_palabras_interes = 0
    
    for palabra in palabras:
        if palabra.startswith(letra_inicio):
            cantidad_palabras_interes += 1
    
    return cantidad_palabras_interes

def generar_reporte_acoso(cantidad_palabras_interes):
    return (f"\nHay {cantidad_palabras_interes} palabras "
            f"que inician con la letra \"{LETRA_INICIO}\"\n"
        )

main()