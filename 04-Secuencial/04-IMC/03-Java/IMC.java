/*
    Programa para calcular el índice de masa corporal (IMC), usando la 
    altura y el peso.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
*/
import modulo.Util;

void main() {
  var peso = Util.ingresarEntero("Peso  (kg): ");
	var altura = ingresarReal("Altura (m): ");

	var imc = calcularIMC(peso, altura);
	var informeIMC = generarInformeIMC(peso, altura, imc);
  
	Util.mostrarMensaje(informeIMC);
}

double ingresarReal(String pregunta) {
  return Double.parseDouble(Util.ingresarTexto(pregunta));
}

double calcularIMC(int peso, double altura) {
  return (double)peso / Math.pow(altura, 2.0);
}

String generarInformeIMC(int peso, double altura, double imc) {
  return String.format(
    "\nCon su peso de %d kg y su altura de %.1f metros"+
    "\nsu índice de masa corporal (IMC) es de %.1f\n", 
    peso, altura, imc);
}
