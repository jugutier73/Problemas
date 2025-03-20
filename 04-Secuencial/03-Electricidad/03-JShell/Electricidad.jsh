/*
    Programa que permita tener un control del consumo de la energía
    eléctrica con relación al consumo del mes anterior, para que 
    así el usuario pueda adoptar hábitos de consumo más conscientes 
    y responsables.
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
*/

void main() {
	var consumoActual = ingresarEntero(
        "Consumo mes actual   (kilovatios): ");
	var consumoAnterior = ingresarEntero(
        "Consumo mes anterior (kilovatios): ");
	var relacionConsumo = calcularRelacionConsumo(consumoActual, consumoAnterior);
	var reporteRelacion = generarReporeRelacion(consumoActual, 
        consumoAnterior, relacionConsumo);
	mostrarMensaje(reporteRelacion);

	System.exit(0);
}

double calcularRelacionConsumo(int consumoActual, int consumoAnterior) {
	return (double)consumoActual / (double)consumoAnterior * 100.0;
}

String generarReporeRelacion(int consumoActual, int consumoAnterior, 
double relacionConsumo) {
	return String.format(
    "\nEl consumo actual de %d kilovatios representa"+
		"\nun %.1f%% con relación al consumo del mes"+
		"\nanterior de %d kilovatios.", 
    consumoActual, relacionConsumo, consumoAnterior);
}

main();