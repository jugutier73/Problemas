"""
    Crear un programa para calcular, de forma clara y transparente, 
    el valor de un trayecto en taxi considerando variables como 
    el horario diurno/nocturno, los recargos de domingos y festivos, 
    y el destino (ya sea dentro del perímetro urbano o fuera de él)
    para determinar si la tarifa cobrada es justa o no.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Julio 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_entero, \
    ingresar_opcion, ingresar_logico

URBANO = 1
PERIFERIA = 2
EXTRA_URBANO = 3
VIA_AEROPUERTO = 4

SIN_COSTO = 0
VALOR_MINIMO = 6000
VALOR_ESPECIAL = 1300

VALOR_PERIFERIA = 3400
VALOR_EXTRA_URBANO = 4900
VALOR_VIA_AEROPUERTO = 1600

def main():
    valor_cobrado = ingresar_entero("Valor cobrado: ")
    valor_taximetro = ingresar_entero("Valor del taxímetro: ")
    es_especial = ingresar_logico("Es domingo, festivo o nocturno (s/n): ")
    destino_trayecto = ingresar_opcion("Destino \n" +
                                       "\t(1) urbano,\n" +
                                       "\t(2) lugar periférico,\n" +
                                       "\t(3) extra urbano\n" +
                                       "\t(4) vía al aeropuerto\n\n" +
                                       "Cuál es su destino: ", 4)

    tarifa_minima  = determinar_tarifa_minima(valor_taximetro)
    tarifa_especial= determinar_tarifa_especial(es_especial)
    tarifa_destino = determinar_tarifa_destino(destino_trayecto) 

    tarifa_real = calcular_tarifa(
        tarifa_minima, tarifa_especial, tarifa_destino)
    
    mensaje_cobro = determinar_mensaje_cobro(valor_cobrado, tarifa_real)

    informe_cobro = generar_informe_cobro(
        valor_cobrado, tarifa_real, mensaje_cobro)

    mostrar_mensaje(informe_cobro)

def determinar_tarifa_minima(valor_taximetro):
    return max(valor_taximetro, VALOR_MINIMO)

def determinar_tarifa_especial(es_especial):
    tarifa_especial = SIN_COSTO

    if es_especial:
        tarifa_especial = VALOR_ESPECIAL
        
    return tarifa_especial

def determinar_tarifa_destino(destino_trayecto):
    tarifa_destino = SIN_COSTO

    if destino_trayecto == PERIFERIA:
        tarifa_destino = VALOR_PERIFERIA
    elif destino_trayecto == EXTRA_URBANO:
        tarifa_destino = VALOR_EXTRA_URBANO
    elif destino_trayecto == VIA_AEROPUERTO:
        tarifa_destino = VALOR_VIA_AEROPUERTO

    return tarifa_destino

def calcular_tarifa(tarifa_minima, tarifa_especial, tarifa_destino):
    return tarifa_minima + tarifa_especial + tarifa_destino

def determinar_mensaje_cobro(valor_cobrado, tarifa_real):
    mensaje_cobro = "JUSTO"

    if valor_cobrado != tarifa_real:
        mensaje_cobro = "INJUSTO"

    return mensaje_cobro

def generar_informe_cobro(valor_cobrado, tarifa_real, mensaje_cobro):
    return (
        f"\nValor cobrado \t${valor_cobrado:7d}"
        f"\nValor real \t${tarifa_real:7d}\n"
        f"\nPor lo anterior el cobro es {mensaje_cobro}\n"
    )

main()