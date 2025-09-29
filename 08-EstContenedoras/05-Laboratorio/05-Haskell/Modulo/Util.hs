{-
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
-}

module Modulo.Util (mostrarMensaje, mostrarError, ingresarTexto, ingresarEntero, ingresarReal, ingresarLogico, ingresarOpcion, ingresarColeccion, ordenarColeccion, convertirColeccionCadena, contarSegunCriterio) where

import Data.Char (toLower)
import System.IO (stderr, hPutStr)
import Text.Read (readMaybe)
import Data.List (sortBy)
import Text.Printf (printf)

{-|
   Muestra un mensaje (cadena) en la salida estandar.

   @param mensaje Mensaje que se desea mostrar.
-}
mostrarMensaje :: String -> IO ()
mostrarMensaje mensaje = do
  putStr mensaje


{-|
  Muestra un mensaje (cadena) en el error estandar.
  
  @param mensaje Mensaje que se desea mostrar como error.
-}
mostrarError :: String -> IO ()
mostrarError mensaje = do
  hPutStr stderr mensaje


{-|
  Devuelve el texto ingresado por el usuario como respuesta a una pregunta.

  @param pregunta Texto que se le presenta al usuario como pregunta.
  @return Texto ingresado por el usuario
-}
ingresarTexto :: String -> IO String
ingresarTexto pregunta = do
  mostrarMensaje pregunta
  getLine


{-|
  Devuelve el entero ingresado por el usuario como respuesta a una pregunta.

  @param pregunta Texto que se le presenta al usuario como pregunta.
  @return Valor ingresado por el usuario.
-}
ingresarEntero :: String -> IO Int
ingresarEntero pregunta = do
  texto <- ingresarTexto pregunta
  entero <- case readMaybe texto :: Maybe Int of
            Just valor  -> pure valor
            Nothing     -> do
              mostrarError "El valor ingresado es inválido, intente nuevamente.\n\n"
              ingresarEntero pregunta
  return entero


{-|
  Devuelve el real ingresado por el usuario como respuesta a una pregunta.

  @param pregunta Texto que se le presenta al usuario como pregunta.
  @return Valor ingresado por el usuario.
-}
ingresarReal :: String -> IO Double
ingresarReal pregunta = do
  texto <- ingresarTexto pregunta
  real <- case readMaybe texto :: Maybe Double of
            Just valor  -> pure valor
            Nothing     -> do
              mostrarError "El valor ingresado es inválido, intente nuevamente.\n\n"
              ingresarReal pregunta
  return real


{-|
  Devuelve el booleano ingresado por el usuario como respuesta a una pregunta.

  @param pregunta Texto que se le presenta al usuario como pregunta.
  @return Valor ingresado por el usuario.
-}
ingresarLogico :: String -> IO Bool
ingresarLogico pregunta = do
    texto <- ingresarTexto pregunta
    let opcion = map toLower texto

    if opcion == "s" || opcion == "n"
        then pure (opcion == "s")
        else do
          putStrLn "No es una opción válida, intente nuevamente.\n\n"
          ingresarLogico pregunta


{-|
  Devuelve el entero ingresado por el usuario como respuesta a una pregunta con múltiples opciones del 1 a un máximo especificado.
  
  @param pregunta Texto que se le presenta al usuario como pregunta.
  @param maximaOpcion Entero que indica la cantidad de opciones.
  @return Valor ingresado por el usuario.
-}
ingresarOpcion :: String -> Int -> IO Int
ingresarOpcion pregunta maximaOpcion = do
  valor <- ingresarEntero pregunta
  if valor >= 1 && valor <= maximaOpcion
     then return valor
     else do
       mostrarError "La opción no es válida, intente nuevamente.\n\n"
       ingresarOpcion pregunta maximaOpcion


{-|
  Devuelve una colección, con los elementos obtenidos de un función especificada, hasta que el usuario indique que ya no sea ingresar más datos.
  
  @param ingresarElemento Función que solicita un dato al usuario.
  @return Coleccion con los valores ingresados por medio de la función indicada.
-}
ingresarColeccion :: IO tipo -> IO [tipo]
ingresarColeccion ingresarElemento = adicionarElemento []
  where
    adicionarElemento coleccion = do
      elemento <- ingresarElemento

      hayMas <- ingresarLogico "Hay más datos (s/n): "
      if hayMas 
      then adicionarElemento (coleccion ++ [elemento]) 
      else pure  (coleccion ++ [elemento])



{-|
  Ordena una coleccion de elementos, según el campo y orden indicado.
  
  @param coleccion Colección de elementos que se desean ordenar.
  @param comparador Función utilizada como criterio de ordenación.
  @param descendente Valor booleano para indicar si se ordena descendentemente.
  @return Nueva colección con todos los datos ordenados por el criterio y sentido indicado.
-}
ordenarColeccion :: [a] -> (a -> a -> Ordering) -> Bool -> [a]
ordenarColeccion xs cmp descendente
  | descendente = sortBy (flip cmp) xs
  | otherwise   = sortBy cmp xs


{-|
   Convierte una coleccion de elementos a una cadena, según una función dada como argumento que convierte un elemento.
   
   @param titulo Título que se utiliza para los elementos de la colección.
   @param coleccion Colección de elementos a mostrar.
   @param convertirElementoCadena Función que convierte un elemento a una cadena.
   @return Un texto con un título y todos los elementos de la colección.
-}
convertirColeccionCadena :: String -> [tipo] -> (tipo -> String) -> String
convertirColeccionCadena titulo lista convertirElementoCadena =
  printf ("\n%s\n%s") titulo 
    (concatMap 
      (\elemento -> "\t" ++ convertirElementoCadena elemento) lista
    )


{-|
   Cuenta la cantidad de elementos de una colección que cumplen un criterio dado como argumento.
   
   @param coleccion Colección de elementos a procesar.
   @param aplicarCriterio Función que toma un elemento e indica si cumple un cierto criterio.
   @param valorCriterio  Valor de refencia para el criterio.
   @return Cantidad de elementos de la colección que cumplemen el 
   el criterio indicado.
-}
contarSegunCriterio :: [tipo1] -> (tipo1 -> tipo2 -> Bool) -> tipo2 -> Int
contarSegunCriterio coleccion aplicarCriterio valorCriterio = 
  length (filter (\c -> (aplicarCriterio c valorCriterio) ) coleccion)