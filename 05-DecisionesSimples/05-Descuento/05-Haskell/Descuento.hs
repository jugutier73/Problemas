{-
   Crear un programa para aplicar un descuento a la tarifa del
   transporte público según el perfil del usuario, como estudiantes o
   adultos mayores. Considerando que los domingos y festivos la tarifa
   es mayor.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarLogico)

tarifaNormal, tarifaDomingoFestivo :: Int
tarifaNormal = 2900
tarifaDomingoFestivo = 3000

porcentajeDescuentoEstudiante, porcentajeDescuentoAdultoMayor :: Double
porcentajeDescuentoEstudiante = 10.0
porcentajeDescuentoAdultoMayor = 15.0

main :: IO()
main = do
  esDomingoFestivo <- ingresarLogico "Es domingo o festivo (s/n): "
  esEstudiante <- ingresarLogico "Es estudiante (s/n): "
  esAdultoMayor <- ingresarLogico "Es adulto mayor (s/n): "

  let tarifaDia = obtenerTarifaDia esDomingoFestivo

  let porcentajeDescuento = obtenerPorcentajeDescuento esEstudiante  esAdultoMayor

  let valorDescuento = calcularValorDescuento tarifaDia porcentajeDescuento

  let valorTarifa = calcularValorTarifa tarifaDia valorDescuento

  let reciboTarifa = generarReciboTarifa tarifaDia valorTarifa porcentajeDescuento valorDescuento
  
  mostrarMensaje reciboTarifa

obtenerTarifaDia :: Bool -> Int
obtenerTarifaDia esDomingoFestivo =
    if esDomingoFestivo 
      then tarifaDomingoFestivo 
      else tarifaNormal

obtenerPorcentajeDescuento :: Bool -> Bool -> Double
obtenerPorcentajeDescuento esEstudiante esAdultoMayor
    | esAdultoMayor = porcentajeDescuentoAdultoMayor
    | esEstudiante  = porcentajeDescuentoEstudiante
    | otherwise     = 0.0

calcularValorDescuento :: Int -> Double -> Double
calcularValorDescuento tarifaDia porcentajeDescuento =
    fromIntegral tarifaDia * porcentajeDescuento / 100.0

calcularValorTarifa :: Int -> Double -> Double
calcularValorTarifa tarifaDia valorDescuento =
    fromIntegral tarifaDia - valorDescuento

generarReciboTarifa :: Int -> Double -> Double -> Double -> String
generarReciboTarifa tarifaDia valorTarifa porcentajeDescuento valorDescuento =
    let mensajeBase = printf "\naLa tarifa es de %d\n" tarifaDia
    in if valorDescuento > 0
       then printf ("\n %s" ++
            "\nsu tarifa a pagar es de $%.0f" ++ 
            " por tener" ++
            "\nun descuento del %.1f%%" ++
            " equivalente a $%.0f\n")
            mensajeBase
            valorTarifa
            porcentajeDescuento
            valorDescuento
       else mensajeBase
