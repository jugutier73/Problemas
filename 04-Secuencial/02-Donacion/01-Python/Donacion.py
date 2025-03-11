"""
    Programa para la iniciativa Amigo Social que permita generar 
    el documento de recibido que sirva como soporte contable de 
    la donación.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Febrero 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_texto

def main():
    nombre_casa = ingresar_texto("Nombre casa adulto mayor: ")
    cantidad_recolectada = ingresar_texto("Cantidad recolectada: ")
    recibo = generar_recibo(nombre_casa, cantidad_recolectada)
    mostrar_mensaje(recibo)

def ingresar_entero(pregunta):
    return int(ingresar_texto(pregunta))

def generar_recibo(nombre_casa, cantidad_recolectada):
    return (
        f"\nLa iniciativa Amigo Social tiene el gusto de entregar"
        f"\nuna donación de ${cantidad_recolectada} pesos colombianos"
        f"\na la casa del adulto mayor {nombre_casa}."
        f"\n\n_______________________"
        f"\nFirma representante legal"
    )

main()