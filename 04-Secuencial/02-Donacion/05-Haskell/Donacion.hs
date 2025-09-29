{-
   Programa para la iniciativa Amigo Social que permita generar
   el documento de recibido que sirva como soporte contable de
   la donación.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarTexto)

main :: IO()
main = do
  nombreCasa <- ingresarTexto "Nombre casa adulto mayor: "
  cantidadRecolectada <- ingresarEntero "Cantidad recolectada: "

  let reciboDonacion = generarRecibo nombreCasa cantidadRecolectada
  
  mostrarMensaje reciboDonacion
  
ingresarEntero :: String -> IO Int
ingresarEntero pregunta = read <$> ingresarTexto pregunta     
  
generarRecibo :: String -> Int -> String
generarRecibo nombreCasa cantidadRecolectada =
  printf
    ("\nLa iniciativa Amigo Social tiene el gusto de entregar" ++
    "\nuna donación de $%d pesos colombianos" ++
    "\na la casa del adulto mayor %s." ++
    "\n\n_______________________" ++
    "\nFirma representante legal\n")
    cantidadRecolectada
    nombreCasa