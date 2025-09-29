"""
   Crear un programa para garantizar que los pacientes sean atendidos
   de manera equitativa y oportuna. El orden de llegada debe
   equilibrarse con la prioridad médica de cada caso

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, ingresar_opcion, mostrar_mensaje, ingresar_coleccion, ordenar_coleccion, convertir_coleccion_cadena

NIVELES = 3

NIVEL_RIESGO = 1
NIVEL_URGENCIA = 2
NIVEL_PRIORITARIO = 3

def main():
    citas = ingresar_coleccion(ingresar_cita)

    citas_por_nivel = ordenar_coleccion(citas, seleccionar_campo_nivel, False)

    cantidad_riesgo      = contar_segun_criterio(citas, tener_nivel_riesgo, NIVEL_RIESGO)
    cantidad_urgencia    = contar_segun_criterio(citas, tener_nivel_riesgo, NIVEL_URGENCIA)
    cantidad_prioritario = contar_segun_criterio(citas, tener_nivel_riesgo, NIVEL_PRIORITARIO)

    reporte_citas = generar_reporte_citas(
        citas_por_nivel, 
        cantidad_riesgo, 
        cantidad_urgencia, 
        cantidad_prioritario)

    mostrar_mensaje(reporte_citas)

def ingresar_cita():
    mostrar_mensaje("\nIngrese los datos de la cita:\n")
    nombre = ingresar_texto  ("\tIngrese el nombre paciente: ")
    nivel  = ingresar_opcion (
                        "\tNIVEL DE URGENCIA\n" +
                        "\t\t1: Riesgo\n"+
                        "\t\t2: Urgencia\n"+
                        "\t\t3: Prioritario\n"+
                        "\tIngrese tipo de persona: ", 
                        NIVELES
                     )

    return  { "nombre": nombre, "nivel": nivel }

def seleccionar_campo_nivel(cita):
    return cita["nivel"]

def contar_segun_criterio(coleccion, aplicar_criterio, valor_criterio):
    cantidad = 0

    for elemento in coleccion:
        if aplicar_criterio(elemento, valor_criterio):
            cantidad += 1
    
    return cantidad

def tener_nivel_riesgo(cita, nivel_riesgo):
    return cita["nivel"] == nivel_riesgo

def generar_reporte_citas(citas_por_nivel, 
                          cantidad_riesgo, 
                          cantidad_urgencia, 
                          cantidad_prioritario):

    listado_por_nivel  = convertir_coleccion_cadena(
      "LISTADO POR NIVEL DE URGENCIA", 
      citas_por_nivel,  
      convertir_cita_cadena)
   
    return (f"{listado_por_nivel}\n\n"
            f"Nivel 1 (Riesgo)     : {cantidad_riesgo}\n"
            f"Nivel 2 (Urdencia)   : {cantidad_urgencia}\n"
            f"Nivel 3 (Prioritario): {cantidad_prioritario}\n"
           )

def convertir_cita_cadena(cita):
    return f"{cita['nivel']} : {cita['nombre']}\n"

main()