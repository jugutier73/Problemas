/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Septiembre 2025
   Licencia: GNU GPL v3
*/

/**
 * Muestra un mensaje (cadena) en la salida estandar.
 * 
 * @param {string} mensaje - Mensaje que se desea mostrar.
 */
function mostrarMensaje(mensaje) {
  document.getElementById("salida").value = mensaje;
}


/**
 * Muestra un mensaje (cadena) en el error estandar.
 * 
 * @param {string} mensaje Mensaje que se desea mostrar como error.
 */
function mostrarError(mensaje) {
  alert(mensaje);
}


/**
 * Devuelve el texto ingresado por el usuario como respuesta a una pregunta.
 * 
 * @param {string} componente Nombre del componente donde el usuario ingresa la respuesta
 * @return {string} Texto ingresado por el usuario
 */
function ingresarTexto(componente) {
  return document.getElementById(componente).value;
}


/**
 * Devuelve el entero ingresado por el usuario como respuesta a una pregunta.
 * 
 * @param {string} componente Nombre del componente donde el usuario ingresa la respuesta
 * @return {int} Valor ingresado por el usuario o cero si es un valor inválido.
 */
function ingresarEntero(componente) {
  let entero = Number(ingresarTexto(componente));
  
  if (!Number.isInteger(entero)) {
    mostrarError("El valor ingresado es inválido, intente nuevamente.");
    seleccionarError(componente);
    exit(1);
  }

  return entero;
}


/**
 * Selecciona el contenido de un componente html que tiene una entrada inválida.
 * 
 * @param {string} componente Nombre del componente donde el usuario ingresa la respuesta inválida.
 */
function seleccionarError(componente) {
  document.getElementById(componente).select();
}


/**
 * Devuelve el real ingresado por el usuario como respuesta a una pregunta.
 * 
 * @param {string} componente Nombre del componente donde el usuario ingresa la respuesta
 * @return {float} Valor ingresado por el usuario o cero si es un valor inválido.
 */
function ingresarReal(componente) {
  let real = Number(ingresarTexto(componente));

  if (isNaN(real)) {
    mostrarError("El valor ingresado es inválido, intente nuevamente.");
    seleccionarError(componente);
    exit(1);
  }
  
  return real;
}


/**
 * Devuelve el booleano ingresado por el usuario como respuesta a una pregunta.
 * 
 * @param {string} componente Nombre del componente donde el usuario ingresa la respuesta
 * @return {boolean} Valor ingresado por el usuario.
 */
function ingresarLogico(componente) {
  return document.getElementById(componente).checked;
}


/**
 * Devuelve el entero ingresado por el usuario como respuesta a una pregunta con múltiples opciones del 1 a un máximo especificado.
 * 
 * @param {string} componente Nombre del componente donde el usuario ingresa la respuesta
 * @param {int} maximaOpcion Entero que indica la cantidad de opciones.
 * @return {int} Valor ingresado por el usuario.
 */
function ingresarOpcion(componente) {
	return document.getElementById(componente).selectedIndex + 1;
}