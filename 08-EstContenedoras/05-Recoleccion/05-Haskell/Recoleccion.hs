{-
   Programa que registra los residuos recolectados durante una
   jornada, especificando su código, tipo y condición de
   reparabilidad. El sistema genera un informe con el
   listado de los elementos recolectados, ordenados por tipo y
   reparabilidad, además de presentar un conteo total por cada
   tipo de residuo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, 
                    ingresarTexto, 
                    ingresarOpcion, 
                    ingresarLogico,
                    ingresarColeccion, 
                    convertirColeccionCadena,
                    ordenarColeccion,
                    contarSegunCriterio)                   
import Data.Ord    (comparing)

ancho :: Int
ancho = 15

totalTipos :: Int
totalTipos = 5

bateria, telefono, computador, cargador, otroDispositivo :: Int
bateria = 1; 
telefono = 2; 
computador = 3; 
cargador = 4; 
otroDispositivo = 5

etiqueta :: Int -> String
etiqueta t = case t of
  1 -> "Batería"
  2 -> "Teléfono"
  3 -> "Computador"
  4 -> "Cargador"
  5 -> "Otro"
  _ -> "Desconocido"

data Residuo = Residuo { codigo :: String, tipo :: Int, reparable :: Bool } deriving (Show)

main :: IO()
main = do
  residuos <- ingresarColeccion ingresarResiduo
  
  let residuosClasificados = ordenarColeccion residuos compararCampoTipo False

  let cantidadBaterias = contarSegunCriterio residuos tenerTipo bateria
  
  let cantidadTelefonos = contarSegunCriterio residuos tenerTipo telefono
  
  let cantidadComputadores = contarSegunCriterio residuos tenerTipo computador
  
  let cantidadCargadores = contarSegunCriterio residuos tenerTipo cargador
  
  let cantidadOtros = contarSegunCriterio residuos tenerTipo otroDispositivo
  
  let reporteRecoleccion = generarReporteRecoleccion residuosClasificados cantidadBaterias cantidadTelefonos cantidadComputadores cantidadCargadores cantidadOtros

  mostrarMensaje reporteRecoleccion

ingresarResiduo :: IO Residuo
ingresarResiduo = do
  mostrarMensaje "\nIngrese los datos del residuo electrónico:\n"
  codigo <- ingresarTexto  "\tIngrese el código del residuo: "
  tipo <- ingresarOpcion ( "\tTIPOS DE RESIDUOS\n" 
                            ++ "\t\t1: Batería\n" 
                            ++ "\t\t2: Teléfono\n" 
                            ++ "\t\t3: Computador\n" 
                            ++ "\t\t4: Cargador\n" 
                            ++ "\t\t5: Otro dispositivo)\n"
                            ++ "\tIngrese tipo de residuo      : ")  
                            totalTipos

  reparable <- ingresarLogico "\t\tPuede ser reparado (s/n): "

  return (Residuo codigo tipo reparable)

compararCampoTipo :: Residuo -> Residuo -> Ordering
compararCampoTipo = comparing tipo <> comparing (not . reparable)

tenerTipo :: Residuo -> Int -> Bool
tenerTipo residuo tipoInteres =
  (tipo residuo) == tipoInteres

generarReporteRecoleccion :: [Residuo] -> Int -> Int -> Int -> Int -> Int -> String
generarReporteRecoleccion residuosClasificados cantidadBaterias cantidadTelefonos cantidadComputadores cantidadCargadores cantidadOtros =
  let listadoPorNombre = convertirColeccionCadena 
          "LISTADO DE RESIDUOS CLASIFICADOS" 
          residuosClasificados 
          convertirResiduoCadena

  in printf ("%s\n\n"
              ++ "Cantidad de Baterías     : %d\n" 
              ++ "Cantidad de Teléfonos    : %d\n" 
              ++ "Cantidad de Computadores : %d\n" 
              ++ "Cantidad de Cargadores   : %d\n" 
              ++ "Cantidad de Otros        : %d\n")
              listadoPorNombre cantidadBaterias cantidadTelefonos cantidadComputadores cantidadCargadores cantidadOtros

convertirResiduoCadena :: Residuo -> String
convertirResiduoCadena residuo =
  printf ("%-*s %-*s %s\n")
    ancho (codigo residuo)
    ancho (etiqueta (tipo residuo))
    (if reparable residuo then "(REPARABLE)" else "")
