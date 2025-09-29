/*
  Programa que muestra el mensaje "Hola Mundo" en la pantalla
  Autor: Julián Esteban Gutiérrez Posada
  Fecha: Enero 2025
  Licencia: GNU GPL v3
*/

function main() {
  mostrarMensaje("Hola Mundo\n");
}

function mostrarMensaje(mensaje) {
  document.getElementById("salida").value = mensaje;
}