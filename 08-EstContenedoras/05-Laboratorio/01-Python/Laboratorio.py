"""
   Crear un programa para registrar el ingreso y salida de los
   usuarios de un laboratorio para asegurar que solo personas
   autorizadas -estudiantes, docentes o personal técnico- accedan
   al espacio. El programa debe permitir identificar quiénes estaban 
   presentes y la cantidad por cada tipo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, ingresar_opcion, mostrar_mensaje, ingresar_coleccion, convertir_coleccion_cadena, contar_segun_criterio

ANCHO = 20

TOTAL_TIPO_AUTORIZADOS = 3

ESTUDIANTE = 1
DOCENTE = 2
TECNICO = 3

ETIQUETAS = {
    ESTUDIANTE: "Estudiante",
    DOCENTE:    "Docente",
    TECNICO:    "Técnico"
}

def main():
    ingresos = ingresar_coleccion(ingresar_ingreso)
    salidas  = ingresar_coleccion(ingresar_salida)

    personas_laboratorio = obtener_personas_laboratorio(ingresos, salidas)

    cantidad_estudiantes = contar_segun_criterio(personas_laboratorio, 
                                                 tener_tipo, ESTUDIANTE )

    cantidad_docentes    = contar_segun_criterio(personas_laboratorio, 
                                                 tener_tipo, DOCENTE )

    cantidad_tecnico     = contar_segun_criterio(personas_laboratorio, 
                                                 tener_tipo, TECNICO )

    reporte_ingresos = generar_reporte_ingresos(
        ingresos,
        salidas,
        personas_laboratorio, 
        cantidad_estudiantes, 
        cantidad_docentes, 
        cantidad_tecnico
    )

    mostrar_mensaje(reporte_ingresos)

def ingresar_ingreso():
    mostrar_mensaje("\nIngrese los datos del ingreso:\n")
    documento = ingresar_texto   ("\tIngrese el documento   : ")
    nombre = ingresar_texto ("\tIngrese el nombre      : ")
    tipo   = ingresar_opcion (
                        "\tTIPO DE PERSONAS AUTORIZADAS\n" +
                        "\t\t1: Estudiante\n"+
                        "\t\t2: Docente\n"+
                        "\t\t3: Técnico\n"+
                        "\tIngrese tipo de persona: ", 
                        TOTAL_TIPO_AUTORIZADOS
                     )

    return  { "documento": documento, "nombre": nombre, "tipo": tipo }

def ingresar_salida():
    mostrar_mensaje("\nIngrese los datos de la salida:\n")
    documento = ingresar_texto ("\tIngrese el documento   : ")
    return documento

def obtener_personas_laboratorio(ingresos, salidas):
    personas_laboratorio = []

    for ingreso in ingresos:
        if ingreso["documento"] not in salidas:
            personas_laboratorio.append(ingreso.copy()) 

    return personas_laboratorio

def tener_tipo(ingreso, tipo):
    return ingreso["tipo"] == tipo

def generar_reporte_ingresos(
        ingresos, 
        salidas, 
        personas_laboratorio,
        cantidad_estudiantes, 
        cantidad_docentes,
        cantidad_tecnico):
    
    listado_ingresos  = convertir_coleccion_cadena(
      "LISTADO DE INGRESOS", 
      ingresos,  
      convertir_ingresos_cadena)

    listado_salidas  = convertir_coleccion_cadena(
      "LISTADO DE SALIDAS", 
      salidas,  
      convertir_salida_cadena)
    
    listado_personas_laboratorio  = convertir_coleccion_cadena(
      "LISTADO DE PERSONAS EN EL LABORATORIO", 
      personas_laboratorio,  
      convertir_ingresos_cadena)

   
    return (f"{listado_ingresos}\n\n"
            f"{listado_salidas}\n\n"
            f"{listado_personas_laboratorio}\n\n"
            f"CANTIDAD DE PERSONAS AÚN EN EL LABORATORIO\n"
            f"Cantidad de Estudiantes  : {cantidad_estudiantes}\n"
            f"Cantidad de Docentes     : {cantidad_docentes}\n"
            f"Cantidad de Técnicos     : {cantidad_tecnico}\n"
           )

def convertir_ingresos_cadena(ingreso):
    etiqueta = ETIQUETAS[ingreso["tipo"]]
    
    return (f"{ingreso['documento']:{ANCHO}}"
            f"{ingreso['nombre']:{ANCHO}} {etiqueta}\n"
           )

def convertir_salida_cadena(salida):
    return salida + " "

main()




# def obtener_personas_laboratorio(ingresos, salidas):
#     salidas_set = set(salidas)
#     return [ingreso.copy() for ingreso in ingresos 
#             if ingreso["documento"] not in salidas_set]