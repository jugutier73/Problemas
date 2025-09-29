/*
   Crear un programa para calcular, de forma clara y transparente,
   el valor de un trayecto en taxi considerando variables como
   el horario diurno/nocturno, los recargos de domingos y festivos,
   y el destino (ya sea dentro del perímetro urbano o fuera de él)
   para determinar si la tarifa cobrada es justa o no.
   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Julio 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

final int SIN_COSTO      = 0;
final int URBANO         = 1;
final int PERIFERIA      = 2;
final int EXTRA_URBANO   = 3;
final int VIA_AEROPUERTO = 4;

final int VALOR_MINIMO   = 6000;
final int VALOR_ESPECIAL = 1300;

final int VALOR_PERIFERIA      = 3400;
final int VALOR_EXTRA_URBANO   = 4900;
final int VALOR_VIA_AEROPUERTO = 1600;

void main() {
	var valorCobrado = Util.ingresarEntero("Valor cobrado: ");
	var valorTaximetro = Util.ingresarEntero("Valor del taxímetro: ");
	var esEspecial = Util.ingresarLogico("Es domingo, festivo o nocturno (s/n):");
	var destinoTrayecto = Util.ingresarOpcion("Destino \n"+
		"\t(1) urbano,\n"+
		"\t(2) lugar periférico,\n"+
		"\t(3) extra urbano\n"+
		"\t(4) vía al aeropuerto\n\n"+
		"Cuál es su destino: ", 4);

	var tarifaMinima = determinarTarifaMinima(valorTaximetro);
	var tarifaEspecial = determinarTarifaEspecial(esEspecial);
	var tarifaDestino = determinarTarifaDestino(destinoTrayecto);

	var tarifaReal = calcularTarifa(
		tarifaMinima, tarifaEspecial, tarifaDestino);

	var mensajeCobro = determinarMensajeCobro(valorCobrado, tarifaReal);
	
	var informeCobro = generarInformeCobro(
		valorCobrado, tarifaReal, mensajeCobro);

	Util.mostrarMensaje(informeCobro);
}

int determinarTarifaMinima(int valorTaximetro)  {
	return Math.max(valorTaximetro, VALOR_MINIMO);
}

int determinarTarifaEspecial(boolean esEspecial) {
	var tarifaEspecial = SIN_COSTO;

	if (esEspecial) {
		tarifaEspecial = VALOR_ESPECIAL;
	}

	return tarifaEspecial;
}

int determinarTarifaDestino(int destinoTrayecto) {
	var tarifaDestino = SIN_COSTO;

	if (destinoTrayecto == PERIFERIA) {
		tarifaDestino = VALOR_PERIFERIA;
	} else if (destinoTrayecto == EXTRA_URBANO) {
			tarifaDestino = VALOR_EXTRA_URBANO;
	} else if (destinoTrayecto == VIA_AEROPUERTO) {
			tarifaDestino = VALOR_VIA_AEROPUERTO;
	}

	return tarifaDestino;
}

int calcularTarifa(int tarifaMinima, int tarifaEspecial, int tarifaDestino){
	return tarifaMinima + tarifaEspecial + tarifaDestino;
}

String determinarMensajeCobro(int valorCobrado, int tarifaReal){
	var mensajeCobro = "JUSTO";

	if (valorCobrado != tarifaReal) {
		mensajeCobro = "INJUSTO";
	}

	return mensajeCobro;
}

String generarInformeCobro(int valorCobrado, int tarifaReal, String mensajeCobro) {
	return  String.format(
		"\nValor cobrado \t$%7d"+ 
			"\nValor real \t$%7d\n"+
			"\nPor lo anterior el cobro es %s\n",
		valorCobrado, tarifaReal, mensajeCobro);
}
