"""
    Programa para calcular la cantidad de CO2 emitido por el uso de
    transporte particular (carro y moto), el empleo del transporte 
    público (buses)
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_real

HUELLA_CARRO = 121.0
HUELLA_MOTO  = 53.0
HUELLA_BUS   = 49.0

def main():
    km_carro = ingresar_real("Total de kilómetros recorridos en carro: ")
    km_moto  = ingresar_real("Total de kilómetros recorridos en moto : ")
    km_buses = ingresar_real("Total de kilómetros recorridos en bus  : ")

    huella = calcular_huella_carbono(km_carro, km_moto, km_buses)
    informe_huella = generar_huella(km_carro, km_moto, km_buses, huella)
    
    mostrar_mensaje(informe_huella)

def calcular_huella_carbono(km_carro, km_moto, km_buses):
    return km_carro * HUELLA_CARRO + \
           km_moto  * HUELLA_MOTO + \
           km_buses * HUELLA_BUS

def generar_huella(km_carro, km_moto, km_buses, huella):
    return (
        f"\nCon {km_carro}, {km_moto}, {km_buses} km de recorrido"
        f"\nen carro, moto y bus respectivamente,"
        f"\nsu huella de carbono por el uso de"
        f"\ntransporte es de {huella:.1f} kg de CO2.\n"
    )

main()