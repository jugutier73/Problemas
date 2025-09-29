/*
   Programa que registra los residuos recolectados durante una
   jornada, especificando su código, tipo y condición de
   reparabilidad. El sistema genera un informe con el
   listado de los elementos recolectados, ordenados por tipo y
   reparabilidad, además de presentar un conteo total por cada
   tipo de residuo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

final int ANCHO = 15;

final int TOTAL_TIPOS = 5;

final int BATERIA          = 1;
final int TELEFONO         = 2;
final int COMPUTADOR       = 3;
final int CARGADOR         = 4;
final int OTRO_DISPOSITIVO = 5;


final Map<Integer, String> ETIQUETAS = Map.of(
	BATERIA, "Batería",
	TELEFONO, "Teléfono",
	COMPUTADOR, "Computador",
	CARGADOR, "Cargador",
	OTRO_DISPOSITIVO, "Otro"
);

record Residuo(String codigo, int tipo, boolean reparable) {}

void main() {
	var residuos = Util.ingresarColeccion(() -> ingresarResiduo());

    var residuosClasificados = Util.ordenarColeccion(residuos, 
		compararCampoTipo(), false);

	var cantidadBaterias = Util.contarSegunCriterio(residuos, 
		(residuo, tipo) -> tenerTipo(residuo, tipo), BATERIA);
	
	var cantidadTelefonos = Util.contarSegunCriterio(residuos, 
		(residuo, tipo) -> tenerTipo(residuo, tipo), TELEFONO);

	var cantidadComputadores = Util.contarSegunCriterio(residuos, 
		(residuo, tipo) -> tenerTipo(residuo, tipo), COMPUTADOR);

	var cantidadCargadores = Util.contarSegunCriterio(residuos, 
		(residuo, tipo) -> tenerTipo(residuo, tipo), CARGADOR);

	var cantidadOtros = Util.contarSegunCriterio(residuos, 
		(residuo, tipo) -> tenerTipo(residuo, tipo), OTRO_DISPOSITIVO);

	var reporteRecoleccion = generarReporteRecoleccion(
		residuosClasificados,
		cantidadBaterias,
		cantidadTelefonos,
		cantidadComputadores,
		cantidadCargadores,
		cantidadOtros);

	Util.mostrarMensaje(reporteRecoleccion);
}

Residuo ingresarResiduo(){
	Util.mostrarMensaje("\nIngrese los datos del residuo electrónico:\n");
	var codigo = Util.ingresarTexto("\tIngrese el código del residuo: ");
	var tipo = Util.ingresarOpcion(
		"\tTIPOS DE RESIDUOS\n"+
			"\t\t1: Batería\n"+
			"\t\t2: Teléfono\n"+
			"\t\t3: Computador\n"+
			"\t\t4: Cargador\n"+
			"\t\t5: Otro dispositivo)\n"+
			"\tIngrese tipo de residuo      : ", 
			TOTAL_TIPOS);
	var reparable = Util.ingresarLogico("\tPuede ser reparado (s/n): ");

	return new Residuo(codigo, tipo, reparable);
}

Comparator<Residuo> compararCampoTipo() {
    return Comparator.comparingInt(Residuo::tipo)
			.thenComparing(Residuo::reparable, Comparator.reverseOrder());
}

boolean tenerTipo(Residuo residuo,int tipoInteres)  {
	return residuo.tipo() == tipoInteres;
}

String generarReporteRecoleccion(List<Residuo> residuosClasificados,
						   int cantidadBaterias,
						   int cantidadTelefonos,
						   int cantidadComputadores,
						   int cantidadCargadores,
						   int cantidadOtros){

	var listadoResiduosClasificados = Util.convertirColeccionCadena(
			"LISTADO DE RESIDUOS CLASIFICADOS", 
			residuosClasificados,
			(residuo) -> convertirResiduoCadena(residuo));

	return listadoResiduosClasificados + "\n" 
			+ "Cantidad de Baterías     : " + cantidadBaterias + "\n" 
			+ "Cantidad de Teléfonos    : " + cantidadTelefonos + "\n"
			+ "Cantidad de Computadores : " + cantidadComputadores + "\n"
			+ "Cantidad de Cargadores   : " + cantidadCargadores + "\n"
			+ "Cantidad de Otros        : " + cantidadOtros + "\n";
}

String convertirResiduoCadena(Residuo residuo){
	var etiqueta = ETIQUETAS.get(residuo.tipo());
	var reparable = residuo.reparable() ? "(REPARABLE)" : "";

	return String.format("%-" + ANCHO + "s %-" + ANCHO + "s %s\n", 
	residuo.codigo(), etiqueta, reparable);
}
