/*
   Crear un programa para registrar el ingreso y salida de los
   usuarios de un laboratorio para asegurar que solo personas
   autorizadas -estudiantes, docentes o personal técnico- accedan
   al espacio. El programa debe permitir identificar quiénes estaban 
   presentes y la cantidad por cada tipo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

const ANCHO = 20;

const ESTUDIANTE = 1;
const DOCENTE = 2;
const TECNICO = 3;

const ETIQUETAS = {
	1: "Estudiante",
	2: "Docente",
	3: "Técnico"
};

function main() {

	let ingresos = ingresarColeccion(ingresarIngreso);
	let salidas = ingresarColeccion(ingresarSalidas);

	let personasLaboratorio = obtenerPersonasLaboratorio(ingresos, salidas);

	let cantidadEstudiantes = contarSegunCriterio(personasLaboratorio, 
		tenerTipo, ESTUDIANTE);

	let cantidadDocentes = contarSegunCriterio(personasLaboratorio,
		 tenerTipo, DOCENTE);

	let cantidadTecnicos = contarSegunCriterio(personasLaboratorio, 
		tenerTipo, TECNICO);

	let reporteIngresos = generarReporteIngresos(
		ingresos,
		salidas,
		personasLaboratorio,
		cantidadEstudiantes,
		cantidadDocentes,
		cantidadTecnicos);

	mostrarMensaje(reporteIngresos);
}

function ingresarIngreso() {
	alert("Ingrese los datos del ingreso:");

	let documento = prompt("Ingrese el documento   : ");
	let nombre = prompt("Ingrese el nombre      : ");
	let tipo = Number(prompt("TIPO DE PERSONAS AUTORIZADAS\n" +
		"    1: Estudiante\n" +
		"    2: Docente\n" +
		"    3: Técnico\n" +
		"Ingrese tipo de persona      : ")) || 0;

	return { documento, nombre, tipo };
}

function ingresarSalidas() {
	alert("Ingrese los datos de la salida:");

	let documento = prompt("Ingrese el documento   : ");

	return documento;
}

function obtenerPersonasLaboratorio(ingresos, salidas) {
  let salidasSet = new Set(salidas);

  return ingresos.filter(ing => !salidasSet.has(ing.documento));
}

function tenerTipo(ingresos, tipo) {
	return ingresos.tipo == tipo;
}

function generarReporteIngresos(ingresos,
	salidas,
	personasLaboratorio,
	cantidadEstudiantes,
	cantidadDocentes,
	cantidadTecnicos) {
	listadoIngresos = convertirColeccionCadena(
		"LISTADO DE INGRESOS",
		ingresos,
		convertirIngresoCadena);

	listadoSalidas = convertirColeccionCadena(
		"LISTADO DE SALIDAS",
		salidas,
		convertirSalidaCadena);

	listadoIngresos = convertirColeccionCadena(
		"LISTADO DE PERSONAS EN EL LABORATORIO",
		personasLaboratorio,
		convertirIngresoCadena);	

	return `${listadoIngresos}\n\n` +
		`${listadoSalidas}\n\n` +
		`${listadoIngresos}\n\n` +
		`Cantidad de Estudiantes : ${cantidadEstudiantes}\n` +
		`Cantidad de Docentes    : ${cantidadDocentes}\n` +
		`Cantidad de Tecnicos    : ${cantidadTecnicos}\n`;
}

function convertirIngresoCadena(ingreso) {
	etiqueta = ETIQUETAS[ingreso.tipo];

	return `${ingreso.documento.padEnd(ANCHO)} `+
	       ` ${ingreso.nombre.padEnd(ANCHO)} ${etiqueta}\n`;
}

function convertirSalidaCadena(salida) {
	return salida;
}