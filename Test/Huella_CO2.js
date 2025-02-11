const FACTOR_CARRO = 121.0;
const FACTOR_MOTO = 53.0;
const FACTOR_BICICLETA = 3.0;

const FACTOR_BUS = 49.0;

const FACTOR_CARNE_ROJA = 7.19;
const FACTOR_PESCADO = 3.91;
const FACTOR_VEGETARIANA = 3.81;

function main() {
  mostrarMensaje("KILÓMETROS RECORRIDOS EN VEHÍCULOS PARTICULARES");
  const kmCarro = ingresarReal("kmCarro");
  const kmMoto = ingresarReal("kmMoto");
  const kmBicicleta = ingresarReal("kmBicicleta");

  mostrarMensaje("\nKILÓMETROS RECORRIDOS EN TRANSPORTE URBANO");
  const trayectosBus = ingresarReal("trayectosBus");

  mostrarMensaje("\nCANTIDAD COMIDAS CON");
  const comidasCarneRoja = ingresarEntero("comidasCarneRoja");
  const comidasPescado = ingresarEntero("comidasPescado");
  const comidasVegetarianas = ingresarEntero("comidasVegetarianas");

  const kgCo2Recorridos = calcularKgCo2Recorridos(kmCarro,
    kmMoto, kmBicicleta);
  const kgCo2Trayectos = calcularKgCo2Trayectos(trayectosBus);
  const kgCo2Comidas = calcularKgCo2Comidas(comidasCarneRoja,
    comidasPescado, comidasVegetarianas);
  const kgCo2Total = calcularKgCo2Total(kgCo2Recorridos,
    kgCo2Trayectos, kgCo2Comidas);

  const mensajeSalida = generarMensajeSalida(kgCo2Recorridos,
    kgCo2Trayectos, kgCo2Comidas, kgCo2Total);

  mostrarMensaje(mensajeSalida);
}

function mostrarMensaje(mensaje) {
  document.getElementById("salida").value = mensaje;
}

function ingresarEntero(componente) {
  const entero = parseInt(document.getElementById(componente).value, 10);
  return isNaN(entero) ? 0 : entero;
}

function ingresarReal(componente) {
  const real = parseFloat(document.getElementById(componente).value);
  return isNaN(real) ? 0.0 : real;
}

function calcularKgCo2Recorridos(kmCarro, kmMoto, kmBicicleta) {
  return (
    kmCarro * FACTOR_CARRO + 
    kmMoto * FACTOR_MOTO + 
    kmBicicleta * FACTOR_BICICLETA
  );
}

function calcularKgCo2Trayectos(trayectosBus) {
  return trayectosBus * FACTOR_BUS;
}

function calcularKgCo2Comidas(
  comidasCarneRoja, 
  comidasPescado, 
  comidasVegetarianas) {
  return ( 
    comidasCarneRoja * FACTOR_CARNE_ROJA + 
    comidasPescado * FACTOR_PESCADO + 
    comidasVegetarianas * FACTOR_VEGETARIANA
  );
}

function calcularKgCo2Total(
  kgCo2Recorridos, 
  kgCo2Trayectos, 
  kgCo2Comidas) {
  return kgCo2Recorridos + kgCo2Trayectos + kgCo2Comidas;
}

function generarMensajeSalida(
  kgCo2Recorridos, 
  kgCo2Trayectos, 
  kgCo2Comidas, 
  kgCo2Total) {
  return `\nSu huella de carbono en kg de CO2 por:
- RECORRIDOS ${kgCo2Recorridos.toFixed(2)} kg
- TRAYECTOS ${kgCo2Trayectos.toFixed(2)} kg
- COMIDAS ${kgCo2Comidas.toFixed(2)} kg

Para un TOTAL de ${kgCo2Total.toFixed(2)} kg de emisiones de CO2.`;
}

