"""
    Crear un programa para recomendar un medio de transporte
    según el tipo de distancia a la universidad, si está lloviendo,
    y si hay o no transporte público.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Julio 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_opcion, \
                        ingresar_logico

CERQUITA = 1
LEJOS    = 3

def main():
    tipo_distancia = ingresar_opcion("Vive\n"+
                                     "\t(1) cerquita,\n"+
                                     "\t(2) cerca,\n"+
                                     "\t(3) lejos: \n\n"+
                                     "Cuál es el tipo de distancia: ", 3)
    esta_lloviendo = ingresar_logico("Está lloviendo (s/n): ")
    hay_transporte = ingresar_logico("Hay transporte público (s/n): ")  
    
    medio_transporte = recomendar_medio_transporte(
        tipo_distancia, esta_lloviendo, hay_transporte)
    reporte_transporte = generar_reporte_transporte(medio_transporte)

    mostrar_mensaje(reporte_transporte)

def recomendar_medio_transporte(tipo_distancia, esta_lloviendo, hay_transporte):
    if tipo_distancia == CERQUITA and not esta_lloviendo:
        medio_transporte = "caminar o usar bicicleta"
    else:
        if tipo_distancia == LEJOS or not hay_transporte:
            medio_transporte = "carro compartido"
        else:
            medio_transporte = "transporte público"

    return medio_transporte

def generar_reporte_transporte(medio_transporte):
    return f"\nMedio de transporte recomendado: {medio_transporte}.\n"

main()