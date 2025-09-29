{-
   Crear un programa para calcular, de forma clara y transparente,
   el valor de un trayecto en taxi considerando variables como
   el horario diurno/nocturno, los recargos de domingos y festivos,
   y el destino (ya sea dentro del perímetro urbano o fuera de él)
   para determinar si la tarifa cobrada es justa o no.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarEntero, ingresarOpcion, ingresarLogico)

urbano, periferia, extraUrbano, viaAeropuerto :: Int
urbano        = 1
periferia     = 2
extraUrbano   = 3
viaAeropuerto = 4

sinCosto, valorMinimo, valorEspecial :: Int
sinCosto      = 0
valorMinimo   = 6000
valorEspecial = 1300

valorPeriferia, valorExtraUrbano, valorViaAeropuerto :: Int
valorPeriferia     = 3400
valorExtraUrbano   = 4900
valorViaAeropuerto = 1600

main :: IO()
main = do
  valorCobrado <- ingresarEntero "Valor cobrado: "
  valorTaximetro <- ingresarEntero "Valor del taxímetro: "
  esEspecial <- ingresarLogico "Es domingo, festivo o nocturno (s/n): "
  destinoTrayecto <- ingresarOpcion ( "Destino \n" ++
    "\t(1) urbano,\n" ++
    "\t(2) lugar periférico,\n" ++
    "\t(3) extra urbano\n" ++
    "\t(4) vía al aeropuerto\n\n" ++
    "Cuál es su destino: " ) 4
  
  let tarifaMinima = determinarTarifaMinima valorTaximetro
  let tarifaEspecial = determinarTarifaEspecial esEspecial
  let tarifaDestino = determinarTarifaDestino destinoTrayecto

  let tarifaReal = calcularTarifa tarifaMinima tarifaEspecial tarifaDestino

  let mensajeCobro = determinarMensajeCobro valorCobrado tarifaReal
  
  let informeCobro = generarInformeCobro valorCobrado tarifaReal mensajeCobro
  
  mostrarMensaje informeCobro

determinarTarifaMinima :: Int -> Int
determinarTarifaMinima valorTaximetro = 
  max valorTaximetro valorMinimo

determinarTarifaEspecial :: Bool -> Int
determinarTarifaEspecial esEspecial
  | esEspecial = valorEspecial
  | otherwise  = sinCosto

determinarTarifaDestino :: Int -> Int
determinarTarifaDestino destinoTrayecto
  | destinoTrayecto == periferia     = valorPeriferia
  | destinoTrayecto == extraUrbano   = valorExtraUrbano
  | destinoTrayecto == viaAeropuerto = valorViaAeropuerto
  | otherwise                        = sinCosto

calcularTarifa :: Int -> Int -> Int -> Int
calcularTarifa tarifaMinima tarifaEspecial tarifaDestino =
  tarifaMinima + tarifaEspecial + tarifaDestino

determinarMensajeCobro ::  Int -> Int -> String
determinarMensajeCobro valorCobrado tarifaReal
  | valorCobrado == tarifaReal = "JUSTO"
  | otherwise                  = "INJUSTO"

generarInformeCobro :: Int -> Int -> String -> String
generarInformeCobro valorCobrado tarifaReal mensajeCobro =
  printf ("\nValor cobrado \t$%7d" ++
    "\nValor real   \t$%7d"++
    "\n\nPor lo anterior el cobro es %s\n" )
    valorCobrado tarifaReal mensajeCobro