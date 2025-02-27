{-
   Programa para la empresa de seguridad Aros S.A.
   que permita resaltar el nombre del empleado destacado
   del mes ofreciendo unas felicitaciones públicas.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje)

main :: IO()
main = do
  nombreEmpleado <- ingresarTexto "Nombre del empleado destacado: "
  nombreMes <- ingresarTexto "Nombre del mes: "
  let felicitaciones = generarFelicitaciones nombreEmpleado nombreMes
  mostrarMensaje felicitaciones

ingresarTexto :: String -> IO String
ingresarTexto pregunta = do
  mostrarMensaje pregunta
  getLine

generarFelicitaciones :: String -> String -> String
generarFelicitaciones nombreEmpleado nombreMes =
  printf
    ("\nLa empresa de seguridad Aros S.A. quiere felicitar" ++
    "\npúblicamente a %s como nuestro" ++
    "\nempleado destacado del mes de %s," ++
    "\nmuchas gracias por su excelencia.")
    nombreEmpleado
    nombreMes