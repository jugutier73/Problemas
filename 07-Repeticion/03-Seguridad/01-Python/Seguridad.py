"""
    Crear un programa para contar cuántas letras mayúsculas, 
    cuántas minúsculas y cuántos dígitos contiene una contraseña

    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Septiembre 2025
    Licencia: GNU GPL v3
"""

from modulo.util import ingresar_texto, mostrar_mensaje

def main():
    contrasenia = ingresar_texto("Ingrese la contraseña a analizar: ") 
    
    cantidad_mayusculas = contar_mayusculas(contrasenia)
    cantidad_minusculas = contar_minusculas(contrasenia)
    cantidad_digitos    = contar_digitos   (contrasenia)
    
    reporte_contrasenia = generar_reporte_contrasenia(contrasenia, 
        cantidad_mayusculas, cantidad_minusculas, cantidad_digitos)

    mostrar_mensaje(reporte_contrasenia)

def contar_mayusculas(texto):
    cantidad_mayusculas = 0

    for caracter in texto:
        if caracter.isupper():
            cantidad_mayusculas += 1

    return cantidad_mayusculas

def contar_minusculas(texto):
    cantidad_minusculas = 0

    for caracter in texto:
        if caracter.islower():
            cantidad_minusculas += 1

    return cantidad_minusculas

def contar_digitos(texto):
    cantidad_digitos = 0

    for caracter in texto:
        if caracter.isdigit():
            cantidad_digitos += 1

    return cantidad_digitos

def generar_reporte_contrasenia(contrasenia, cantidad_mayusculas, cantidad_minusculas, cantidad_digitos):
    return (f"\nEn la constraseña \"{contrasenia}\" hay:\n"
            f"{cantidad_mayusculas} mayúscula(s), "
            f"{cantidad_minusculas} minúscula(s) y "
            f"{cantidad_digitos} dígito(s)\n"
        )

main()