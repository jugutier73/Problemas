/*
   Crear un programa para a clasificación de residuos según su
   material. Residuos aprovechable (limpios y secos), usar bolsa
   blanca; residuos orgánicos, usar bolsa verde; lo anterior si
   existe ruta de recolección selectiva; en otro caso, usar bolsa
   de color negro.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

const APROVECHABLE = 1;
const OTRO         = 3;

function main() {
	let tipoResiduo = ingresarOpcion('tipoResiduo');
	let hayRecoleccionSelectiva = ingresarLogico('hayRecoleccionSelectiva');

	let colorBolsa = recomendarColorBolsa(
		tipoResiduo, hayRecoleccionSelectiva);
		
	let reporteBolsaRecoleccion = generarReporteBolsa(colorBolsa);

	mostrarMensaje(reporteBolsaRecoleccion);
}

function ingresarOpcion(componente) {
	return document.getElementById(componente).selectedIndex + 1;
}

function recomendarColorBolsa(tipoResiduo,hayRecoleccionSelectiva) {
	let colorBolsa;

	if (tipoResiduo == OTRO || !hayRecoleccionSelectiva) {
		colorBolsa = "NEGRA";
	} else {
		if (tipoResiduo == APROVECHABLE) {
			colorBolsa = "BLANCA";
		} else {
			colorBolsa = "VERDE";
		}
	}

	return colorBolsa;
}

function generarReporteBolsa(colorBolsa){
	return `\nLa bolsa recomendada es de color ${colorBolsa}.\n`;
}