/*
   Crear un programa para contar cuántas letras mayúsculas,
   cuántas minúsculas y cuántos dígitos contiene una contraseña

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

void main() {
	var contrasenia = Util.ingresarTexto("Ingrese la contraseña analizar: ");

	var cantidadMayusculas = contarMayusculas(contrasenia);
	var cantidadMinusculas = contarMinusculas(contrasenia);
	var cantidadDigitos = contarDigitos(contrasenia);

	var reporteContrasenia = generarReporteContrasenia(contrasenia,
		cantidadMayusculas, cantidadMinusculas,	cantidadDigitos);

	Util.mostrarMensaje(reporteContrasenia);
}

int contarMayusculas(String texto)  { 
	return contarCaracteres(texto, esMayuscula); 
}

int contarMinusculas(String texto)  { 
	return contarCaracteres(texto, esMinuscula); 
}

int contarDigitos(String texto)     { 
	return contarCaracteres(texto, esDigito); 
}

IntPredicate esMayuscula = Character::isUpperCase;
IntPredicate esMinuscula = Character::isLowerCase;
IntPredicate esDigito    = Character::isDigit;

int contarCaracteres(String texto, IntPredicate funcion) {
	var cantidadCaracteres = 0;

	for (var caracter : texto.toCharArray())  {
		if (funcion.test(caracter)) {
			cantidadCaracteres++;
		}
	}

	return cantidadCaracteres;
}

String generarReporteContrasenia(String contrasenia, 
int cantidadMayusculas, int cantidadMinusculas, int cantidadDigitos) {
	return String.format("\nEn la constraseña \"%s\" hay:\n"+
		"%d mayúscula(s), "+
		"%d minúscula(s) y "+
		"%d dígito(s)\n",
		contrasenia, cantidadMayusculas, cantidadMinusculas, cantidadDigitos);
}
