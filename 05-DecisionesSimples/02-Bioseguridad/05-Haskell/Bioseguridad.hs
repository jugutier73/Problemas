{-
    Crear un programa para programa para la verificación de medidas 
    de bioseguridad para el ingreso a un evento público.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
-}

import Data.Char (toLower)
import Modulo.Util (mostrarMensaje, ingresarTexto)

main :: IO()
main = do
  tieneVacunas <- ingresarLogico "Tiene vacunas (s/n): "
  restadosNegativos <- ingresarLogico "Pruebas negativas (s/n): "
  tieneSintomas <- ingresarLogico "Tiene síntomas (s/n): "

  let reporteIngreso = generarReporteIngreso tieneVacunas restadosNegativos tieneSintomas
  
  mostrarMensaje reporteIngreso

ingresarLogico :: String -> IO Bool
ingresarLogico pregunta = do
    texto <- ingresarTexto pregunta

    let opcion = map toLower texto

    if opcion /= "s" && opcion /= "n"
        then putStrLn "No es una opción válida, se asume \"n\"\n\n"
        else pure()
        
    return (opcion == "s")

generarReporteIngreso :: Bool -> Bool -> Bool -> String
generarReporteIngreso tieneVacunas restadosNegativos tieneSintomas =
    if tieneVacunas && restadosNegativos && not tieneSintomas
    then "\nLa persona puede ingresar al evento\n"
    else "\nLa persona no puede ingresar al evento\n"