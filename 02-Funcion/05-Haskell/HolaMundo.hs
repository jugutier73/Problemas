{-
    Programa que muestra el mensaje "Hola Mundo" en la pantalla
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Enero 2025
    Licencia: GNU GPL v3
-}

main :: IO()
main = do
  mostrarMensaje "Hola Mundo\n"

mostrarMensaje :: String -> IO ()
mostrarMensaje mensaje = do
  putStr mensaje