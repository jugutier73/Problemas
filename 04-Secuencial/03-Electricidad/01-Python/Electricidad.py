"""
    Programa que permita tener un control del consumo de la energía
    eléctrica con relación al consumo del mes anterior, para que 
    así el usuario pueda adoptar hábitos de consumo más conscientes 
    y responsables.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import ingresar_entero, mostrar_mensaje

def main():
    consumo_actual = ingresar_entero("Consumo mes actual   (kilovatios): ")
    consumo_anterior=ingresar_entero("Consumo mes anterior (kilovatios): ")

    relacion_consumo = calcular_relacion_consumo(consumo_actual, consumo_anterior)
    
    reporte_relacion = generar_repore_relacion(consumo_actual, consumo_anterior, relacion_consumo)

    mostrar_mensaje(reporte_relacion)

def calcular_relacion_consumo(consumoActual, consumoAnterior):
    return consumoActual / consumoAnterior * 100.0

def generar_repore_relacion(consumo_actual, consumo_anterior, relacion_consumo):
    return (
        f"\nEl consumo actual de {consumo_actual} kilovatios representa"
        f"\nun {relacion_consumo:.1f}% con relación al consumo del mes"
        f"\nanterior de {consumo_anterior} kilovatios.\n"
    )

main()
