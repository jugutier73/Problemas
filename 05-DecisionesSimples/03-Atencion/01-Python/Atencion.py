"""
    Programa para establecer la prioridad de atención 
    en los centros de salud para identificar a los pacientes con 
    mayor riesgo o vulnerabilidad
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_entero, ingresar_logico

EDAD_RECIEN_NACIDO = 1
EDAD_ADULTO_MAYOR = 60

def main():
    edad_paciente = ingresar_entero("Edad del paciente: ")
    enfermedad_cronica = ingresar_logico("Enfermedad crónica (s/n): ")
    estado_inmunosupresion = ingresar_logico("Estado de inmunosupresión (s/n): ")

    reporte_atencion = generar_reporte_atencion(
        edad_paciente, enfermedad_cronica, estado_inmunosupresion)

    mostrar_mensaje(reporte_atencion)

def generar_reporte_atencion(
        edad_paciente, enfermedad_cronica, estado_inmunosupresion):
    mensaje = "\nEl paciente es de atención general\n" 

    if edad_paciente < EDAD_RECIEN_NACIDO or \
       edad_paciente > EDAD_ADULTO_MAYOR or \
       enfermedad_cronica or estado_inmunosupresion:
        mensaje = "\nEl paciente es de atención prioritaria\n"
        
    return mensaje

main()