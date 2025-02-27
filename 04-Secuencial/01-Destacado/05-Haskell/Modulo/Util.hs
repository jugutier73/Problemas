{-
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
-}

module Modulo.Util (mostrarMensaje) where

mostrarMensaje :: String -> IO ()
mostrarMensaje mensaje = do
  putStr mensaje