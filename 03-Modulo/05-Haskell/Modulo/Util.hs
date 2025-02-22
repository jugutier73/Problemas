{-
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GLP v3
-}

module Modulo.Util (mostrarMensaje) where

mostrarMensaje :: String -> IO ()
mostrarMensaje mensaje = do
  putStr mensaje