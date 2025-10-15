/*
   Crear un programa para organizar la entrada de los beneficiarios
   y garantice la atención en orden de llegada. Se debe registrar
   datos clave (nombre, edad y presencia de necesidades especiales).
   Además de informar la cantidad de personas con necesidades y el
   promedio de edades.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

function main() {

  let reservasComedor = ingresarColeccion(ingresarReserva);

  let cantidadConNecesidades = contarSegunCriterio(reservasComedor, tenerNecesidadEspecial, true);

  let promedioEdades = calcularPromedioEdades(reservasComedor);

  let reporteReservas = generarReporteReservasComedor(
		reservasComedor,
		cantidadConNecesidades,
		promedioEdades);

  mostrarMensaje(reporteReservas);
}

function ingresarReserva() {
  alert("Ingrese los datos de la reserva:");

  let nombre = prompt("Ingrese el nombre de la persona  :");
  let edad = Number(prompt("Ingrese la edad de la persona    :")) || 0;
  let necesidadEspecial = confirm("Tiene necesidades especiales (s:Aceptar / n:Cancelar): ");

  return { nombre, edad, necesidadEspecial };
}

function tenerNecesidadEspecial(reserva, necesidadEspecialInteres) {
  return reserva.necesidadEspecial == necesidadEspecialInteres;
}

function calcularPromedioEdades(reservasComedor) {
	let totalReservas = reservasComedor.length;
	let promedioEdades = 0.0;

	if ( totalReservas > 0) {
		promedioEdades = reservasComedor
			.reduce((sumaEdades, reserva) => sumaEdades + reserva.edad, 0) / totalReservas;
	}
  
  return promedioEdades;
}


function generarReporteReservasComedor(reservasComedor,
									   cantidadConNecesidades,
									   promedioEdades) {
  listadoReservas = convertirColeccionCadena(
      "LISTADO DE RESERVAS", 
      reservasComedor, 
      convertirReservaCadena);

  return `${listadoReservas}\n\n` +
    `Cantidad con necesidades : ${cantidadConNecesidades}\n` +
    `Promedio de edades       : ${promedioEdades.toFixed(1)}\n`;
}

function convertirReservaCadena(reserva) {
	let mensaje = `${reserva.nombre}, ${reserva.edad} años`;

	if (reserva.necesidadEspecial) {
		mensaje += ", con necesidad especial";
	}

  	return mensaje + "\n";
}
