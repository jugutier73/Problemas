{-
    Programa para verificar la elegibilidad de los donantes 
    con base en la edad, el peso y la madurez fisiológica
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
-}

import Modulo.Util (mostrarMensaje, ingresarEntero, ingresarLogico, ingresarReal)

edad_minima, edad_maxima, edad_minima_autorizacion, edad_maxima_autorizacion :: Int
edad_minima = 18
edad_maxima = 65
edad_minima_autorizacion = 16
edad_maxima_autorizacion = 18

peso_minimo :: Double
peso_minimo = 50.0

main :: IO()
main = do
  edadDonante <- ingresarEntero "Edad del donante: "
  autorizacionEdad <- ingresarLogico "Tiene autorización por edad (s/n): "
  pesoDonante <- ingresarReal "Peso del donante: "
  suficienteMadurez <-ingresarLogico "Suficiente madurez fisiológica(s/n):"

  let reporteLegibilidad = generarReporteEligibilidad edadDonante autorizacionEdad pesoDonante suficienteMadurez  

  mostrarMensaje reporteLegibilidad

generarReporteEligibilidad :: Int -> Bool -> Double -> Bool -> String
generarReporteEligibilidad edadDonante autorizacionEdad pesoDonante suficienteMadurez =
    if (edadDonante >= edad_minima && edadDonante <= edad_maxima ||
      edadDonante >= edad_minima_autorizacion &&
      edadDonante <= edad_maxima_autorizacion && autorizacionEdad) &&
      pesoDonante > peso_minimo && suficienteMadurez
    then "\nEl donante es elegible para donar\n"
    else "\nEl donante no cumple las condiciones para donar\n"