"""
   Crear un programa para determinar si hay problemas con el uso
   de los espacios en un frase que el usuario ingrese

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, mostrar_mensaje

def main():
    frase =  ingresar_texto("Ingrese una frase a analizar: ")
    
    frase_corregida = corregir_uso_espacios(frase)
    
    reporte = generar_reporte_alfabetizacion(frase, frase_corregida)

    mostrar_mensaje(reporte)

def corregir_uso_espacios(frase):
    inicio_frase = False
    espacio_pendiente = False 
    frase_corregida = ""

    for simbolo in frase:
        if simbolo == " ":
            if inicio_frase:
                espacio_pendiente = True
        else:
            if espacio_pendiente:
                frase_corregida += " "
                espacio_pendiente = False
            frase_corregida += simbolo            
            inicio_frase = True
            
    return frase_corregida

def generar_reporte_alfabetizacion(frase, frase_corregida):
    reporte_alfabetizacion =  f"\nLa frase \"{frase}\" hace un uso "

    if frase == frase_corregida:
        reporte_alfabetizacion += f"CORRECTO de espacios."
    else:
        reporte_alfabetizacion += (f"INCORRECTO de espacios,\n" 
                                   f"lo correcto es \"{frase_corregida}\"")
                                   
    return reporte_alfabetizacion + "\n"

main()