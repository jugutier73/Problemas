"""
    Crear un programa para evaluar el nivel de estrés de los estudiantes

    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_opcion

OPCIONES = 5

NIVEL_BAJO = 19
NIVEL_MODERADO = 25

def main():
    respuesta01 = ingresar_respuesta(
        "¿con qué frecuencia te has sentido afectado por algo que ocurrió inesperadamente?")
    respuesta02 = ingresar_respuesta(
        "¿con qué frecuencia te has sentido incapaz de controlar las  cosas importantes en tu vida?")
    respuesta03 = ingresar_respuesta(
        "¿con qué frecuencia te has sentido nervioso o estresado?")
    respuesta04 = ingresar_respuesta_invertida(
        "¿con qué frecuencia has manejado con éxito los pequeños problemas irritantes de la vida?")
    respuesta05 = ingresar_respuesta_invertida(
        "¿con qué frecuencia has sentido que has afrontado efectivamente los cambios importantes que han estado ocurriendo en tu vida?")
    respuesta06 = ingresar_respuesta_invertida(
        "¿con qué frecuencia has estado seguro sobre tu capacidad para manejar tus problemas personales?")
    respuesta07 = ingresar_respuesta_invertida(
        "¿con qué frecuencia has sentido que las cosas van bien?")
    respuesta08 = ingresar_respuesta(
        "¿con qué frecuencia has sentido que no podías afrontar todas las cosas que tenías que hacer?")
    respuesta09 = ingresar_respuesta_invertida(
        "¿con qué frecuencia has podido controlar las dificultades de tu vida?")
    respuesta10 = ingresar_respuesta_invertida(
        "¿con qué frecuencia has sentido que tenías todo bajo control?")
    respuesta11 = ingresar_respuesta(
        "¿con qué  frecuencia has estado enfadado porque las cosas que te han ocurrido estaban fuera de tu control?")
    respuesta12 = ingresar_respuesta(
        "¿con qué frecuencia has pensado sobre las cosas que te faltan por hacer?")
    respuesta13 = ingresar_respuesta_invertida(
        "¿con qué frecuencia has podido controlar la forma de pasar el tiempo?")
    respuesta14 = ingresar_respuesta(
        "¿con qué frecuencia has sentido que las dificultades se acumulan tanto que no puedes superarlas?")

    puntaje_total = calcular_puntaje_total(respuesta01, respuesta02, 
        respuesta03, respuesta04, respuesta05, respuesta06, 
        respuesta07, respuesta08, respuesta09, respuesta10,
        respuesta11, respuesta12, respuesta13, respuesta14)
    
    nivel_estres = obtener_nivel_estres(puntaje_total)

    reporte_estres = generar_reporte_estres(nivel_estres, puntaje_total)

    mostrar_mensaje(reporte_estres)

def ingresar_respuesta(pregunta):
    # Se restar uno para obtener una escala de 0 a 4
    return ingresar_opcion(f"\nEn el último mes, {pregunta}\n\n" +
                           "\t(1) Nunca,\n" +
                           "\t(2) Casi nunca,\n" +
                           "\t(3) De vez en cuando\n" +
                           "\t(4) A menudo\n" +
                           "\t(5) Muy a menudo\n\n" +
                           "\tCuál es su frecuencia: ", OPCIONES) - 1

def ingresar_respuesta_invertida(pregunta):
    # Se usa valor absoluto(abs) al restar 4 para tener la escala 4 a 0
    return abs(ingresar_respuesta(pregunta) - OPCIONES + 1)

def calcular_puntaje_total(respuesta01, respuesta02, respuesta03, 
        respuesta04, respuesta05, respuesta06, respuesta07, 
        respuesta08, respuesta09, respuesta10, respuesta11, 
        respuesta12, respuesta13, respuesta14):
    return respuesta01 + respuesta02 + respuesta03 + respuesta04 + \
           respuesta05 + respuesta06 + respuesta07 + respuesta08 + \
           respuesta09 + respuesta10 + respuesta11 + respuesta12 + \
           respuesta13 + respuesta14

def obtener_nivel_estres(puntaje_total):

    if puntaje_total < NIVEL_BAJO:
        nivel_estres = "BAJO"
    else:
        if puntaje_total < NIVEL_MODERADO:
            nivel_estres = "MODERADO"
        else:
            nivel_estres = "ALTO"

    return nivel_estres

def generar_reporte_estres(nivel_estres, puntaje_total):
    return (f"\nSu nivel de estrés es {nivel_estres}, "
            f"con un puntaje de {puntaje_total}.\n"
    )

main()