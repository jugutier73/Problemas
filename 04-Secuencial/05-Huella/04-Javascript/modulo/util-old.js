/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Marzo 2025
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
 * Devuelve el texto ingresado por el usuario como respuesta a una pregunta
 * 
 * @param {string} componente Nombre del componente donde el usuario 
 *                            ingresa la respuesta
 * @return {string} Texto ingresado por el usuario
 */
function ingresarTexto(componente) {
  return document.getElementById(componente).value;
}

/**
 * Devuelve el entero ingresado por el usuario como respuesta a una 
 * pregunta.
 * 
 * @param {string} componente Nombre del componente donde el usuario 
 *                            ingresa la respuesta
 * @return {int} Valor ingresado por el usuario o cero si es un valor 
 *               inválido.
 */
function ingresarEntero(componente) {
  return parseInt(ingresarTexto(componente));
}

/**
 * Devuelve el real ingresado por el usuario como respuesta a una pregunta.
 * 
 * @param {string} componente Nombre del componente donde el usuario 
 *                            ingresa la respuesta
 * @return {float} Valor ingresado por el usuario o cero si es un valor
 *                 inválido.
 */
function ingresarReal(componente) {
  return parseFloat(ingresarTexto(componente));
}