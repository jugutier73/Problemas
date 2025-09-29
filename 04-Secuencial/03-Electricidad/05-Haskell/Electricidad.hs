{-
    Programa que permita tener un control del consumo de la energía
    eléctrica con relación al consumo del mes anterior, para que 
    así el usuario pueda adoptar hábitos de consumo más conscientes 
    y responsables.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarEntero)

main :: IO()
main = do
  consumoActual <- ingresarEntero "Consumo mes actual   (kilovatios): "
  consumoAnterior <- ingresarEntero "Consumo mes anterior (kilovatios): "

  let relacionConsumo = calcularRelacionConsumo consumoActual consumoAnterior
  let reporteRelacion = generarReporeRelacion consumoActual consumoAnterior relacionConsumo
  
  mostrarMensaje reporteRelacion


calcularRelacionConsumo :: Int -> Int -> Double
calcularRelacionConsumo consumoActual consumoAnterior =
  fromIntegral consumoActual / fromIntegral consumoAnterior * 100.0


generarReporeRelacion :: Int -> Int -> Double -> String
generarReporeRelacion consumoActual consumoAnterior relacionConsumo =
  printf
    ("\nEl consumo actual de %d kilovatios representa" ++
    "\nun %.1f%% con relación al consumo del mes" ++
    "\nanterior de %d kilovatios.\n")
    consumoActual
    relacionConsumo
    consumoAnterior