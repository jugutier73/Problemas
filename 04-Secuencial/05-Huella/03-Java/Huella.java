/*
    Programa para calcular la cantidad de CO2 emitido por el uso de
    transporte particular (carro y moto), el empleo del transporte 
    público (buses)
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
*/

import modulo.Util;

final double HUELLA_CARRO = 121.0;
final double HUELLA_MOTO  = 53.0;
final double HUELLA_BUS   = 49.0;

void main() {
  var kmCarro = Util.ingresarReal("Total de kilómetros en carro: ");
  var kmMoto = Util.ingresarReal("Total de kilómetros en moto : ");
  var kmBuses = Util.ingresarReal("Total de kilómetros en buses: ");

  var huella = calcularHuellaCarbono(kmCarro, kmMoto, kmBuses);
  var informeHuella = generarHuella(kmCarro, kmMoto, kmBuses, huella);

  Util.mostrarMensaje(informeHuella);
}

double calcularHuellaCarbono(double kmCarro, double kmMoto, double kmBuses){
	return kmCarro * HUELLA_CARRO +
		kmMoto * HUELLA_MOTO +
		kmBuses * HUELLA_BUS;
}

String generarHuella(double kmCarro, double kmMoto, double kmBuses, double huella) {
	return String.format(
		"\nCon %.1f, %.1f, %.1f km de recorrido"+
			"\nen carro, moto y bus representante,"+
			"\nsu huella de carbono por el uso de"+
			"\ntransporte es de %.1f kg de CO2.\n",
		kmCarro, kmMoto, kmBuses, huella);
}
