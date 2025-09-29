"""
    Crear un programa para a clasificación de residuos según su 
    material. Residuos aprovechable (limpios y secos), usar bolsa 
    blanca; residuos orgánicos, usar bolsa verde; lo anterior si 
    existe ruta de recolección selectiva; en otro caso, usar bolsa 
    de color negro.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Julio 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_entero, \
                        ingresar_logico, mostrar_error

APROVECHABLE = 1
OTRO         = 3

MAX_OPCIONES = 3

def main():
    tipo_residuo = ingresar_opcion("Residuo \n"+
                                   "\t(1) aprovechable (limpia y seca),\n"+
                                   "\t(2) orgánico,\n"+
                                   "\t(3) otro\n\n"+
                                   "Cuál es el tipo de residuo: ", MAX_OPCIONES)
    hay_recoleccion_selectiva = ingresar_logico(
        "Hay ruta de recolección selectiva (s/n): ")
    
    color_bolsa = recomendar_color_bolsa(
        tipo_residuo, hay_recoleccion_selectiva)
    reporte_bolsa_recoleccion = generar_reporte_bolsa(color_bolsa)

    mostrar_mensaje(reporte_bolsa_recoleccion)

def ingresar_opcion(pregunta, maxima_opcion):
    opcion = ingresar_entero(pregunta)

    if opcion < 1 or opcion > maxima_opcion:
        mostrar_error("La opción no es válida, se asume 1\n\n")
        opcion = 1

    return opcion 

def recomendar_color_bolsa(tipo_residuo, hay_recoleccion_selectiva):
    if tipo_residuo == OTRO or not hay_recoleccion_selectiva:
        color_bolsa = "NEGRA"
    else:
        if tipo_residuo == APROVECHABLE:
            color_bolsa = "BLANCA"
        else:
            color_bolsa = "VERDE"

    return color_bolsa

def generar_reporte_bolsa(color_bolsa):
    return f"\nLa bolsa recomendada es de color {color_bolsa}.\n"

main()