{-
   Crear un programa para determinar si hay problemas con el uso
   de los espacios en un frase que el usuario ingrese

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarTexto)

main :: IO()
main = do
    frase <- ingresarTexto "Ingrese un texto con los mensajes a analizar: "
    
    let fraseCorregida = corregirUsoEspacios frase
    
    let reporte = generarReporteAlfabetizacion frase fraseCorregida
    
    mostrarMensaje reporte

corregirUsoEspacios :: String -> String
corregirUsoEspacios = (\(_,_,frase) -> frase) 
                     . foldl analizar (False, False, "")
  where
    analizar (inicioFrase, espacioPendiente, fraseCorregida) simbolo
      | simbolo == ' ' =
          if inicioFrase
            then (inicioFrase, True, fraseCorregida)    
            else (inicioFrase, espacioPendiente, fraseCorregida)
      | otherwise =
          let fraseCorregida' = if espacioPendiente 
                       then fraseCorregida ++ [' ', simbolo]
                       else fraseCorregida ++ [simbolo]
          in (True, False, fraseCorregida')

generarReporteAlfabetizacion :: String -> String -> String
generarReporteAlfabetizacion frase fraseCorregida =
    printf "\nLa frase \"%s\" hace un uso %s\n"
    frase
    (if frase == fraseCorregida
       then "CORRECTO de espacios."
       else printf "INCORRECTO de espacios,\nlo correcto es \"%s\"" fraseCorregida)