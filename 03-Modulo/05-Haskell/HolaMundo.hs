{-
    Programa que muestra el mensaje "Hola Mundo" en la pantalla
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Febrero 2025
    Licencia: GNU GLP v3
-}

import Modulo.Util (mostrarMensaje)

main :: IO()
main = do
  mostrarMensaje "Hola Mundo"