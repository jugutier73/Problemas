"""
   Crear un programa para organizar la entrada de los beneficiarios
   y garantice la atención en orden de llegada. Se debe registrar
   datos clave (nombre, edad y presencia de necesidades especiales).
   Además de informar la cantidad de personas con necesidades y el
   promedio de edades.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, ingresar_entero, ingresar_logico, mostrar_mensaje, ingresar_coleccion, convertir_coleccion_cadena, contar_segun_criterio

def main():
    reservas_comedor = ingresar_coleccion(ingresar_reverva)

    cantidad_con_necesidades = contar_segun_criterio(reservas_comedor, tener_necesidad_especial, True)

    promedio_edades = calcular_promedio_edades(reservas_comedor)

    reporte_reservas_comedor = generar_reporte_reservas_comedor(
        reservas_comedor, 
        promedio_edades, 
        cantidad_con_necesidades
    )

    mostrar_mensaje(reporte_reservas_comedor)

def ingresar_reverva():
    mostrar_mensaje("\nIngrese los datos de la reserva:\n")
    nombre = ingresar_texto ("\tIngrese el nombre de la persona  : ")
    edad = ingresar_entero ("\tIngrese la edad de la persona    : ")
    necesidad_especial = ingresar_logico ("\tTiene necesidades especiales (s/n): ")

    return  { "nombre": nombre, 
              "edad": edad, 
              "necesidad_especial": necesidad_especial }

def tener_necesidad_especial(reservas_comedor, necesidad_especial_interes):
    return reservas_comedor["necesidad_especial"] == necesidad_especial_interes

def calcular_promedio_edades(reservas_comedor):
    total_reservas = len(reservas_comedor)
    promedio_edades = 0.0
    suma_edades = 0.0
    
    for reserva in reservas_comedor:
        suma_edades += reserva["edad"]

    if total_reservas > 0:
        promedio_edades = suma_edades / total_reservas
    
    return promedio_edades

def generar_reporte_reservas_comedor(reservas_comedor, 
                                     promedio_edad,
                                     cantidad_necesidades_especiales):

    listado_reservas  = convertir_coleccion_cadena(
      "LISTADO DE RESERVAS", reservas_comedor,  convertir_reserva_cadena)
   
    return (f"{listado_reservas}\n\n"
            f"Cantidad con necesidades : {cantidad_necesidades_especiales}\n"
            f"Promedio de edades       : {promedio_edad:.1f}\n"
           )

def convertir_reserva_cadena(reserva):
    mensaje =  f"{reserva['nombre']}, {reserva['edad']} años"

    if reserva['necesidad_especial']:
        mensaje += ", con necesidad especial"
  
    return mensaje + "\n"

main()


# def contar_necesidades_especiales(reservas_comedor):
#     return sum(reserva["necesidad_especial"] for reserva in reservas_comedor)


# def calcular_promedio_edades(reservas_comedor):
#     total_reservas = len(reservas_comedor)

#     promedio_edades = (sum(reserva["edad"] for reserva in reservas_comedor) / total_reservas) \
#                         if total_reservas > 0 else 0.0

#     return promedio_edades