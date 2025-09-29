/*
   Crear un programa para contar cuantas palabras comienzan con
   la letra "p" en un mensaje.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

const LETRA_INICIO = 'p';

function main() {
	let texto = ingresarTexto('texto');

	let cantidadPalabras = contarPalabrasInician(texto, LETRA_INICIO);
	let reporteCiberacoso = generarReporteAcoso(cantidadPalabras);

	mostrarMensaje(reporteCiberacoso);
}

function contarPalabrasInician(texto, letraInicio) {
    let palabras = texto.toLowerCase().split(" ");

    let cantidadPalabrasInteres = 0;

    for (let palabra of palabras) {
        if (palabra.startsWith(letraInicio)) {
            cantidadPalabrasInteres++;
        }
    }
    return cantidadPalabrasInteres;
}

function generarReporteAcoso(cantidadPalabrasInteres) {
	return `\nHay ${cantidadPalabrasInteres} palabras ` +
		     `que inician con la letra \"${LETRA_INICIO}\".\n`;
}