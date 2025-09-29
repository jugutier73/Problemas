"""
   Programa que registra los residuos recolectados durante una
   jornada, especificando su código, tipo y condición de
   reparabilidad. El sistema genera un informe con el
   listado de los elementos recolectados, ordenados por tipo y
   reparabilidad, además de presentar un conteo total por cada
   tipo de residuo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto,  ingresar_opcion, \
                        ingresar_logico, mostrar_mensaje, \
                        ingresar_coleccion, ordenar_coleccion,\
                        convertir_coleccion_cadena, \
                        contar_segun_criterio

TOTAL_TIPOS = 5

BATERIA = 1
TELEFONO = 2
COMPUTADOR = 3
CARGADOR = 4
OTRO_DISPOSITIVO = 5

ANCHO = 15

ETIQUETAS = {
    BATERIA:         "Batería",
    TELEFONO:        "Teléfono",
    COMPUTADOR:      "Computador",
    CARGADOR:        "Cargador",
    OTRO_DISPOSITIVO:"Otro",
}

def main():
    residuos = ingresar_coleccion(ingresar_residuo)

    residuos_clasificados = ordenar_coleccion(residuos,
                                              seleccionar_campo_tipo,False)

    cantidad_baterias     = contar_segun_criterio(residuos, 
                                                  tener_tipo, BATERIA)
    
    cantidad_telefonos    = contar_segun_criterio(residuos, 
                                                  tener_tipo, TELEFONO)
    
    cantidad_computadores = contar_segun_criterio(residuos, 
                                                  tener_tipo, COMPUTADOR)
    
    cantidad_cargadores   = contar_segun_criterio(residuos, 
                                                  tener_tipo, CARGADOR)
    
    cantidad_otros        = contar_segun_criterio(residuos, 
                                                  tener_tipo,
                                                  OTRO_DISPOSITIVO)

    reporte_recoleccion = generar_reporte_recoleccion(
        residuos_clasificados, 
        cantidad_baterias, 
        cantidad_telefonos,
        cantidad_computadores,
        cantidad_cargadores,
        cantidad_otros
    )

    mostrar_mensaje(reporte_recoleccion)

def ingresar_residuo():
    mostrar_mensaje("\nIngrese los datos del residuo electrónico:\n")
    codigo = ingresar_texto   ("\tIngrese el código del residuo: ")
    tipo   = ingresar_opcion  (
                        "\tTIPOS DE RESIDUOS\n" +
                        "\t\t1: Batería\n"+
                        "\t\t2: Teléfono\n"+
                        "\t\t3: Computador\n"+
                        "\t\t4: Cargador\n"+
                        "\t\t5: Otro dispositivo)\n" +
                        "\tIngrese tipo de residuo      : ", 
                        TOTAL_TIPOS
                     )
    reparable = ingresar_logico ("\tPuede ser reparado (s/n): ")

    return  { "codigo": codigo, "tipo": tipo, "reparable": reparable }

def seleccionar_campo_tipo(residuo):
    return (residuo["tipo"], -residuo["reparable"])

def tener_tipo(residuo, tipo_interes):
    return residuo["tipo"] == tipo_interes

def generar_reporte_recoleccion(
        residuos_clasificados, 
        cantidad_baterias, 
        cantidad_telefonos,
        cantidad_computadores,
        cantidad_cargadores,
        cantidad_otros):

    listado_residuos  = convertir_coleccion_cadena(
      "LISTADO DE RESIDUOS CLASIFICADOS", 
      residuos_clasificados,  
      convertir_residuo_cadena)
   
    return (f"{listado_residuos}\n\n"
            f"Cantidad de Baterías     : {cantidad_baterias}\n"
            f"Cantidad de Teléfonos    : {cantidad_telefonos}\n"
            f"Cantidad de Computadores : {cantidad_computadores}\n"
            f"Cantidad de Cargadores   : {cantidad_cargadores}\n"
            f"Cantidad de Otros        : {cantidad_otros}\n"
           )

def convertir_residuo_cadena(residuo):
    etiqueta = ETIQUETAS[residuo["tipo"]]
    reparable = "(REPARABLE)" if residuo["reparable"] else ""
    
    return f"{residuo['codigo']:{ANCHO}} {etiqueta:{ANCHO}} {reparable}\n"

main()
