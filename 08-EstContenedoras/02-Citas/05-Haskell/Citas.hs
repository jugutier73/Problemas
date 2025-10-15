{-
   Crear un programa para garantizar que los pacientes sean atendidos
   de manera equitativa y oportuna. El orden de llegada debe
   equilibrarse con la prioridad médica de cada caso

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, 
                    ingresarTexto, 
                    ingresarOpcion, 
                    ingresarColeccion, 
                    ordenarColeccion, 
                    convertirColeccionCadena)
import Data.List (sortBy, maximumBy)
import Data.Function (on)
import Data.Ord (comparing)

nivel_riesgo, nivel_urgencia, nivel_prioritario :: Int
niveles = 3
nivel_riesgo      = 1
nivel_urgencia    = 2
nivel_prioritario = 3

data Cita = Cita { nombre :: String, nivel :: Int } deriving (Show)

main :: IO()
main = do
  citas <- ingresarColeccion ingresarCita
  
  let citasPorNivel       = ordenarColeccion citas compararNivel False

  let cantidadRiesgo      = contarSegunCriterio citas tenerNivelRiesgo nivel_riesgo 

  let cantidadUrgencia    = contarSegunCriterio citas tenerNivelRiesgo nivel_urgencia
  
  let cantidadPrioritario = contarSegunCriterio citas tenerNivelRiesgo nivel_prioritario
  
  let reporteCitas  = generarReporteCitas 
                        citasPorNivel 
                        cantidadRiesgo 
                        cantidadUrgencia  
                        cantidadPrioritario  

  mostrarMensaje reporteCitas

ingresarCita :: IO Cita
ingresarCita = do
  mostrarMensaje "\nIngrese los datos de la cita:\n"
  nombre <- ingresarTexto  "\tIngrese el nombre paciente: "
  nivel <- ingresarOpcion ("\tNIVEL DE URGENCIA\n"
                            ++ "\t\t1: Riesgo\n"
                            ++ "\t\t2: Urgencia\n"
                            ++ "\t\t3: Prioritario\n"
                            ++ "\tIngrese tipo de persona: ") 
                          niveles

  return (Cita nombre nivel)

compararNivel :: Cita -> Cita -> Ordering
compararNivel = comparing nivel

contarSegunCriterio :: [tipo1] -> (tipo1 -> tipo2 -> Bool) -> tipo2 -> Int
contarSegunCriterio coleccion aplicarCriterio valorCriterio = 
  length (filter (\c -> (aplicarCriterio c valorCriterio) ) coleccion)

tenerNivelRiesgo :: Cita -> Int -> Bool
tenerNivelRiesgo cita nivelInteres = 
  (nivel cita) == nivelInteres

generarReporteCitas :: [Cita] -> Int -> Int -> Int -> String
generarReporteCitas citasPorNivel cantidadRiesgo cantidadUrgencia cantidadPrioritario =
  let listadoPorNombre = convertirColeccionCadena 
          "LISTADO POR NIVEL DE URGENCIA" 
          citasPorNivel 
          convertirCitaCadena
  in printf ("%s\n\n"
              ++ "Nivel 1 (Riesgo)     : %d\n" 
              ++ "Nivel 2 (Urdencia)   : %d\n"
              ++ "Nivel 3 (Prioritario): %d\n")
              listadoPorNombre cantidadRiesgo cantidadUrgencia cantidadPrioritario

convertirCitaCadena :: Cita -> String
convertirCitaCadena cita = 
  printf ("%d : %s\n") (nivel cita) (nombre cita)