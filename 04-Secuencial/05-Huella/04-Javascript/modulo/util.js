/*
   Módulo de utilidades (funciones de apoyo)
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Febrero 2025
   Licencia: GNU GPL v3
*/

function mostrarMensaje(mensaje) {
  document.getElementById("salida").value = mensaje;
}

function ingresarTexto(componente) {
  return document.getElementById(componente).value;
}

function ingresaEntero(componente) {
  return parseInt(ingresarTexto(componente));
}

function ingresaReal(componente) {
  return parseFloat(ingresarTexto(componente));
}