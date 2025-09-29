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

const URBANO         = 1;
const PERIFERIA      = 2;
const EXTRA_URBANO   = 3;
const VIA_AEROPUERTO = 4;

const SIN_COSTO      = 0;
const VALOR_MINIMO   = 6000;
const VALOR_ESPECIAL = 1300;

const VALOR_PERIFERIA      = 3400;
const VALOR_EXTRA_URBANO   = 4900;
const VALOR_VIA_AEROPUERTO = 1600;

function main() {
	let valorCobrado = ingresarEntero('valorCobrado');
	let valorTaximetro = ingresarEntero('valorTaximetro');
	let esEspecial = ingresarLogico('esEspecial');
	let destinoTrayecto = ingresarOpcion('destinoTrayecto');

	let tarifaMinima = determinarTarifaMinima(valorTaximetro);
	let tarifaEspecial = determinarTarifaEspecial(esEspecial);
	let tarifaDestino = determinarTarifaDestino(destinoTrayecto);

	let tarifaReal = calcularTarifa(
		tarifaMinima, tarifaEspecial, tarifaDestino);

	let mensajeCobro = determinarMensajeCobro(valorCobrado, tarifaReal);
	
	let informeCobro = generarInformeCobro(
		valorCobrado, tarifaReal, mensajeCobro);

	mostrarMensaje(informeCobro);
}

function determinarTarifaMinima(valorTaximetro)  {
	return Math.max(valorTaximetro, VALOR_MINIMO);
}

function determinarTarifaEspecial(esEspecial) {
	let tarifaEspecial = SIN_COSTO;

	if (esEspecial) {
		tarifaEspecial = VALOR_ESPECIAL;
	}

	return tarifaEspecial;
}

function determinarTarifaDestino(destinoTrayecto) {
	let tarifaDestino = SIN_COSTO;

	if (destinoTrayecto == PERIFERIA) {
		tarifaDestino = VALOR_PERIFERIA;
	} else if (destinoTrayecto == EXTRA_URBANO) {
			tarifaDestino = VALOR_EXTRA_URBANO;
	} else if (destinoTrayecto == VIA_AEROPUERTO) {
			tarifaDestino = VALOR_VIA_AEROPUERTO;
	}

	return tarifaDestino;
}

function calcularTarifa(valorTaximetro, esEspecial, destinoTrayecto) {
	return valorTaximetro + esEspecial + destinoTrayecto;
}

function determinarMensajeCobro(valorCobrado, tarifaReal){
	let mensajeCobro = "JUSTO";

	if (valorCobrado != tarifaReal) {
		mensajeCobro = "INJUSTO";
	}

	return mensajeCobro;
}

function generarInformeCobro(valorCobrado, tarifaReal, mensajeCobro) {
	return  `\nValor cobrado\t$${valorCobrado.toString().padStart(7, ' ')}` +
            `\nValor real \t$${tarifaReal.toString().padStart(7, ' ')}\n` +
            `\nPor lo anterior el cobro es ${mensajeCobro}\n`;
}