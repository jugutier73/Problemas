/*
   Crear un programa para garantizar que los pacientes sean atendidos
   de manera equitativa y oportuna. El orden de llegada debe
   equilibrarse con la prioridad médica de cada caso

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

const NIVELES = 3;

const NIVEL_RIESGO      = 1;
const NIVEL_URGENCIA    = 2;
const NIVEL_PRIORITARIO = 3;

function main() {
  let citas = ingresarColeccion(ingresarCita);

  let citasPorNivel = ordenarColeccion(citas, 
    compararCampoNivel, false);

  let cantidadRiesgo      = contarSegunCriterio(citas, 
    tenerNivelRiesgo, NIVEL_RIESGO);

  let cantidadUrgencia    = contarSegunCriterio(citas, 
    tenerNivelRiesgo, NIVEL_URGENCIA);

  let cantidadPrioritario = contarSegunCriterio(citas, 
    tenerNivelRiesgo, NIVEL_PRIORITARIO);

  let reporteCitas = generarReporteCitas(citasPorNivel,
                                         cantidadRiesgo,
                                         cantidadUrgencia,
                                         cantidadPrioritario);

  mostrarMensaje(reporteCitas);
}

function ingresarCita() {
  alert("Ingrese los datos de la cita:");

  let nombre = prompt("Ingrese el nombre paciente:");
  let nivel = Number(prompt("NIVEL DE URGENCIA\n"+
		"  1: Riesgo\n"+
		"  2: Urgencia\n"+
		"  3: Prioritario\n"+
		"Ingrese tipo de persona: ")) || 0;

  return { nombre, nivel };
}

function compararCampoNivel(cita1, cita2) {
  return cita1.nivel - cita2.nivel;
}

function contarSegunCriterio(coleccion, aplicarCriterio, valorCriterio) {
  return coleccion.filter((elemento) => aplicarCriterio(elemento, valorCriterio) ).length;
}

function tenerNivelRiesgo(cita, nivelRiesgo) {
	return cita.nivel == nivelRiesgo
}

function generarReporteCitas(citasPorNivel,
                             cantidadRiesgo,
                             cantidadUrgencia,
                             cantidadPrioritario) {
  listadoPorNivel = convertirColeccionCadena(
      "LISTADO POR NIVEL DE URGENCIA", 
      citasPorNivel, 
      convertirCitaCadena);

  return `${listadoPorNivel}\n\n` +
    `Nivel 1 (Riesgo)     : ${cantidadRiesgo}\n` +
    `Nivel 2 (Urdencia)   : ${cantidadUrgencia}\n` + 
    `Nivel 3 (Prioritario): ${cantidadPrioritario}\n`;
}

function convertirCitaCadena(cita) {
  return `${cita.nivel} : ${cita.nombre}\n`;
}
