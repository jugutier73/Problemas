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