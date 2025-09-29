{-
   Crear un programa para registrar el ingreso y salida de los
   usuarios de un laboratorio para asegurar que solo personas
   autorizadas -estudiantes, docentes o personal técnico- accedan
   al espacio. El programa debe permitir identificar quiénes estaban 
   presentes y la cantidad por cada tipo.

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
import qualified Data.Set as Set

ancho :: Int
ancho = 20

totalTipos :: Int
totalTipos = 3

estudiante, docente, tecnico :: Int
estudiante = 1; 
docente = 2; 
tecnico = 3; 

etiqueta :: Int -> String
etiqueta t = case t of
  1 -> "Estudiante"
  2 -> "Docente"
  3 -> "Técnico"

data Ingreso = Ingreso { documento :: String, nombre :: String, tipo :: Int} deriving (Show)

main :: IO()
main = do
  ingresos <- ingresarColeccion ingresarIngreso
  salidas <- ingresarColeccion ingresarSalida
  
  let personasLaboratorio = obtenerPersonasLaboratorio ingresos salidas
  
  let cantidadEstudiantes = contarSegunCriterio personasLaboratorio tenerTipo estudiante

  let cantidadDocentes = contarSegunCriterio personasLaboratorio tenerTipo docente

  let cantidadTecnicos = contarSegunCriterio personasLaboratorio tenerTipo tecnico
  
  let reporteIngresos = generarReporteIngresos ingresos salidas personasLaboratorio cantidadEstudiantes cantidadDocentes cantidadTecnicos

  mostrarMensaje reporteIngresos

ingresarIngreso :: IO Ingreso
ingresarIngreso = do
  mostrarMensaje "\nIngrese los datos del ingreso:\n"
  documento <- ingresarTexto  "\tIngrese el documento   : "
  nombre <- ingresarTexto  "\tIngrese el nombre      : "
  tipo <- ingresarOpcion ( "\tTIPO DE PERSONAS AUTORIZADAS\n" 
                            ++ "\t\t1: Estudiante\n" 
                            ++ "\t\t2: Docente\n" 
                            ++ "\t\t3: Técnico\n" 
                            ++ "\tIngrese tipo de persona      : ")  
                            totalTipos

  return (Ingreso documento nombre tipo)

ingresarSalida :: IO String 
ingresarSalida = do
  mostrarMensaje "\nIngrese los datos de la salida:\n"
  documento <- ingresarTexto  "\tIngrese el documento   : "

  return documento

obtenerPersonasLaboratorio :: [Ingreso] -> [String] -> [Ingreso]
obtenerPersonasLaboratorio ingresos salidas =
  let salidasSet = Set.fromList salidas

  in filter (\ingreso -> documento ingreso `Set.notMember` salidasSet) ingresos

tenerTipo :: Ingreso -> Int -> Bool
tenerTipo ingreso tipoInteres = (tipo ingreso) == tipoInteres

generarReporteIngresos :: [Ingreso] -> [String] -> [Ingreso] -> Int -> Int -> Int -> String
generarReporteIngresos ingresos salidas personasLaboratorio cantidadEstudiantes cantidadDocentes cantidadTecnicos =
  let listadoIngresos = convertirColeccionCadena 
          "LISTADO DE INGRESOS" 
          ingresos 
          convertirIngresoCadena
          
      listadoSalidas = convertirColeccionCadena 
          "LISTADO DE SALIDAS" 
          salidas 
          convertirSalidaCadena
          
      listadoPersonasLaboratorio = convertirColeccionCadena 
          "LISTADO DE PERSONAS EN EL LABORATORIO" 
          personasLaboratorio 
          convertirIngresoCadena

  in printf ("%s\n\n"
              ++ "%s\n\n"
              ++ "%s\n\n"
              ++ "Cantidad de Estudiante : %d\n" 
              ++ "Cantidad de Docente    : %d\n" 
              ++ "Cantidad de Técnico    : %d\n")
              listadoIngresos listadoSalidas listadoPersonasLaboratorio cantidadEstudiantes cantidadDocentes cantidadTecnicos

convertirIngresoCadena :: Ingreso -> String
convertirIngresoCadena ingreso =
  printf ("%-*s %-*s %s\n")
    ancho (documento ingreso)
    ancho (nombre ingreso)
    (etiqueta (tipo ingreso))

convertirSalidaCadena :: String -> String
convertirSalidaCadena salida = salida