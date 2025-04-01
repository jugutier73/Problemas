"""
    Programa para calcular el índice de masa corporal (IMC), usando la 
    altura y el peso.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
"""

from modulo.util import mostrar_mensaje, ingresar_entero, ingresar_texto

def main():
    peso = ingresar_entero("Peso  (kg): ")
    altura = ingresar_real("Altura (m): ")
    imc = calcular_IMC(peso, altura);            
    informeIMC = generar_informe_IMC(peso, altura, imc);
    mostrar_mensaje(informeIMC)

def ingresar_real(pregunta):
    return float(ingresar_texto(pregunta))

def calcular_IMC(peso, altura):
    return peso / (altura**2.0);

def generar_informe_IMC(peso, altura, imc):
    return (
        f"\nCon su peso de {peso} kg y su altura de {altura:.1f} metros"
        f"\nsu índice de masa corporal (IMC) es de {imc:.1f}"
    )

main()
