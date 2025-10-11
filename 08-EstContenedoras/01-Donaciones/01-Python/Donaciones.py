"""
    Crear un programa para registrar y consultar de manera ágil
    todas las donaciones recibidas. Esta debería ofrecer el total 
    recaudado, listados alfabéticos y descendentes por valor, 
    identificar a los aportantes que superen un umbral específico y 
    señalar al mayor donante.

    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Diciembre 2025
    Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, \
                        ingresar_entero, \
                        ingresar_logico, \
                        mostrar_mensaje

UMBRAL = 10000

def main():
    donantes = ingresar_coleccion(ingresar_donante)

    donantes_por_nombre   = ordenar_coleccion(donantes, 
                                              comparar_nombre,   
                                              False)
    
    donantes_por_donacion = ordenar_coleccion(donantes, 
                                              comparar_donacion, 
                                              True)

    mayores_donantes   = obtener_mayores_donantes(donantes, UMBRAL)
    mayor_donante      = obtener_mayor_donante   (donantes)
    suma_donantes      = obtener_suma_donantes   (donantes)

    reporte_donantes = generar_reporte_donantes(donantes_por_nombre,
                                                donantes_por_donacion, 
                                                mayores_donantes, 
                                                mayor_donante, 
                                                suma_donantes)

    mostrar_mensaje(reporte_donantes)

def ingresar_coleccion(ingresar_elemento):
    coleccion = []
    hay_mas = True

    while hay_mas:
        coleccion.append(ingresar_elemento())

        hay_mas = ingresar_logico("Hay más datos (s/n): ")
        
    return coleccion

def ingresar_donante():
    mostrar_mensaje("\nIngrese los datos de un donante:\n")
    nombre   = ingresar_texto  ("\tIngrese el nombre: ")
    donacion = ingresar_entero ("\tIngrese donación : ")

    return  { "nombre": nombre, "donacion": donacion }

def ordenar_coleccion(coleccion, comparador, descendente):
    return sorted(coleccion, key=comparador, reverse=descendente)

def comparar_nombre(donacion):
    return donacion["nombre"]

def comparar_donacion(donacion):
    return donacion["donacion"]

def obtener_mayores_donantes(donantes, valor_limite):
    mayores_donantes = []

    for elemento in donantes:
        if elemento["donacion"] > valor_limite:
            mayores_donantes.append(elemento)
    
    return mayores_donantes

def obtener_mayor_donante(donantes):
    mayor_donante = None

    for elemento in donantes:
        if mayor_donante is None or \
           mayor_donante["donacion"] < elemento["donacion"]:
            mayor_donante = elemento
    
    return mayor_donante

def obtener_suma_donantes(donantes):
    suma_donantes = 0

    for elemento in donantes:
        suma_donantes += elemento["donacion"]
    
    return suma_donantes

def generar_reporte_donantes(donantes_por_nombre,       
                               donantes_por_donacion, 
                               donantes_mayores, 
                               mayor_donante,  
                               suma_donantes):

    listado_por_nombre  = convertir_coleccion_cadena(
      "LISTADO EN ORDEN ALFABÉTICO", 
      donantes_por_nombre, 
      convertir_donante_cadena)
    
    listado_por_donacion  = convertir_coleccion_cadena(
      "LISTADO ORDENADO POR DONACIÓN", 
      donantes_por_donacion, 
      convertir_donante_cadena)
    
    listado_mayores  = convertir_coleccion_cadena(
      f"LISTADO DONANTES MAYORES A ${UMBRAL}", donantes_mayores, convertir_donante_cadena)
    
    return (f"{listado_por_nombre}\n"
            f"{listado_por_donacion}\n"
            f"{listado_mayores}\n"
            f"El mayor donante: {mayor_donante["nombre"]}\n"
            f"Total de donaciones ${suma_donantes}\n"
           )

def convertir_coleccion_cadena(titulo, coleccion, convertir_elemento_cadena):
    mensaje = f"\n{titulo}\n"

    for elemento in coleccion:
        mensaje += "\t" + convertir_elemento_cadena(elemento)

    return mensaje

def convertir_donante_cadena(donante):
    return f"${donante['donacion']:10} \t {donante['nombre']}\n"

main()