import Text.Printf (printf)

factorCarro, factorMoto, factorBicicleta :: Double
factorCarro = 121.0
factorMoto = 53.0
factorBicicleta = 3.0

factorBus :: Double
factorBus = 49.0

factorCarneRoja, factorPescado, factorVegetariana :: Double
factorCarneRoja = 7.19
factorPescado = 3.91
factorVegetariana = 3.81

main :: IO ()
main = do
  mostrarMensaje "KILÓMETROS RECORRIDOS EN VEHÍCULOS PARTICULARES\n"
  kmCarro <- ingresarReal "- carro: "
  kmMoto <- ingresarReal "- moto: "
  kmBicicleta <- ingresarReal "- bicicleta: "

  mostrarMensaje "\nKILÓMETROS RECORRIDOS EN TRANSPORTE URBANO\n"
  trayectosBus <- ingresarReal "- bus: "

  mostrarMensaje "\nCANTIDAD COMIDAS CON\n"
  comidasCarneRoja <- ingresarEntero "- carne roja: "
  comidasPescado <- ingresarEntero "- pescado: "
  comidasVegetarianas <- ingresarEntero "- vegetarianas: "

  let kgCo2Recorridos = calcularKgCo2Recorridos kmCarro kmMoto kmBicicleta
  
  let kgCo2Trayectos = calcularKgCo2Trayectos trayectosBus

  let kgCo2Comidas = calcularKgCo2Comidas comidasCarneRoja comidasPescado comidasVegetarianas

  let kgCo2Total = calcularKgCo2Total kgCo2Recorridos kgCo2Trayectos kgCo2Comidas

  let mensajeSalida = generarMensajeSalida kgCo2Recorridos kgCo2Trayectos kgCo2Comidas kgCo2Total

  mostrarMensaje mensajeSalida

mostrarMensaje :: String -> IO ()
mostrarMensaje mensajeSalida = do
  putStr mensajeSalida

ingresarEntero :: String -> IO Int
ingresarEntero pregunta = do
  mostrarMensaje pregunta
  readLn

ingresarReal :: String -> IO Double
ingresarReal pregunta = do
  mostrarMensaje pregunta
  readLn

calcularKgCo2Recorridos :: Double -> Double -> Double -> Double
calcularKgCo2Recorridos kmCarro kmMoto kmBicicleta =
  kmCarro * factorCarro
    + kmMoto * factorMoto
    + kmBicicleta * factorBicicleta

calcularKgCo2Trayectos :: Double -> Double
calcularKgCo2Trayectos trayectosBus =
  trayectosBus * factorBus

calcularKgCo2Comidas :: Int -> Int -> Int -> Double
calcularKgCo2Comidas comidasCarneRoja comidasPescado comidasVegetarianas =
  fromIntegral comidasCarneRoja * factorCarneRoja
    + fromIntegral comidasPescado * factorPescado
    + fromIntegral comidasVegetarianas * factorVegetariana

calcularKgCo2Total :: Double -> Double -> Double -> Double
calcularKgCo2Total kgCo2Recorridos kgCo2Trayectos kgCo2Comidas =
  kgCo2Recorridos
    + kgCo2Trayectos
    + kgCo2Comidas

generarMensajeSalida :: Double -> Double -> Double -> Double -> String
generarMensajeSalida kgCo2Recorridos kgCo2Trayectos kgCo2Comidas kgCo2Total =
  printf
    "\nSu huella de carbono en kg de CO2 por:\n\
    \- RECORRIDOS %.2f kg\n\
    \- TRAYECTOS %.2f kg\n\
    \- COMIDAS %.2f kg\n\
    \\nPara un TOTAL de %.2f kg de emisiones de CO2."
    kgCo2Recorridos
    kgCo2Trayectos
    kgCo2Comidas
    kgCo2Total
