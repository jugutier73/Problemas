package main

import (
	"fmt"
)

const (
	FACTOR_CARRO     = 121.0
	FACTOR_MOTO      = 53.0
	FACTOR_BICICLETA = 3.0

	FACTOR_BUS = 49.0

	FACTOR_CARNE_ROJA  = 7.19
	FACTOR_PESCADO     = 3.91
	FACTOR_VEGETARIANA = 3.81
)

func main() {
	mostrarMensaje("KILÓMETROS RECORRIDOS EN VEHÍCULOS PARTICULARES")
	kmCarro := ingresarReal("- carro: ")
	kmMoto := ingresarReal("- moto: ")
	kmBicicleta := ingresarReal("- bicicleta: ")

	mostrarMensaje("\nKILÓMETROS RECORRIDOS EN TRANSPORTE URBANO")
	trayectosBus := ingresarReal("- bus: ")

	mostrarMensaje("\nCANTIDAD COMIDAS CON")
	comidasCarneRoja := ingresarEntero("- carne roja: ")
	comidasPescado := ingresarEntero("- pescado: ")
	comidasVegetarianas := ingresarEntero("- vegetarianas: ")

	kgCO2Recorridos := calcularKgCO2Recorridos(
		kmCarro,
		kmMoto,
		kmBicicleta)

	kgCO2Trayectos := calcularKgCO2Trayectos(trayectosBus)

	kgCO2Comidas := calcularKgCO2Comidas(
		comidasCarneRoja,
		comidasPescado,
		comidasVegetarianas)

	kgCO2Total := calcularKgCO2Total(
		kgCO2Recorridos,
		kgCO2Trayectos,
		kgCO2Comidas)

	mensajeSalida := generarMensajeSalida(
		kgCO2Recorridos,
		kgCO2Trayectos,
		kgCO2Comidas,
		kgCO2Total)

	mostrarMensaje(mensajeSalida)
}

func mostrarMensaje(mensajeSalida string) {
	fmt.Println(mensajeSalida)
}

func ingresarReal(pregunta string) float64 {
	var real float64
	fmt.Print(pregunta)
	fmt.Scanln(&real)
	return real
}

func ingresarEntero(pregunta string) int {
	var entero int
	fmt.Print(pregunta)
	fmt.Scanln(&entero)
	return entero
}

func calcularKgCO2Recorridos(
	kmCarro,
	kmMoto,
	kmBicicleta float64) float64 {
	return kmCarro*FACTOR_CARRO + kmMoto*FACTOR_MOTO +
		kmBicicleta*FACTOR_BICICLETA
}

func calcularKgCO2Trayectos(trayectosBus float64) float64 {
	return float64(trayectosBus) * FACTOR_BUS
}

func calcularKgCO2Comidas(
	comidasCarneRoja,
	comidasCarneBlanca,
	comidasVegetarianas int) float64 {
	return float64(comidasCarneRoja)*FACTOR_CARNE_ROJA +
		float64(comidasCarneBlanca)*FACTOR_PESCADO +
		float64(comidasVegetarianas)*FACTOR_VEGETARIANA
}

func calcularKgCO2Total(
	kgCO2Recorridos,
	kgCO2Trayectos,
	kgCO2Comidas float64) float64 {
	return kgCO2Recorridos + kgCO2Trayectos + kgCO2Comidas
}

func generarMensajeSalida(kgCO2Recorridos,
	kgCO2Trayectos,
	kgCO2Comidas,
	kgCO2Total float64) string {
	return fmt.Sprintf("\nSu huella de carbono en kg de CO2 por:\n"+
		"- RECORRIDOS %.2f kg\n"+
		"- TRAYECTOS %.2f kg\n"+
		"- COMIDAS %.2f kg\n"+
		"\nPara un TOTAL de %.2f kg de emisiones de CO2.",
		kgCO2Recorridos, kgCO2Trayectos, kgCO2Comidas, kgCO2Total)
}
