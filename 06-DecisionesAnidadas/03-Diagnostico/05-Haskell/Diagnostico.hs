{-
   Crear un programa para registrar síntomas y recibir recomendaciones 
   simples y seguras como reposar, hidratarse o consultar a un profesional 
   de la salud, según la gravedad o combinación de síntomas reportados.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
-}

import Text.Printf (printf)
import Modulo.Util (mostrarMensaje, ingresarReal, ingresarLogico)

fiebreAlta, fiebre :: Double
fiebreAlta = 39.5
fiebre     = 37.5

mensajeEspecialista, mensajeFiebreAlta, mensajeFiebre, mensajeDolorCabeza, mensajeCongestion :: String
mensajeEspecialista = "\n\
 \- Consultar un especialista.\n\
 \- Anotar los síntomas y cuándo comenzaron.\n\
 \- Evitar esfuerzos físicos y actividades intensas."

mensajeFiebreAlta = "\n\
 \- Solicitar una cita medica con urgencia.\n\
 \- Usar paños húmedos y fríos en la frente.\n\
 \- Permanecer en un lugar fresco y ventilado."

mensajeFiebre       = "\n\
 \- Descansar lo suficiente.\n\
 \- Hidratarse bebiendo agua u otros líquidos."

mensajeDolorCabeza  = "\n\
 \- Realizar ejercicio cervical isométrico.\n\
 \- Descansar en un lugar oscuro y silencioso.\n\
 \- Beber agua, ya que la deshidratación puede empeorar el dolor."

mensajeCongestion   = "\n\
 \- Realizar lavados nasales con solución salina.\n\
 \- Usar almohadas extras para dormir con la cabeza elevada."

main :: IO()
main = do
  temperatura <- ingresarReal "Temperatura corporal: "
  sintomasVariosDias <- ingresarLogico "Síntomas por 2 o más días (s/n): "
  malestarIntenso <- ingresarLogico "Tiene malestar intenso (s/n): "
  dolorCabeza <- ingresarLogico "Tiene dolor de cabeza (s/n): "
  congestionNasal <- ingresarLogico "Tiene congestion nasal (s/n): "
  
  let recomendacionesEspecialista = recibirRecomendaciones (sintomasVariosDias || malestarIntenso)  mensajeEspecialista

  let recomendacionesFiebreAlta = recibirRecomendaciones (temperatura >= fiebreAlta) mensajeFiebreAlta

  let recomendacionesFiebre = recibirRecomendaciones (temperatura >= fiebre) mensajeFiebre

  let recomendacionesDolorCabeza = recibirRecomendaciones dolorCabeza mensajeDolorCabeza

  let recomendacionesCongestion = recibirRecomendaciones congestionNasal mensajeCongestion
  
  let recomendaciones = generarReporteRecomendaciones recomendacionesEspecialista recomendacionesFiebreAlta recomendacionesFiebre recomendacionesDolorCabeza recomendacionesCongestion
    
  mostrarMensaje recomendaciones

recibirRecomendaciones :: Bool -> String -> String
recibirRecomendaciones condicion recomendaciones
  | condicion = recomendaciones 
  | otherwise = ""

generarReporteRecomendaciones :: String -> String -> String -> String -> String -> String 
generarReporteRecomendaciones recomendacionesEspecialista recomendacionesFiebreAlta recomendacionesFiebre recomendacionesDolorCabeza recomendacionesCongestion
  | recomendacionesEspecialista /= "" =
     base ++ recomendacionesEspecialista
  | otherwise =
      base
      ++ recomendacionesFiebreAlta
      ++ recomendacionesFiebre
      ++ recomendacionesDolorCabeza
      ++ recomendacionesCongestion
  where
    base = "\nSe recomienda:" ++ "\n"
