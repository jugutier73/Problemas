"""
   Crear un programa para registrar síntomas y recibir recomendaciones 
   simples y seguras como reposar, hidratarse o consultar a un profesional 
   de la salud, según la gravedad o combinación de síntomas reportados.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
"""

from modulo.util import ingresar_logico, ingresar_real, mostrar_mensaje

FIEBRE_ALTA = 39.5
FIEBRE = 37.5

MENSAJE_ESPECIALISTA = """
 - Consultar un especialista.
 - Anotar los síntomas y cuándo comenzaron.
 - Evitar esfuerzos físicos y actividades intensas."""

MENSAJE_FIEBRE_ALTA = """
 - Solicitar una cita medica con urgencia.
 - Usar paños húmedos y fríos en la frente.
 - Permanecer en un lugar fresco y ventilado."""

MENSAJE_FIEBRE = """
 - Descansar lo suficiente.
 - Hidratarse bebiendo agua u otros líquidos."""

MENSAJE_DOLOR_CABEZA = """
 - Realizar ejercicio cervical isométrico.
 - Descansar en un lugar oscuro y silencioso.
 - Beber agua, ya que la deshidratación puede empeorar el dolor."""

MENSAJE_CONGESTION = """
 - Realizar lavados nasales con solución salina.
 - Usar almohadas extras para dormir con la cabeza elevada."""

def main():
    temperatura = ingresar_real("Temperatura corporal: ")
    sintomas_varios_dias = ingresar_logico("Síntomas por 2 o más días (s/n): ")
    malestar_intenso = ingresar_logico("Tiene malestar intenso (s/n): ")
    dolor_cabeza = ingresar_logico("Tiene dolor de cabeza (s/n): ")
    congestion_nasal = ingresar_logico("Tiene congestion nasal (s/n): ")

    recomendaciones_especialista = recibir_recomendaciones(
        sintomas_varios_dias or malestar_intenso, MENSAJE_ESPECIALISTA)
    
    recomendaciones_fiebre_alta = recibir_recomendaciones(
        temperatura >= FIEBRE_ALTA, MENSAJE_FIEBRE_ALTA)
    
    recomendaciones_fiebre = recibir_recomendaciones(
        temperatura >= FIEBRE, MENSAJE_FIEBRE)
    
    recomendaciones_dolor_cabeza = recibir_recomendaciones(
        dolor_cabeza, MENSAJE_DOLOR_CABEZA)
    
    recomendaciones_congestion = recibir_recomendaciones(
        congestion_nasal, MENSAJE_CONGESTION)

    recomendaciones = generar_reporte_recomendaciones(
        recomendaciones_especialista, recomendaciones_fiebre_alta,
        recomendaciones_fiebre,       recomendaciones_dolor_cabeza,
        recomendaciones_congestion)

    mostrar_mensaje(recomendaciones)

def recibir_recomendaciones(condicion, recomendaciones):
    return recomendaciones if condicion else ""

def generar_reporte_recomendaciones(
        recomendaciones_especialista, recomendaciones_FIEBRE_ALTA,
        recomendaciones_fiebre,       recomendaciones_dolor_cabeza,
        recomendaciones_congestion):
    reporte_recomendaciones = "\nSe recomienda:"

    if recomendaciones_especialista != "":
        reporte_recomendaciones += recomendaciones_especialista
    else:
        reporte_recomendaciones += recomendaciones_FIEBRE_ALTA + \
            recomendaciones_fiebre + recomendaciones_dolor_cabeza + \
            recomendaciones_congestion

    return reporte_recomendaciones + "\n"

main()
