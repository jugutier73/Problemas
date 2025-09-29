{-
   Crear un programa para contar cuántas letras mayúsculas,
   cuántas minúsculas y cuántos dígitos contiene una contraseña

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarTexto)
import Data.Char (isUpper, isLower, isDigit)

main :: IO()
main = do
    contrasenia <- ingresarTexto "Ingrese la contraseña a analizar: "
    
    let cantidadMayusculas = contarMayusculas contrasenia
    let cantidadMinusculas = contarMinusculas contrasenia
    let cantidadDigitos = contarDigitos contrasenia

    let reporteEstres = generarReporteContrasenia contrasenia cantidadMayusculas cantidadMinusculas cantidadDigitos

    mostrarMensaje reporteEstres

contarMayusculas :: String -> Int
contarMayusculas texto = contarCaracteres texto esMayuscula

contarMinusculas :: String -> Int
contarMinusculas texto = contarCaracteres texto esMinuscula

contarDigitos :: String -> Int
contarDigitos texto = contarCaracteres texto esDigito

contarCaracteres :: String -> (Char -> Bool) -> Int
contarCaracteres texto funcion = length (filter funcion texto)

esMayuscula :: Char -> Bool
esMayuscula caracter =  isUpper caracter

esMinuscula :: Char -> Bool
esMinuscula caracter =  isLower caracter

esDigito :: Char -> Bool
esDigito caracter =  isDigit caracter

generarReporteContrasenia :: String -> Int -> Int -> Int -> String
generarReporteContrasenia contrasenia cantidadMayusculas cantidadMinusculas cantidadDigitos =
  printf ("\nEn la constraseña \"%s\" hay:\n" ++
          "%d mayúscula(s), "++
          "%d minúscula(s) y "++
          "%d dígito(s)\n")
          contrasenia cantidadMayusculas cantidadMinusculas cantidadDigitos
