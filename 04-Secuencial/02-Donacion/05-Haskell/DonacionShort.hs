main = do
  putStr "Nombre casa adulto mayor: "
  nombreCasa <- getLine

  putStr "Cantidad recolectada: "
  cantidadRecolectada :: Int <- readLn

  putStrLn (
    "\nLa iniciativa Amigo Social tiene el gusto de entregar" ++
    "\nuna donación de $" ++ show cantidadRecolectada ++ 
     " pesos colombianos" ++
    "\na la casa del adulto mayor " ++ nombreCasa ++ "." ++
    "\n\n_______________________" ++
    "\nFirma representante legal\n" )
