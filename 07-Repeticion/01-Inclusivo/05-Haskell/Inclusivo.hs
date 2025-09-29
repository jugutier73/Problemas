{-
   Crear un programa para contar cuantas veces se emplean símbolos 
   "x" o la "@" en un mensaje.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarTexto)

simboloInclusivo1, simboloInclusivo2 :: Char
simboloInclusivo1 = 'x'
simboloInclusivo2 = '@'

main :: IO()
main = do
    texto <- ingresarTexto "Ingrese un texto con los mensajes: "
    
    let cantidadSimbolos = contarEmpleoSimbolos texto

    let reporteSimbolos = generarReporteInclusivo cantidadSimbolos

    mostrarMensaje reporteSimbolos

contarEmpleoSimbolos :: String -> Int

contarEmpleoSimbolos [] = 0

contarEmpleoSimbolos (c:cs)
  | c == simboloInclusivo1 ||  c == simboloInclusivo2 = 
    1 + contarEmpleoSimbolos cs
  | otherwise = contarEmpleoSimbolos cs

generarReporteInclusivo :: Int -> String
generarReporteInclusivo cantidadSimbolosInclusivos =
  printf ("\nSe emplearon %d veces " ++
          "los símbolos inclusivos \"%c\" y \"%c\".\n") 
          cantidadSimbolosInclusivos 
          simboloInclusivo1 simboloInclusivo2