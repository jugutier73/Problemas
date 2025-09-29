/*
   Crear un programa para contar cuántas letras mayúsculas,
   cuántas minúsculas y cuántos dígitos contiene una contraseña

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

function main() {
	let contrasenia = ingresarTexto('contrasenia'); 

	let cantidadMayusculas = contarMayusculas(contrasenia);
	let cantidadMinusculas = contarMinusculas(contrasenia);
	let cantidadDigitos = contarDigitos(contrasenia);

	let reporteEstres = generarReporteContrasenia(contrasenia,
		cantidadMayusculas,cantidadMinusculas, cantidadDigitos);

	mostrarMensaje(reporteEstres);
}

function contarMayusculas (texto) {
	return contarCaracteres(texto, esMayuscula);
} 

function contarMinusculas  (texto) {
	return contarCaracteres(texto, esMinuscula);
}

function contarDigitos     (texto) {
	return contarCaracteres(texto, esDigito);
}

function esMayuscula (caracter) { return /\p{Lu}/u.test(caracter); }
function esMinuscula (caracter) { return /\p{Ll}/u.test(caracter); }
function esDigito    (caracter) { return /\p{Nd}/u.test(caracter); }

function contarCaracteres(texto, funcion) {
    let cantidadCaracteres = 0;

	for (let i = 0 ; i < texto.length ; i++) {
		if (funcion(texto.charAt(i))) {
			cantidadCaracteres += 1
		}
	}
	
	return cantidadCaracteres;
}

function generarReporteContrasenia(contrasenia,
	cantidadMayusculas, cantidadMinusculas, cantidadDigitos) {
	return `\nEn la constraseña \"${contrasenia}\"  hay:\n ` +
		   `${cantidadMayusculas} mayúscula(s), ` +
		   `${cantidadMinusculas} minúscula(s) y ` +
		   `${cantidadDigitos} dígito(s).\n`;
}