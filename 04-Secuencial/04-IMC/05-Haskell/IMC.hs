{-
   Programa para calcular el índice de masa corporal (IMC), usando la
   altura y el peso.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarEntero)

main :: IO()
main = do
  peso <- ingresarEntero "Peso  (kg): "
  altura <- ingresarReal "Altura (m): "
  let imc = calcularIMC peso altura
  let informeIMC = generarInformeIMC peso altura imc
  mostrarMensaje informeIMC

ingresarReal :: String -> IO Double
ingresarReal pregunta = do
  mostrarMensaje pregunta
  readLn

calcularIMC :: Int -> Double -> Double
calcularIMC peso altura =
  fromIntegral peso / altura ** 2

generarInformeIMC :: Int -> Double -> Double -> String
generarInformeIMC peso altura imc =
  printf
    ("\nCon su peso de %d kg y su altura de %.1f metros" ++
    "\nsu índice de masa corporal (IMC) es de %.1f")
    peso
    altura
    imc