/*
    Crear un programa para programa para la verificación de medidas 
    de bioseguridad para el ingreso a un evento público.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Mayo 2025
    Licencia: GNU GPL v3
*/

function main() {
  let tieneVacunas = ingresarLogico('tieneVacunas');
  let restadosNegativos = ingresarLogico('restadosNegativos');
  let tieneSintomas = ingresarLogico('tieneSintomas');

  let reporteIngreso = generarReporteIngreso(
    tieneVacunas, restadosNegativos, tieneSintomas);  

  mostrarMensaje(reporteIngreso);
}

function ingresarLogico(componente) {
  return document.getElementById(componente).checked;
}

function generarReporteIngreso(tieneVacunas, restadosNegativos, tieneSintomas) {
 let mensaje = `\nLa persona no puede ingresar al evento\n`;          

 if (tieneVacunas && restadosNegativos && !tieneSintomas){
   mensaje = `\nLa persona puede ingresar al evento\n`;          
  }
  
 return mensaje;
}     

