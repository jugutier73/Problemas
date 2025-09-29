/*
   Programa que registra los residuos recolectados durante una
   jornada, especificando su código, tipo y condición de
   reparabilidad. El sistema genera un informe con el
   listado de los elementos recolectados, ordenados por tipo y
   reparabilidad, además de presentar un conteo total por cada
   tipo de residuo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

const ANCHO = 15;

const BATERIA = 1;
const TELEFONO = 2;
const COMPUTADOR = 3;
const CARGADOR = 4;
const OTRO_DISPOSITIVO = 5;

const ETIQUETAS = {
	1: "Batería",
	2: "Teléfono",
	3: "Computador",
	4: "Cargador",
	5: "Otro"
};

function main() {
	let residuos = ingresarColeccion(ingresarResiduo);

	let residuosClasificados = ordenarColeccion(residuos, 
		compararCampoTipo, false);

	let cantidadBaterias = contarSegunCriterio(residuos, 
		tenerTipo, BATERIA);

	let cantidadTelefonos = contarSegunCriterio(residuos, 
		tenerTipo, TELEFONO);

	let cantidadComputadores = contarSegunCriterio(residuos, 
		tenerTipo, COMPUTADOR);

	let cantidadCargadores = contarSegunCriterio(residuos, 
		tenerTipo, CARGADOR);

	let cantidadOtros = contarSegunCriterio(residuos, 
		tenerTipo, OTRO_DISPOSITIVO);

	let reporteRecoleccion = generarReporteRecoleccion(
		residuosClasificados,
		cantidadBaterias,
		cantidadTelefonos,
		cantidadComputadores,
		cantidadCargadores,
		cantidadOtros);

	mostrarMensaje(reporteRecoleccion);
}

function ingresarResiduo() {
	alert("Ingrese los datos del residuo electrónico:");

	let codigo = prompt("Ingrese el código del residuo: ");
	let tipo = Number(prompt("TIPOS DE RESIDUOS\n" +
		"    1: Batería\n" +
		"    2: Teléfono\n" +
		"    3: Computador\n" +
		"    4: Cargador\n" +
		"    5: Otro dispositivo)\n" +
		"Ingrese tipo de residuo      : ")) || 0;
	let reparable = confirm("Puede ser reparado (s:Aceptar / n:Cancelar): ");

	return { codigo, tipo, reparable };
}

function compararCampoTipo(residuo1, residuo2) {
	let comparacionTipo = residuo1.tipo - residuo2.tipo;

	if (comparacionTipo == 0) {
		comparacionTipo = residuo2.reparable - residuo1.reparable;
	}

	return comparacionTipo;
}


function tenerTipo(residuo, tipoInteres) {
	return residuo.tipo == tipoInteres;
}


function generarReporteRecoleccion(residuosClasificados,
	cantidadBaterias,
	cantidadTelefonos,
	cantidadComputadores,
	cantidadCargadores,
	cantidadOtros) {
	listadoResiduosClasificados = convertirColeccionCadena(
		"LISTADO DE RESIDUOS CLASIFICADOS",
		residuosClasificados,
		convertirResiduoCadena);

	return `${listadoResiduosClasificados}\n\n` +
		`Cantidad de Baterías     : ${cantidadBaterias}\n` +
		`Cantidad de Teléfonos    : ${cantidadTelefonos}\n` +
		`Cantidad de Computadores : ${cantidadComputadores}\n` +
		`Cantidad de Cargadores   : ${cantidadCargadores}\n` +
		`Cantidad de Otros        : ${cantidadOtros}\n`;
}

function convertirResiduoCadena(residuo) {
	etiqueta = ETIQUETAS[residuo.tipo];
	reparable = residuo.reparable ? "(REPARABLE)" : "";

	return `${residuo.codigo.padEnd(ANCHO)} ${etiqueta.padEnd(ANCHO)} ${reparable}\n`;
}
