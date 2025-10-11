{-
   Crear un programa para registrar y consultar de manera ágil
   todas las donaciones recibidas. Esta debería ofrecer el total 
   recaudado, listados alfabéticos y descendentes por valor, 
   identificar a los aportantes que superen un umbral específico y 
   señalar al mayor donante.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarTexto, ingresarEntero, ingresarLogico)
import Data.List (sortBy, maximumBy)
import Data.Function (on)
import Data.Ord (comparing)

umbral :: Int
umbral = 10000

data Donante = Donante { nombre :: String, donacion :: Int } deriving (Show)

main :: IO()
main = do
  donantes <- ingresarColeccion ingresarDonante
  
  let donantesPorNombre = ordenarColeccion donantes compararNombre False
  let donantesPorDonacion = ordenarColeccion donantes compararDonacion True

  let mayoresDonantes = obtenerMayoresDonaciones donantes umbral 
  let mayorDonante    = obtenerMayorDonante donantes
  let sumaDonaciones  = obtenerSumaDonaciones donantes
  
  let reporteDonaciones  = generarReporteDonaciones 
                              donantesPorNombre 
                              donantesPorDonacion 
                              mayoresDonantes  
                              mayorDonante  
                              sumaDonaciones 

  mostrarMensaje reporteDonaciones

ingresarColeccion :: IO tipo -> IO [tipo]
ingresarColeccion ingresarElemento = adicionarElemento []
  where
    adicionarElemento coleccion = do
      elemento <- ingresarElemento

      hayMas <- ingresarLogico "Hay más datos (s/n): "
      if hayMas 
      then adicionarElemento (coleccion ++ [elemento]) 
      else pure  (coleccion ++ [elemento])

ingresarDonante :: IO Donante
ingresarDonante = do
  mostrarMensaje "\nIngrese los datos de un donante:\n"
  nombre <- ingresarTexto  "\tIngrese el nombre: "
  donacion <- ingresarEntero "\tIngrese donación : "

  return (Donante nombre donacion)

ordenarColeccion :: [a] -> (a -> a -> Ordering) -> Bool -> [a]
ordenarColeccion coleccion comparador descendente
  | descendente = sortBy (flip comparador) coleccion
  | otherwise   = sortBy comparador coleccion

compararNombre :: Donante -> Donante -> Ordering
compararNombre = comparing nombre

compararDonacion :: Donante -> Donante -> Ordering
compararDonacion = comparing donacion

obtenerMayoresDonaciones :: [Donante] -> Int -> [Donante]
obtenerMayoresDonaciones donantes limite = filter (\d -> donacion d > limite) donantes

obtenerMayorDonante :: [Donante] -> Donante
obtenerMayorDonante = maximumBy (on compare donacion)

obtenerSumaDonaciones :: [Donante] -> Int
obtenerSumaDonaciones = foldr ((+) . donacion) 0

generarReporteDonaciones :: [Donante] -> [Donante] -> [Donante] -> Donante -> Int -> String
generarReporteDonaciones donantesPorNombre donantesPorDonacion mayoresDonantes  mayorDonante  sumaDonaciones  =
  let listadoPorNombre = convertirColeccionCadena 
          "LISTADO EN ORDEN ALFABÉTICO" 
          donantesPorNombre 
          obtenerDonante

      listadoPorDonacion = convertirColeccionCadena 
          "LISTADO ORDENADO POR DONACIÓN" 
          donantesPorDonacion 
          obtenerDonante

      listadoMayores =  convertirColeccionCadena 
          ("LISTADO DONACIONES MAYORES A $" ++ show umbral) 
          mayoresDonantes 
          obtenerDonante

  in printf ("%s\n%s\n%s\n"
              ++ "\nEl mayor donante: %s" 
              ++ "\nTotal de donaciones $%d\n")
              listadoPorNombre listadoPorDonacion listadoMayores 
              (nombre mayorDonante) sumaDonaciones

convertirColeccionCadena :: String -> [tipo] -> (tipo -> String) -> String
convertirColeccionCadena titulo coleccion convertirElementoCadena =
  printf ("\n%s\n%s") titulo 
    (concatMap 
      (\elemento -> "\t" ++ convertirElementoCadena elemento) coleccion
    )

obtenerDonante :: Donante -> String
obtenerDonante donante = 
  printf ("$%10d \t %s\n") (donacion donante) (nombre donante)