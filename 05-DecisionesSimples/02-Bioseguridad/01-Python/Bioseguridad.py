"""
    Crear un programa para programa para la verificación de medidas 
    de bioseguridad para el ingreso a un evento público.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_texto, mostrar_error

def main():
    tiene_vacunas = ingresar_logico("Tiene vacunas (s/n): ")
    resultados_negativos = ingresar_logico("Pruebas negativas (s/n): ")
    tiene_sintomas = ingresar_logico("Tiene síntomas (s/n): ")

    reporte_ingreso = generar_reporte_ingreso(tiene_vacunas, 
                                              resultados_negativos, 
                                              tiene_sintomas)

    mostrar_mensaje(reporte_ingreso)

def ingresar_logico(pregunta):
    opcion = ingresar_texto(pregunta).lower()

    if opcion != "s" and opcion != "n":
        mostrar_error("No es una opción válida, se asume \"n\"\n\n")

    return opcion == "s"

def generar_reporte_ingreso(
        tiene_vacunas, resultados_negativos, tiene_sintomas):
    mensaje = "\nLa persona no puede ingresar al evento\n" 

    if tiene_vacunas and resultados_negativos and not tiene_sintomas:
        mensaje = "\nLa persona puede ingresar al evento\n"

    return mensaje

main()