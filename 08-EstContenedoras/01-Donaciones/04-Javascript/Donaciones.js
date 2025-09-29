/*
   Crear un programa para registrar y consultar de manera ágil
   todas las donaciones recibidas. Esta debería ofrecer el total 
   recaudado, listados alfabéticos y descendentes por valor, 
   identificar a los aportantes que superen un umbral específico y 
   señalar al mayor donante.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

const UMBRAL = 10000;

function main() {

  let donantes = ingresarColeccion(ingresarDonante);

  let donantesPorNombre     = ordenarColeccion(donantes, compararNombre,   false);
  let donantesPorOrdenacion = ordenarColeccion(donantes, compararDonacion, true);

  let mayoresDonantes = obtenerMayoresDonaciones(donantes, UMBRAL);
  let mayorDonante = obtenerMayorDonante(donantes);
  let sumaDonaciones = obtenerSumaDonaciones(donantes);

  let reporteDonaciones = generarReporteDonaciones(
    donantesPorNombre,
    donantesPorOrdenacion, 
    mayoresDonantes, 
    mayorDonante, 
    sumaDonaciones);

  mostrarMensaje(reporteDonaciones);
}

function ingresarColeccion(ingresarElemento) {
  const coleccion = [];
  let hayMas = true;

  while (hayMas) {
    coleccion.push(ingresarElemento());

    hayMas = confirm("Hay más datos (s:Aceptar / n:Cancelar): ");
  }

  return coleccion;
}

function ingresarDonante() {
  alert("Ingrese los datos de un donante:");

  let nombre = prompt("Ingrese el nombre:");
  let donacion = Number(prompt("Ingrese donación:")) || 0;

  return { nombre, donacion };
}

function ordenarColeccion(coleccion, comparador, descendente) {
  return [...coleccion].sort(
    (donante1, donante2) => descendente ? 
      comparador(donante2, donante1) :
      comparador(donante1, donante2)
  );
}

function compararNombre(d1, d2) {
  return d1.nombre.localeCompare(d2.nombre);
}

function compararDonacion(d1, d2) {
  return d1.donacion - d2.donacion;
}

function obtenerMayoresDonaciones(coleccion, limite) {
  return coleccion.filter(donante => donante["donacion"] > limite);
}

function obtenerMayorDonante(coleccion) {
  return coleccion.reduce((donante1, donante2) =>
    donante1 == null ||
      donante2["donacion"] > donante1["donacion"] ?
      donante2 : donante1, null);
}

function obtenerSumaDonaciones(coleccion) {
  return coleccion.reduce(
    (sumaDonaciones, donante) => sumaDonaciones + donante["donacion"], 0
  );
}

function generarReporteDonaciones(donantesPorNombre, 
                                  donantesPorDonacion, 
                                  mayoresDonantes, 
                                  mayorDonante, 
                                  sumaDonaciones) {
  listadoPorNombre = convertirColeccionCadena(
      "LISTADO EN ORDEN ALFABÉTICO", 
      donantesPorNombre, 
      obtenerDonante);

  listadoPorDonacion = convertirColeccionCadena(
      "LISTADO ORDENADO POR DONACIÓN", 
      donantesPorDonacion, 
      obtenerDonante);

  listadoMayores = convertirColeccionCadena(
      `LISTADO DONACIONES MAYORES A $${UMBRAL}`, 
      mayoresDonantes, 
      obtenerDonante);

  return `${listadoPorNombre}\n${listadoPorDonacion}\n${listadoMayores}\n`+
    `El mayor donante: ${mayorDonante.nombre}\n` +
    `Total de donantes $${sumaDonaciones}\n`;
}

function convertirColeccionCadena(titulo, lista, convertirElementoCadena) {
  let mensaje = `\n${titulo}\n`;

  for (const elemento of lista) {
    mensaje += "\t" + convertirElementoCadena(elemento);
  }

  return mensaje;
}

function obtenerDonante(donante) {
  return `$${String(donante.donacion).padStart(10)} \t ${donante.nombre}\n`;
}
