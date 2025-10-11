{-
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
-}

module Modulo.Util (mostrarMensaje, ingresarTexto, ingresarEntero, ingresarReal) where

{-|
   Muestra un mensaje (cadena) en la salida estandar.

   @param mensaje Mensaje que se desea mostrar.
-}
mostrarMensaje :: String -> IO ()
mostrarMensaje mensaje = do
  putStr mensaje

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
ingresarEntero pregunta = read <$> ingresarTexto pregunta 

{-|
  Devuelve el real ingresado por el usuario como respuesta a una pregunta.

  @param pregunta Texto que se le presenta al usuario como pregunta.
  @return Valor ingresado por el usuario.
-}
ingresarReal :: String -> IO Double
ingresarReal pregunta = read <$> ingresarTexto pregunta 