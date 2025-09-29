"""
    Programa para verificar la elegibilidad de los donantes 
    con base en la edad, el peso y la madurez fisiológica
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_entero, \
                        ingresar_real, ingresar_logico

EDAD_MINIMA = 18
EDAD_MAXIMA = 65
EDAD_MINIMA_AUTORIZACION = 16
EDAD_MAXIMA_AUTORIZACION = 18
PESO_MINIMO = 50.0

def main():
    edad_donante = ingresar_entero("Edad del donante: ")
    autorizacion_edad = ingresar_logico(
        "Tiene autorización por edad (s/n): ")
    peso_donante = ingresar_real("Peso del donante: ")
    suficiente_madurez = ingresar_logico(
        "Suficiente madurez fisiológica(s/n): ")

    reporte_eligibilidad = generar_reporte_eligibilidad(
        edad_donante, autorizacion_edad, peso_donante, suficiente_madurez)

    mostrar_mensaje(reporte_eligibilidad)

def generar_reporte_eligibilidad(
        edad_donante, autorizacion_edad, peso_donante, suficiente_madurez):
    mensaje = "\nEl donante no cumple las condiciones para donar\n" 

    if (edad_donante >= EDAD_MINIMA and  edad_donante <= EDAD_MAXIMA or \
        edad_donante >= EDAD_MINIMA_AUTORIZACION and \
        edad_donante <= EDAD_MAXIMA_AUTORIZACION and \
        autorizacion_edad) and \
        peso_donante > PESO_MINIMO and suficiente_madurez:
        mensaje = "\nEl donante es elegible para donar\n"
        
    return mensaje

main()