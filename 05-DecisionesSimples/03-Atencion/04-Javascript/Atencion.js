/*
   Programa para establecer la prioridad de atención
   en los centros de salud para identificar a los pacientes con
   mayor riesgo o vulnerabilidad
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

const EDAD_RECIEN_NACIDO = 1;
const EDAD_ADULTO_MAYOR = 60;

function main() {
    let edadPaciente = ingresarEntero('edadPaciente');
    let enfermedadCronica = ingresarLogico('enfermedadCronica');
    let estadoInmunosupresion = ingresarLogico('estadoInmunosupresion');

    let reporteAtencion = generarReporteAtencion(edadPaciente, enfermedadCronica, estadoInmunosupresion);  

    mostrarMensaje(reporteAtencion);
}

function generarReporteAtencion(edadPaciente, enfermedadCronica, estadoInmunosupresion) {
  let mensaje = `\nEl paciente es de atención general\n`;

  if (edadPaciente < EDAD_RECIEN_NACIDO || 
      edadPaciente > EDAD_ADULTO_MAYOR  || 
      enfermedadCronica || estadoInmunosupresion){
    mensaje = `\nEl paciente es de atención prioritaria\n`;          
  }
  
  return mensaje;
}         