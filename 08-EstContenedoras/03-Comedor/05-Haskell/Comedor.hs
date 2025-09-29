{-
   Crear un programa para organizar la entrada de los beneficiarios
   y garantice la atención en orden de llegada. Se debe registrar
   datos clave (nombre, edad y presencia de necesidades especiales).
   Además de informar la cantidad de personas con necesidades y el
   promedio de edades.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, 
                    ingresarTexto, 
                    ingresarEntero, 
                    ingresarLogico,
                    ingresarColeccion, 
                    convertirColeccionCadena,
                    contarSegunCriterio)


data Reserva = Reserva { nombre :: String, edad :: Int, necesidadEspecial :: Bool } deriving (Show)

main :: IO()
main = do
  reservasComedor <- ingresarColeccion ingresarReverva
  
  let cantidadConNecesidades = contarSegunCriterio reservasComedor tenerNecesidadEspecial True

  let promedioEdades = calcularPromedioEdades reservasComedor
  
  let reporteReservas = generarReporteReservasComedor 
                          reservasComedor 
                          cantidadConNecesidades 
                          promedioEdades

  mostrarMensaje reporteReservas

ingresarReverva :: IO Reserva
ingresarReverva = do
  mostrarMensaje "\nIngrese los datos de la reserva:\n"
  nombre <- ingresarTexto  "\tIngrese el nombre de la persona   : "
  edad <- ingresarEntero "\tIngrese la edad de la persona     : "
  necesidadEspecial <- ingresarLogico "\tTiene necesidades especiales (s/n): "

  return (Reserva nombre edad necesidadEspecial)

tenerNecesidadEspecial :: Reserva -> Bool -> Bool
tenerNecesidadEspecial reserva necesidadEspecialInteres =
  (necesidadEspecial reserva) == necesidadEspecialInteres

calcularPromedioEdades :: [Reserva] -> Double
calcularPromedioEdades reservasComedor
  | totalReservas == 0  = 0
  | otherwise = fromIntegral suma / fromIntegral totalReservas
  where
    totalReservas = length reservasComedor
    suma          = sum (map edad reservasComedor)

generarReporteReservasComedor :: [Reserva] -> Int -> Double -> String
generarReporteReservasComedor reservasComedor cantidadConNecesidades promedioEdades =
  let listadoPorNombre = convertirColeccionCadena 
          "LISTADO DE RESERVAS" 
          reservasComedor 
          convertirReservaCadena
  in printf ("%s\n\n"
              ++ "Cantidad con necesidades : %d\n" 
              ++ "Promedio de edades       : %.1f\n")
              listadoPorNombre cantidadConNecesidades promedioEdades

convertirReservaCadena :: Reserva -> String
convertirReservaCadena reserva =
  printf ("%s, %d años%s\n")
    (nombre reserva)
    (edad reserva)
    (if necesidadEspecial reserva then ", con necesidad especial" else "")