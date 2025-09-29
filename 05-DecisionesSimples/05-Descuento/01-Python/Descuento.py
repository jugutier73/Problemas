"""
    Crear un programa para aplicar un descuento a la tarifa del 
    transporte público según el perfil del usuario, como estudiantes o 
    adultos mayores. Considerando que los domingos y festivos la tarifa 
    es mayor.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_logico

TARIFA_NORMAL = 2900
TARIFA_DOMINGO_FESTIVO = 3000

PORCENTAJE_DESCUENTO_ESTUDIANTE = 10.0
PORCENTAJE_DESCUENTO_ADULTO_MAYOR = 15.0

def main():
    es_domingo_festivo = ingresar_logico("Es domingo o festivo (s/n): ")
    es_estudiante = ingresar_logico("Es estudiante (s/n): ")
    es_adulto_mayor = ingresar_logico("Es adulto mayor (s/n): ")

    tarifa_dia = obtener_tarifa_dia(es_domingo_festivo)

    porcentaje_descuento = obtener_porcentaje_descuento(
        es_estudiante, es_adulto_mayor)

    valor_descuento = calcular_valor_descuento(
        tarifa_dia, porcentaje_descuento)

    valor_tarifa = calcular_valor_tarifa(
        tarifa_dia, valor_descuento)

    recibo_tarifa = generar_recibo_tarifa(
        tarifa_dia, valor_tarifa, porcentaje_descuento, valor_descuento)

    mostrar_mensaje(recibo_tarifa)

def obtener_tarifa_dia(es_domingo_festivo):
    tarifa_dia = TARIFA_NORMAL

    if es_domingo_festivo == True:
        tarifa_dia = TARIFA_DOMINGO_FESTIVO

    return tarifa_dia

def obtener_porcentaje_descuento(es_estudiante, es_adulto_mayor):
    porcentaje_descuento = 0.0

    if es_estudiante == True:
        porcentaje_descuento = PORCENTAJE_DESCUENTO_ESTUDIANTE

    if es_adulto_mayor == True:
        porcentaje_descuento = PORCENTAJE_DESCUENTO_ADULTO_MAYOR

    return porcentaje_descuento

def calcular_valor_descuento(tarifa_dia, porcentaje_descuento):
    return tarifa_dia * porcentaje_descuento / 100.0

def calcular_valor_tarifa(tarifa_dia, valor_descuento):
    return tarifa_dia - valor_descuento

def generar_recibo_tarifa(
        tarifa_dia, valor_tarifa, porcentaje_descuento, valor_descuento):
    mensaje = f"\nLa tarifa es de ${tarifa_dia}\n"

    if valor_descuento > 0:
        mensaje = (
            f"\n{mensaje}"
            f"\nsu tarifa a pagar es de ${valor_tarifa:.0f} por tener"
            f"\nun descuento del {porcentaje_descuento:.1f}% "
            f"equivalente  a ${valor_descuento:.0f}\n"
        )
        
    return mensaje

main()