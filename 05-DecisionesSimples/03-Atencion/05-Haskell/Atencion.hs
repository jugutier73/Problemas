{-
   Programa para establecer la prioridad de atención
   en los centros de salud para identificar a los pacientes con
   mayor riesgo o vulnerabilidad
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
-}

import Modulo.Util (mostrarMensaje, ingresarEntero, ingresarLogico)

edad_recien_nacido, edad_adulto_mayor :: Int
edad_recien_nacido = 1
edad_adulto_mayor = 60

main :: IO()
main = do
  edadPaciente <- ingresarEntero "Edad del paciente: "
  enfermedadCronica <- ingresarLogico "Enfermedad crónica (s/n): "
  estadoInmunosupresion <- ingresarLogico "Estado de inmunosupresión(s/n):"

  let reporteIngreso = generarReporteAtencion edadPaciente enfermedadCronica estadoInmunosupresion

  mostrarMensaje reporteIngreso

generarReporteAtencion :: Int -> Bool -> Bool -> String
generarReporteAtencion edadPaciente enfermedadCronica estadoInmunosupresion =
    if edadPaciente < edad_recien_nacido || 
      edadPaciente > edad_adulto_mayor || 
      enfermedadCronica || 
      estadoInmunosupresion
    then "\nEl paciente es de atención prioritaria\n"
    else "\nEl paciente es de atención general\n"