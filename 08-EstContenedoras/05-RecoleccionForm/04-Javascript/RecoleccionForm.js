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

// Estado para almacenar la lista de residuos
const residuos = [];

// Evento del botón de "Agregar"
function agregarResiduo() {
    let residuo = ingresarResiduo();

    residuos.push(residuo);
}

function ingresarResiduo() {
    let codigo    = ingresarTexto ("codigo");
    let tipo      = ingresarTexto ("tipo");
    let reparable = ingresarLogico("reparable");

    return { codigo, tipo, reparable };
}

function actualizarTabla() {
    document.getElementById("codigo").value = "";
    document.getElementById("salida").value = "";

    let tabla = document.getElementById("cuerpoTabla");
    tabla.innerHTML = "";

    residuos.forEach(
        (residuo, indice) => {
            insertarFila(tabla, residuo, indice);
        }
    );
}

function insertarFila(tabla, residuo, indice) {
    let fila = document.createElement("tr");

    fila.innerHTML = `<td>${residuo.codigo}</td>`
        + `<td>${ETIQUETAS[residuo.tipo] ?? "Desconocido"}</td>`
        + `<td>${residuo.reparable ? "(REPARABLE)" : ""}</td>`
        + `<td>`
        + `  <button `
        + `      onclick="borrarResiduo(${indice}); actualizarTabla();">`
        + `      Eliminar`
        + `  </button>`
        + `</td>`;

    tabla.appendChild(fila);
}

// Evento del botón de "Eliminar"
function borrarResiduo(indice) {
    residuos.splice(indice, 1);
}

// Evento del botón de "Limpiar"
function limpiarFormulario() {
    residuos.length = 0;
}

// Evento del botón de "Generar reporte"
function generarReporte() {
    let residuosClasificados = ordenarColeccion(residuos, 
        compararCampoTipo, false);
    
    let cantidadBaterias = contarSegunCriterio(residuos, 
        tenerTipo,BATERIA);
        
    let cantidadTelefonos = contarSegunCriterio(residuos, 
        tenerTipo, TELEFONO);
        
    let cantidadComputadores = contarSegunCriterio(residuos, 
        tenerTipo, COMPUTADOR);
        
    let cantidadCargadores = contarSegunCriterio(residuos, 
        tenerTipo, CARGADOR);
            
    let cantidadOtros = contarSegunCriterio(residuos, 
        tenerTipo, OTRO_DISPOSITIVO);
        
    let listadoResiduosClasificados = convertirColeccionCadena(
        "LISTADO DE RESIDUOS CLASIFICADOS",
        residuosClasificados,
        convertirResiduoCadena);

    let listado = `${listadoResiduosClasificados}\n\n`
        + `Cantidad de Baterías     : ${cantidadBaterias}\n`
        + `Cantidad de Teléfonos    : ${cantidadTelefonos}\n`
        + `Cantidad de Computadores : ${cantidadComputadores}\n`
        + `Cantidad de Cargadores   : ${cantidadCargadores}\n`
        + `Cantidad de Otros        : ${cantidadOtros}\n`;

     mostrarMensaje(listado);
}

function tenerTipo(residuo, tipoInteres) {
	return residuo.tipo == tipoInteres;
}

function compararCampoTipo(residuo1, residuo2) {
    let comparacionTipo = residuo1.tipo - residuo2.tipo;

    if (comparacionTipo == 0) {
        comparacionTipo = residuo2.reparable - residuo1.reparable;
    }

    return comparacionTipo;
}

function convertirResiduoCadena(residuo) {
    etiqueta = ETIQUETAS[residuo.tipo];
    reparable = residuo.reparable ? "(REPARABLE)" : "";

    return `${residuo.codigo.padEnd(ANCHO)} ${etiqueta.padEnd(ANCHO)} ${reparable}\n`;
}