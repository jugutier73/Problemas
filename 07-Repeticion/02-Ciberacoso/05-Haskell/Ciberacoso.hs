{-
   Crear un programa para contar cuantas palabras comienzan con
   la letra "p" en un mensaje.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Data.Char (toLower, isSpace)
import Modulo.Util (mostrarMensaje, ingresarTexto)

letraInicio :: Char
letraInicio = 'p'

main :: IO()
main = do
    texto <- ingresarTexto "Ingrese un texto con los mensajes a analizar: "

    let cantidadPalabras = contarPalabrasInician texto letraInicio

    let reporteCiberacoso = generarReporteAcoso cantidadPalabras

    mostrarMensaje reporteCiberacoso

contarPalabrasInician :: String -> Char -> Int
contarPalabrasInician texto letraInicio = 
  length [palabra | palabra <- words (map toLower texto), 
                    head palabra == letraInicio]

generarReporteAcoso :: Int -> String
generarReporteAcoso cantidadPalabrasInteres =
  printf ("\nHay %d palabras " ++
          "que inician con la letra  \"%c\".\n") 
          cantidadPalabrasInteres  letraInicio


-- contarPalabrasInician :: String -> Char -> Int
-- contarPalabrasInician texto letra =
--   let s   = map toLower texto
--       prev = ' ' : s 
--   in length [() | (p,c) <- zip prev s, isSpace p && c == letra]