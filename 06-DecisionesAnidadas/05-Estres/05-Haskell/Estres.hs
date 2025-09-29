{-
   Crear un programa para evaluar el nivel de estrés de los estudiantes

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarOpcion)

opciones, nivelBajo, nivelModerado :: Int
opciones      = 5
nivelBajo     = 19
nivelModerado = 25

main :: IO()
main = do
    respuesta01 <- ingresarRespuesta "¿con qué frecuencia te has sentido afectado por algo que ocurrió inesperadamente?"
    respuesta02 <- ingresarRespuesta "¿con qué frecuencia te has sentido incapaz de controlar las  cosas importantes en tu vida?"
    respuesta03 <- ingresarRespuesta "¿con qué frecuencia te has sentido nervioso o estresado?"
    respuesta04 <- ingresarRespuestaInvertida "¿con qué frecuencia has manejado con éxito los pequeños problemas irritantes de la vida?"
    respuesta05 <- ingresarRespuestaInvertida "¿con qué frecuencia has sentido que has afrontado efectivamente los cambios importantes que han estado ocurriendo en tu vida?"
    respuesta06 <- ingresarRespuestaInvertida "¿con qué frecuencia has estado seguro sobre tu capacidad para manejar tus problemas personales?"
    respuesta07 <- ingresarRespuestaInvertida "¿con qué frecuencia has sentido que las cosas van bien?"
    respuesta08 <- ingresarRespuesta "¿con qué frecuencia has sentido que no podías afrontar todas las cosas que tenías que hacer?"
    respuesta09 <- ingresarRespuestaInvertida "¿con qué frecuencia has podido controlar las dificultades de tu vida?"
    respuesta10 <- ingresarRespuestaInvertida "¿con qué frecuencia has sentido que tenías todo bajo control?"
    respuesta11 <- ingresarRespuesta "¿con qué  frecuencia has estado enfadado porque las cosas que te han ocurrido estaban fuera de tu control?"
    respuesta12 <- ingresarRespuesta "¿con qué frecuencia has pensado sobre las cosas que te faltan por hacer?"
    respuesta13 <- ingresarRespuestaInvertida "¿con qué frecuencia has podido controlar la forma de pasar el tiempo?"
    respuesta14 <- ingresarRespuesta "¿con qué frecuencia has sentido que las dificultades se acumulan tanto que no puedes superarlas?"

    let puntajeTotal = calcularPuntajeTotal respuesta01 respuesta02 respuesta03 respuesta04 respuesta05 respuesta06 respuesta07 respuesta08 respuesta09 respuesta10 respuesta11 respuesta12 respuesta13 respuesta14
    let nivelEstres = obtenerNivelEstres puntajeTotal
    let reporteEstres = generarReporteEstres nivelEstres puntajeTotal

    mostrarMensaje reporteEstres

ingresarRespuesta :: String -> IO Int
ingresarRespuesta pregunta = do
    -- Se restar uno para obtener una escala de 0 a 4
    opcion <- ingresarOpcion ( printf ("\nEn el último mes, %s\n\n" ++
        "\t(1) Nunca,\n" ++
        "\t(2) Casi nunca,\n" ++
        "\t(3) De vez en cuando\n" ++
        "\t(4) A menudo\n" ++
        "\t(5) Muy a menudo\n\n" ++
        "\tCuál es su frecuencia: ") pregunta) opciones
    return (opcion - 1)

ingresarRespuestaInvertida :: String -> IO Int
ingresarRespuestaInvertida pregunta = do
  opcion <- ingresarRespuesta pregunta
  return (abs (opcion - opciones + 1))

calcularPuntajeTotal :: Int -> Int -> Int -> Int -> Int -> 
   Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int -> 
   Int -> Int
calcularPuntajeTotal respuesta01 respuesta02 respuesta03 
         respuesta04 respuesta05 respuesta06 respuesta07 
         respuesta08 respuesta09 respuesta10 respuesta11 
         respuesta12 respuesta13 respuesta14 = 
  respuesta01 + respuesta02 + respuesta03 + respuesta04 + 
  respuesta05 + respuesta06 + respuesta07 + respuesta08 + 
  respuesta09 + respuesta10 + respuesta11 + respuesta12 + 
  respuesta13 + respuesta14

obtenerNivelEstres :: Int -> String
obtenerNivelEstres puntajeTotal
  | puntajeTotal < nivelBajo     = "BAJO"
  | puntajeTotal < nivelModerado = "MODERADO"
  | otherwise                    = "ALTO"

generarReporteEstres :: String -> Int -> String
generarReporteEstres nivelEstres puntajeTotal =
  printf ("\nSu nivel de estrés es %s" ++
          ", con un puntaje de %d.\n") nivelEstres puntajeTotal
