"""
    Programa para verificar si una etiqueta de reciclaje es válida 
    y reportar los campos con error.

    Formato: T-aaaa-mm
    - T: tipo de residuo 
      (P: plástico, V: vidrio, M: metal, C: cartón/papel, O: orgánico)
    - aaaa: año (> 2025)
    - mm: mes (01-12)

    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Septiembre 2025
    Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, mostrar_mensaje

LONGITUD_ETIQUETA = 9  # T-aaaa-mm

SEPARADOR           = "-"
POSICION_SEPARADOR1 = 1
POSICION_SEPARADOR2 = 6

POSICION_TIPO       = 0

POSICION_ANIO       = 2
LONGITUD_ANIO       = 4

POSICION_MES        = 7
LONGITUD_MES        = 2

MINIMO_ANIO         = 2026
MINIMO_MES          = 1
MAXIMO_MES          = 12

TIPOS_VALIDOS       = "PVMCO"

ETIQUETA_OK         = 0   # 0

ERROR_LONGITUD      = 1   # 2^0
ERROR_FORMATO       = 2   # 2^1
ERROR_TIPO          = 4   # 2^2
ERROR_ANIO          = 8   # 2^3
ERROR_MES           = 16  # 2^4

def main():
    etiqueta =  ingresar_texto("Ingrese la etiqueta a analizar: ")
    
    codigo = verificar_etiqueta_reciclaje(etiqueta)
    
    reporte_etiqueta = generar_reporte_etiqueta(etiqueta, codigo)

    mostrar_mensaje(reporte_etiqueta)

def verificar_etiqueta_reciclaje(etiqueta):
    codigo  = verificar_logitud(etiqueta)
 
    if codigo != ERROR_LONGITUD:
        codigo = verificar_formato(etiqueta) | \
                 verificar_tipo(etiqueta)    | \
                 verificar_anio(etiqueta)    | \
                 verificar_mes(etiqueta)
    
    return codigo

def verificar_logitud(etiqueta):
    codigo_error = ETIQUETA_OK

    if len(etiqueta) != LONGITUD_ETIQUETA:
        codigo_error = ERROR_LONGITUD

    return codigo_error

def verificar_formato(etiqueta):
    codigo_error = ETIQUETA_OK

    if etiqueta[POSICION_SEPARADOR1] != SEPARADOR or \
       etiqueta[POSICION_SEPARADOR2] != SEPARADOR:
        codigo_error = ERROR_FORMATO

    return codigo_error

def verificar_tipo(etiqueta):
    codigo_error = ETIQUETA_OK

    tipo = etiqueta[POSICION_TIPO]

    if tipo not in TIPOS_VALIDOS:
        codigo_error = ERROR_TIPO

    return codigo_error

def verificar_anio(etiqueta):
    codigo_error = ETIQUETA_OK

    anio_etiqueta = etiqueta[POSICION_ANIO:POSICION_ANIO + LONGITUD_ANIO]

    if anio_etiqueta.isdigit():
        anio = int(anio_etiqueta)
        if anio < MINIMO_ANIO:
            codigo_error = ERROR_ANIO
    else:
        codigo_error = ERROR_ANIO

    return codigo_error

def verificar_mes(etiqueta):
    codigo_error = ETIQUETA_OK

    mes_etiqueda = etiqueta[POSICION_MES:POSICION_MES + LONGITUD_MES] 

    if mes_etiqueda.isdigit():
        mes = int(mes_etiqueda)
        if  mes < MINIMO_MES or mes > MAXIMO_MES:
            codigo_error = ERROR_MES
    else:
        codigo_error = ERROR_MES

    return codigo_error

def generar_reporte_etiqueta(etiqueta, codigo):
    reporte =  f"\nLa  etiqueta \"{etiqueta}\" "

    if codigo == ETIQUETA_OK: 
        reporte += "es válida."
    else:
        reporte += "es inválida por tener:"

        if codigo & ERROR_LONGITUD != 0:  
            reporte += "\n  - Error en la longitud"
            
        if codigo & ERROR_FORMATO  != 0:  
            reporte += "\n  - Error en el formato"

        if codigo & ERROR_TIPO     != 0:  
            reporte += "\n  - Error en el tipo"

        if codigo & ERROR_ANIO     != 0:  
            reporte += "\n  - Error en el año"

        if codigo & ERROR_MES      != 0:  
            reporte += "\n  - Error en el mes"

    return reporte +"\n"

main()