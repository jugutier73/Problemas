/*
    Programa para verificar la elegibilidad de los donantes 
    con base en la edad, el peso y la madurez fisiológica
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
*/

const EDAD_MINIMA = 18;
const EDAD_MAXIMA = 65;
const EDAD_MINIMA_AUTORIZACION = 16;
const EDAD_MAXIMA_AUTORIZACION = 18;
const PESO_MINIMO = 50.0;

function main() {
  let edadDonante = ingresarEntero('edadDonante');
  let autorizacionEdad = ingresarLogico('autorizacionEdad');
  let pesoDonante = ingresarReal('pesoDonante');
  let suficienteMadurez = ingresarLogico('suficienteMadurez');

  let reporteLegibilidad = generarReporteEligibilidad(edadDonante, autorizacionEdad, pesoDonante, suficienteMadurez);  

  mostrarMensaje(reporteLegibilidad);
}

function generarReporteEligibilidad(edadDonante, autorizacionEdad, pesoDonante, suficienteMadurez)  {
  let mensaje = "\nEl donante no cumple las condiciones para donar\n"; 

  if ((edadDonante >= EDAD_MINIMA &&  edadDonante <= EDAD_MAXIMA ||
      edadDonante >= EDAD_MINIMA_AUTORIZACION &&  edadDonante <= EDAD_MAXIMA_AUTORIZACION && 
      autorizacionEdad) && pesoDonante > PESO_MINIMO && suficienteMadurez){
    mensaje = "\nEl donante es elegible para donar\n";
  }

  return mensaje;
} 