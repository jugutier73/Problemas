/*
   Crear un programa para contar cuantas veces se emplean símbolos 
   "x" o la "@" en un mensaje.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

const SIMBOLO_INCLUSIVO_1 = 'x';
const SIMBOLO_INCLUSIVO_2 = '@';

function main() {
	let texto = ingresarTexto('texto');

	let cantidadSimbolos = contarEmpleoSimbolos(texto);
	let reporteSimbolos = generarReporteInclusivo(cantidadSimbolos);

	mostrarMensaje(reporteSimbolos);
}

function contarEmpleoSimbolos(texto) {
  
  let cantidadSimbolosInclusivos = 0;

  for (caracter of texto) {
    if (caracter === SIMBOLO_INCLUSIVO_1 || 
        caracter === SIMBOLO_INCLUSIVO_2) {
      cantidadSimbolosInclusivos += 1;
    }
  }

  return cantidadSimbolosInclusivos;
}

function generarReporteInclusivo(cantidadSimbolosInclusivos) {
	return `\nSe emplearon ${cantidadSimbolosInclusivos} veces ` +
		   `los símbolos inclusivos \"${SIMBOLO_INCLUSIVO_1}\" y ` +
		   `\"${SIMBOLO_INCLUSIVO_2}\".\n`;
}
