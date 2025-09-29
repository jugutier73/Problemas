{-
   Crear un programa para mostrar una alerta cuando el índice de la
   calidad del aire (medido con algún instrumento) indique que el aire
   puede resultar perjudicial para la población.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
-}

import Modulo.Util (mostrarMensaje, ingresarReal)

main :: IO()
main = do
  indiceCalidadAire <- ingresarReal "Índice calidad del aire: "

  let alertaCalidadAire = generarAlertaCalidadAire indiceCalidadAire

  mostrarMensaje alertaCalidadAire

generarAlertaCalidadAire :: Double -> String
generarAlertaCalidadAire indiceCalidadAire =
  let mensaje =
        if indiceCalidadAire > 100.0
           then "\nEl aire puede presentar efectos sobre la salud\n"
           else "\nEl aire supone un riesgo bajo para la salud\n"
  in mensaje       