/*
    Programa que permita tener un control del consumo de la energía
    eléctrica con relación al consumo del mes anterior, para que 
    así el usuario pueda adoptar hábitos de consumo más conscientes 
    y responsables.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
*/

function main() {
  let consumoActual = ingresarEntero('consumoActual');
  let consumoAnterior = ingresarEntero('consumoAnterior');

  let relacionConsumo = calcularRelacionConsumo(
    consumoActual, consumoAnterior);
  let reporteRelacion = generarReporteRelacion(
    consumoActual, consumoAnterior, relacionConsumo);
    
  mostrarMensaje(reporteRelacion);
}

function calcularRelacionConsumo(consumoActual, consumoAnterior) {
  return consumoActual / consumoAnterior * 100.0;
}

function generarReporteRelacion(
  consumoActual, consumoAnterior, relacionConsumo) {
  return `\nEl consumo actual de ${consumoActual} kilovatios representa` +
         `\nun ${relacionConsumo.toFixed(1)}% con relación al consumo ` +
         `del mes\nanterior de ${consumoAnterior} kilovatios.\n`;
}