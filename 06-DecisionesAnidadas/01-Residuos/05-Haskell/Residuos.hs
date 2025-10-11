{-
   Crear un programa para a clasificación de residuos según su
   material. Residuos aprovechable (limpios y secos), usar bolsa
   blanca; residuos orgánicos, usar bolsa verde; lo anterior si
   existe ruta de recolección selectiva; en otro caso, usar bolsa
   de color negro.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarEntero, 
                    ingresarLogico, mostrarError)

aprovechable, otro, max_opciones :: Int
aprovechable = 1
otro         = 3
max_opciones = 3

main :: IO()
main = do
  tipoResiduo <- ingresarOpcion ("Residuo \n"++
    "\t(1) aprovechable (limpia y seca),\n"++
    "\t(2) orgánico,\n"++
    "\t(3) otro\n\n"++
    "Cuál es el tipo de residuo: ") max_opciones

  hayRecoleccionSelectiva <- ingresarLogico "Hay ruta de recolección selectiva (s/n): "
  
  let colorBolsa = recomendarColorBolsa tipoResiduo hayRecoleccionSelectiva
  
  let reporteBolsaRecoleccion = generarReporteBolsa colorBolsa
  
  mostrarMensaje reporteBolsaRecoleccion

ingresarOpcion :: String -> Int -> IO Int
ingresarOpcion pregunta maximaOpcion = do
  valor <- ingresarEntero pregunta
  opcion <- if valor < 1 || valor > maximaOpcion
     then mostrarError "La opción no es válida, se asume 1\n\n" >> pure 1
     else pure valor
  return opcion

recomendarColorBolsa :: Int -> Bool  -> String
recomendarColorBolsa tipoResiduo hayRecoleccionSelectiva
    | tipoResiduo == otro || not hayRecoleccionSelectiva = "NEGRA"
    | tipoResiduo == aprovechable                        = "BLANCA"
    | otherwise                                          = "VERDE"

generarReporteBolsa :: String -> String
generarReporteBolsa colorBolsa =
  printf ("\nLa bolsa recomendada es de color %s\n") colorBolsa