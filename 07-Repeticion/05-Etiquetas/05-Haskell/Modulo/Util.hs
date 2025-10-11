{-
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
-}

module Modulo.Util (mostrarMensaje, mostrarError, ingresarTexto, ingresarEntero, ingresarReal, ingresarLogico, ingresarOpcion) where

import Data.Char (toLower)
import System.IO (stderr, hPutStr)
import Text.Read (readMaybe)

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
  Devuelve el entero ingresado por el usuario como respuesta a una pregunta

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
  Devuelve el booleano ingresado por el usuario como respuesta a una 
  pregunta.

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
  Devuelve el entero ingresado por el usuario como respuesta a una pregunta
  con múltiples opciones del 1 a un máximo especificado.
  
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