/*
   Crear un programa para determinar si hay problemas con el uso
   de los espacios en un frase que el usuario ingrese

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

function main() {
	let frase = ingresarTexto('frase');

	let fraseCorregida = corregirUsoEspacios(frase);

	let reporte = generarReporteAlfabetizacion(frase, fraseCorregida);

	mostrarMensaje(reporte);
}

function corregirUsoEspacios(frase) {
	let inicioFrase = false;
	let espacioPendiente = false;
	let fraseCorregida = "";

	for (var simbolo of frase) {
		if (simbolo == ' '){
			if (inicioFrase) {
				espacioPendiente = true;
			}
		} else {
			if (espacioPendiente) {
				fraseCorregida += " ";
				espacioPendiente = false;
			}
			fraseCorregida += simbolo;
			inicioFrase = true;
		}
	}

	return fraseCorregida;
}

function generarReporteAlfabetizacion(frase, fraseCorregida) {
	let reporteAlfabetizacion = `\nLa frase "${frase}" hace un uso `;

	if (frase == fraseCorregida) {
		reporteAlfabetizacion += `CORRECTO de espacios.`;
	} else {
		reporteAlfabetizacion += `INCORRECTO de espacios,\n`+
			`lo correcto es "${fraseCorregida}"`;
	}

	return reporteAlfabetizacion + "\n";
}

