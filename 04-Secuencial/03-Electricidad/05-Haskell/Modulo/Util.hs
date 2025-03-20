{-
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
-}

module Modulo.Util (mostrarMensaje, ingresarTexto, ingresarEntero) where

mostrarMensaje :: String -> IO ()
mostrarMensaje mensaje = do
  putStr mensaje

ingresarTexto :: String -> IO String
ingresarTexto pregunta = do
  mostrarMensaje pregunta
  getLine

ingresarEntero :: String -> IO Int
ingresarEntero pregunta = do
  mostrarMensaje pregunta
  readLn