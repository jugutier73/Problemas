/*
   Programa para la empresa de seguridad Aros S.A.
   que permita resaltar el nombre del empleado destacado
   del mes ofreciendo unas felicitaciones públicas.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
   Licencia: GNU GPL v3
*/

function main() {
  let nombreEmpleado = ingresarTexto("nombreEmpleado");
  let nombreMes = ingresarTexto("nombreMes");

  let felicitaciones = generarFelicitaciones(nombreEmpleado, nombreMes);
  
  mostrarMensaje(felicitaciones);
}

function ingresarTexto(componente) {
  return document.getElementById(componente).value;
}

function generarFelicitaciones(nombreEmpleado, nombreMes) {
  return `\nLa empresa de seguridad Aros S.A. quiere felicitar` +
         `\npúblicamente a ${nombreEmpleado} como nuestro` +
         `\nempleado destacado del mes de ${nombreMes},` +
         `\nmuchas gracias por su excelencia.\n`;
}