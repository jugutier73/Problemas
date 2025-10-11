/*
    Programa que permita tener un control del consumo de la energía
    eléctrica con relación al consumo del mes anterior, para que 
    así el usuario pueda adoptar hábitos de consumo más conscientes 
    y responsables.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
*/

import modulo.Util;

void main() {
	var consumoActual = Util.ingresarEntero(
        "Consumo mes actual   (kilovatios): ");
	var consumoAnterior = Util.ingresarEntero(
        "Consumo mes anterior (kilovatios): ");

	var relacionConsumo = calcularRelacionConsumo(consumoActual,
                                                consumoAnterior);

	var reporteRelacion = generarReporteRelacion(consumoActual, 
                                               consumoAnterior,
                                               relacionConsumo);
        
	Util.mostrarMensaje(reporteRelacion);
}

double calcularRelacionConsumo(int consumoActual, int consumoAnterior) {
	return (double)consumoActual / (double)consumoAnterior * 100.0;
}

String generarReporteRelacion(int consumoActual, int consumoAnterior, 
double relacionConsumo) {
	return String.format(
    "\nEl consumo actual de %d kilovatios representa"+
		"\nun %.1f%% con relación al consumo del mes"+
		"\nanterior de %d kilovatios.\n", 
    consumoActual, relacionConsumo, consumoAnterior);
}