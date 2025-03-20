/*
    Programa para calcular el índice de masa corporal (IMC), usando la 
    altura y el peso.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
*/

void main() {
  var peso = ingresarEntero("Peso  (kg): ");
	var altura = ingresarReal("Altura (m): ");
	var imc = calcularIMC(peso, altura);
	var informeIMC = generarInformeIMC(peso, altura, imc);
	mostrarMensaje(informeIMC);

	System.exit(0);
}

double ingresarReal(String pregunta) {
  return Double.parseDouble(ingresarTexto(pregunta));
}

double calcularIMC(int peso, double altura) {
  return (double)peso / Math.pow(altura, 2.0);
}

String generarInformeIMC(int peso, double altura, double imc) {
  return String.format(
    "\nCon su peso de %d kg y su altura de %.1f metros"+
    "\nsu índice de masa corporal (IMC) es de %.1f", 
    peso, altura, imc);
}

main();