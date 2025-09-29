{-
    Programa para calcular la cantidad de CO2 emitido por el uso de
    transporte particular (carro y moto), el empleo del transporte 
    público (buses)
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarReal)

huellaCarro, huellaMoto, huellaBus :: Double
huellaCarro = 121.0
huellaMoto = 53.0
huellaBus = 49.0

main :: IO()
main = do
  kmCarro <- ingresarReal "Total de kilómetros en carro: "
  kmMoto <- ingresarReal "Total de kilómetros en moto : "
  kmBuses <- ingresarReal "Total de kilómetros en buses: "

  let huella = calcularHuellaCarbono kmCarro kmMoto kmBuses
  let informeHuella = generarHuella kmCarro kmMoto kmBuses huella
  
  mostrarMensaje informeHuella

calcularHuellaCarbono :: Double -> Double -> Double -> Double
calcularHuellaCarbono kmCarro kmMoto kmBuses =
  kmCarro * huellaCarro + kmMoto * huellaMoto + kmBuses * huellaBus

generarHuella :: Double -> Double -> Double -> Double -> String
generarHuella kmCarro kmMoto kmBuses huella =
  printf
    ("\nCon %.1f, %.1f, %.1f km de recorrido" ++
    "\nen carro, moto y bus representante,"++
    "\nsu huella de carbono por el uso de"++
    "\ntransporte es de %.1f kg de CO2.\n"
    )
    kmCarro
    kmMoto
    kmBuses
    huella