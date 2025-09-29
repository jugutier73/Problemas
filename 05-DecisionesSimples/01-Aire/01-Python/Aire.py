"""
    Crear un programa para mostrar una alerta cuando el índice de la 
    calidad del aire (medido con algún instrumento) indique que el aire
    puede resultar perjudicial para la población.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_real

def main():
    indice_calidad_aire = ingresar_real("Índice calidad del aire: ")

    alerta_calidad_aire = generar_alerta_calidad_aire(indice_calidad_aire)

    mostrar_mensaje(alerta_calidad_aire)

def generar_alerta_calidad_aire(indice_calidad_aire):
    mensaje = "\nEl aire supone un riesgo bajo para la salud\n"

    if indice_calidad_aire > 100.0:
        mensaje = "\nEl aire puede presentar efectos sobre la salud\n"
        
    return mensaje

main()