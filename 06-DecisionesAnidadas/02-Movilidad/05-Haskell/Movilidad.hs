{-
    Crear un programa para recomendar un medio de transporte
    según el tipo de distancia a la universidad, si está lloviendo,
    y si hay o no transporte público.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Julio 2025
    Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarLogico, ingresarOpcion)

cerquita, lejos :: Int
cerquita = 1
lejos    = 3

main :: IO()
main = do
  tipoDistancia <- ingresarOpcion ("Vive\n" ++
    "\t(1) cerquita,\n" ++
    "\t(2) cerca,\n" ++
    "\t(3) lejos: \n\n" ++
    "Cuál es el tipo de distancia: ") 3

  estaLloviendo <- ingresarLogico "Está lloviendo (s/n): "

  hayTransporte <- ingresarLogico "Hay transporte público (s/n): "
  
  let medioTransporte = recomendarMedioTransporte tipoDistancia estaLloviendo hayTransporte
  
  let reporteTransporte = generarReporteTransporte medioTransporte
  
  mostrarMensaje reporteTransporte

recomendarMedioTransporte :: Int -> Bool -> Bool -> String
recomendarMedioTransporte tipoDistancia estaLloviendo hayTransporte
    | tipoDistancia == cerquita && not estaLloviendo = "caminar o usar bicicleta"
    | tipoDistancia == lejos    || not hayTransporte = "carro compartido"
    | otherwise                                      = "transporte público"

generarReporteTransporte :: String -> String
generarReporteTransporte medioTransporte =
  printf ("\nMedio de transporte recomendado: %s\n") medioTransporte