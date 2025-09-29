/*
   Crear un programa para aplicar un descuento a la tarifa del
   transporte público según el perfil del usuario, como estudiantes o
   adultos mayores. Considerando que los domingos y festivos la tarifa
   es mayor.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Mayo 2025
   Licencia: GNU GPL v3
*/

const TARIFA_NORMAL = 2900;
const TARIFA_DOMINGO_FESTIVO = 3000;

const PORCENTAJE_DESCUENTO_ESTUDIANTE = 10.0;
const PORCENTAJE_DESCUENTO_ADULTO_MAYOR = 15.0;

function main() {
  let esDomingoFestivo = ingresarLogico('esDomingoFestivo');
  let esEstudiante = ingresarLogico('esEstudiante');
  let esAdultoMayor = ingresarLogico('esAdultoMayor');

  let tarifaDia = obtenerTarifaDia(esDomingoFestivo);

  let porcentajeDescuento = obtenerPorcentajeDescuento(
    esEstudiante, esAdultoMayor);

  let valorDescuento = calcularValorDescuento(
    tarifaDia, porcentajeDescuento);

  let valorTarifa = calcularValorTarifa(
    tarifaDia, valorDescuento);

  let reciboTarifa = generarReciboTarifa(
    tarifaDia, valorTarifa, porcentajeDescuento, valorDescuento);

  mostrarMensaje(reciboTarifa);
}

function obtenerTarifaDia(esDomingoFestivo) {
  let tarifaDia = TARIFA_NORMAL;

  if (esDomingoFestivo == true) {
    tarifaDia = TARIFA_DOMINGO_FESTIVO;
  }

  return tarifaDia;
}

function obtenerPorcentajeDescuento(esEstudiante, esAdultoMayor) {
  let porcentajeDescuento = 0.0;

  if (esEstudiante == true) {
    porcentajeDescuento = PORCENTAJE_DESCUENTO_ESTUDIANTE;
  }

  if (esAdultoMayor == true) {
    porcentajeDescuento = PORCENTAJE_DESCUENTO_ADULTO_MAYOR;
  }

  return porcentajeDescuento;
}

function calcularValorDescuento(tarifaDia, porcentajeDescuento) {
  return tarifaDia * porcentajeDescuento / 100.0;
}

function calcularValorTarifa(tarifaDia, valorDescuento) {
  return tarifaDia - valorDescuento;
}

function generarReciboTarifa(tarifaDia, valorTarifa, porcentajeDescuento, valorDescuento) {
  let mensaje = `\nLa tarifa es de $${tarifaDia}\n`;

  if (valorDescuento > 0) {
    mensaje = `${mensaje}` +
      `\nsu tarifa a pagar es de $${valorTarifa.toFixed(0)} por tener ` +
      `\nun descuento del ${porcentajeDescuento.toFixed(1)}\% ` +
      `equivalente  a $${valorDescuento.toFixed(0)}\n`;
  }
  
  return mensaje;
}     