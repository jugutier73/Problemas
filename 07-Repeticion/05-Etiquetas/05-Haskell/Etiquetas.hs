{-
   Programa para verificar si una etiqueta de reciclaje es válida
   y reportar los campos con error.

   Formato: T-aaaa-mm
   - T: tipo de residuo
     (P: plástico, V: vidrio, M: metal, C: cartón/papel, O: orgánico)
   - aaaa: año (> 2025)
   - mm: mes (01-12)

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Text.Read   (readMaybe)
import Modulo.Util (mostrarMensaje, ingresarTexto)
import Data.Bits ((.|.), (.&.))

separador :: Char
separador = '-'

longitudEtiqueta, posicionSeparador1, posicionSeparador2 :: Int
posicionTipo, posicionAnio, longitudAnio, posicionMes, longitudMes :: Int
minimoAnio, minimoMes, maximoMes :: Int
longitudEtiqueta   = 9

posicionSeparador1 = 1
posicionSeparador2 = 6

posicionTipo       = 0

posicionAnio       = 2
longitudAnio       = 4

posicionMes        = 7
longitudMes        = 2

minimoAnio         = 2026
minimoMes          = 1
maximoMes          = 12

tiposValidos :: String
tiposValidos = "PVMCO"

etiquetaOK, errorLongitud, errorFormato, errorTipo, errorAnio, errorMes :: Int
etiquetaOK    = 0
errorLongitud = 1
errorFormato  = 2
errorTipo     = 4
errorAnio     = 8
errorMes      = 16

main :: IO()
main = do
    etiqueta <- ingresarTexto "Ingrese la etiqueta a analizar:  "
    
    let codigo = verificarEtiquetaReciclaje etiqueta
    
    let reporteEtiqueta = generarReporteEtiqueta etiqueta codigo
    
    mostrarMensaje reporteEtiqueta

verificarEtiquetaReciclaje :: String -> Int
verificarEtiquetaReciclaje etiqueta =
  let codigo = verificarLongitud etiqueta
  in if codigo /= errorLongitud
        then
          verificarFormato etiqueta .|. 
          verificarTipo etiqueta    .|. 
          verificarAnio etiqueta    .|. 
          verificarMes etiqueta
        else codigo

verificarLongitud :: String -> Int
verificarLongitud etiqueta =
  if length etiqueta /= longitudEtiqueta 
    then errorLongitud
    else etiquetaOK 

verificarFormato :: String -> Int
verificarFormato etiqueta
  | head (drop posicionSeparador1 etiqueta) /= separador  ||
    head (drop posicionSeparador2 etiqueta) /= separador  = errorFormato
  | otherwise = etiquetaOK

verificarTipo :: String -> Int
verificarTipo etiqueta
  | elem (head etiqueta) tiposValidos = etiquetaOK
  | otherwise             = errorTipo

verificarAnio :: String -> Int
verificarAnio etiqueta =
  let anioEtiqueta = take longitudAnio (drop posicionAnio etiqueta)
  in case readMaybe anioEtiqueta :: Maybe Int of
       Just anio | anio >= minimoAnio -> etiquetaOK
       _ -> errorAnio

verificarMes :: String -> Int
verificarMes etiqueta =
  let mesEtiqueta = take longitudMes (drop posicionMes etiqueta)
  in case readMaybe mesEtiqueta :: Maybe Int of
       Just mes | mes >= minimoMes && mes <= maximoMes -> etiquetaOK
       _ -> errorMes

generarReporteEtiqueta :: String -> Int -> String
generarReporteEtiqueta etiqueta codigo =
  printf "\nLa  etiqueta \"%s\" %s" etiqueta cuerpo
  where
    cuerpo
      | codigo == etiquetaOK = "es válida.\n"
      | otherwise = printf "es inválida por tener:\n%s" (unlines errores)

    errores =
      [ "  - Error en la longitud" | codigo .&. errorLongitud /= 0 ] ++
      [ "  - Error en el formato"  | codigo .&. errorFormato  /= 0 ] ++
      [ "  - Error en el tipo"     | codigo .&. errorTipo     /= 0 ] ++
      [ "  - Error en el año"      | codigo .&. errorAnio     /= 0 ] ++
      [ "  - Error en el mes"      | codigo .&. errorMes      /= 0 ]