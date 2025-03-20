/*
    Programa para calcular el índice de masa corporal (IMC), usando la 
    altura y el peso.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
*/

function main() {
  let peso = ingresaEntero('peso');
  let altura = ingresaReal('altura');
  let imc = calcularIMC(peso, altura);
	let informeIMC = generarInformeIMC(peso, altura, imc);
	mostrarMensaje(informeIMC);
}

function ingresaReal(componente) {
  return parseFloat(ingresarTexto(componente));
}

function calcularIMC(peso, altura) {
  return peso/Math.pow(altura,2.0);
}

function generarInformeIMC(peso, altura, imc) {
  return `\nCon su peso de ${peso} kg y su altura de ${altura.toFixed(1)} metros` +
         `\nsu \xEDndice de masa corporal (IMC) es de ${imc.toFixed(1)}`;
}