/*
    Programa para calcular la cantidad de CO2 emitido por el uso de
    transporte particular (carro y moto), el empleo del transporte 
    público (buses)
    Autor: Julián Esteban Gutiérrez Posada
    Fecha: Marzo 2025
    Licencia: GNU GPL v3
*/

const HUELLA_CARRO = 121.0;
const HUELLA_MOTO  = 53.0;
const HUELLA_BUS   = 49.0;

function main() {
  let kmCarro = ingresarReal('kmCarro');
  let kmMoto = ingresarReal('kmMoto');
  let kmBuses = ingresarReal('kmBuses');

	let huella = calcularHuellaCarbono(kmCarro, kmMoto, kmBuses);
	let informeHuella = generarHuella(kmCarro, kmMoto, kmBuses, huella);

	mostrarMensaje(informeHuella);
}

function calcularHuellaCarbono(kmCarro, kmMoto, kmBuses) {
	return kmCarro * HUELLA_CARRO +
		kmMoto * HUELLA_MOTO +
		kmBuses * HUELLA_BUS;
}

function generarHuella(kmCarro, kmMoto, kmBuses, huella) {
	return `\nCon ${kmCarro.toFixed(1)}, ${kmMoto.toFixed(1)}, ${kmBuses.toFixed(1)} km de recorrido`+
         `\nen carro, moto y bus representante,`+
			   `\nsu huella de carbono por el uso de`+
			   `\ntransporte es de ${huella.toFixed(1)} kg de CO2.\n`;
}
