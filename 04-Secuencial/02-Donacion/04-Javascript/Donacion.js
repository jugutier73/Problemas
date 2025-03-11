/*
   Programa para la iniciativa Amigo Social que permita generar
   el documento de recibido que sirva como soporte contable de
   la donación.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
*/

function main() {
  let nombreCasa = ingresarTexto("nombreAdultoMayor");
  let cantidadRecolectada = ingresarTexto("cantidadRecolectada");
  let reciboDonacion = generarRecibo(nombreCasa, cantidadRecolectada);
  mostrarMensaje(reciboDonacion);
}

function ingresaEntero(componente) {
  return parseInt(ingresarTexto(componente));
}

function generarRecibo(nombreCasa, cantidadRecolectada) {
  return `\nLa iniciativa Amigo Social tiene el gusto de entregar` +
         `\nuna donación de $${cantidadRecolectada} pesos colombianos` +
         `\na la casa del adulto mayor ${nombreCasa}.` +
         `\n\n_______________________`+
			   `\nFirma representante legal`
}