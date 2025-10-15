/*
   Crear un programa para registrar el ingreso y salida de los
   usuarios de un laboratorio para asegurar que solo personas
   autorizadas -estudiantes, docentes o personal técnico- accedan
   al espacio. El programa debe permitir identificar quiénes estaban 
   presentes y la cantidad por cada tipo.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

final int ANCHO = 20;

final int TOTAL_TIPO_AUTORIZADOS = 5;

final int ESTUDIANTE = 1;
final int DOCENTE    = 2;
final int TECNICO    = 3;

final Map<Integer, String> ETIQUETAS = Map.of(
	ESTUDIANTE, "Estudiante",
	DOCENTE, "Docente",
	TECNICO, "Técnico"
);

record Ingreso(String documento, String nombre, int tipo) {}

void main() {
	var ingresos = Util.ingresarColeccion(() -> ingresarIngreso());
	var salidas  = Util.ingresarColeccion(() -> ingresarSalidas());
	
	var personasLaboratorio = obtenerPersonasLaboratorio(ingresos, salidas);

	var cantidadEstudiantes = Util.contarSegunCriterio(personasLaboratorio, 
		(ingreso, tipo) -> tenerTipo(ingreso, tipo), ESTUDIANTE);

	var cantidadDocentes = Util.contarSegunCriterio(personasLaboratorio, 
		(ingreso, tipo) -> tenerTipo(ingreso, tipo), DOCENTE);
	
	var cantidadTecnicos = Util.contarSegunCriterio(personasLaboratorio, 
		(ingreso, tipo) -> tenerTipo(ingreso, tipo), TECNICO);
	
	var reporteIngresos = generarReporteIngresos(
		ingresos,
		salidas,
		personasLaboratorio,
		cantidadEstudiantes,
		cantidadDocentes,
		cantidadTecnicos);

	Util.mostrarMensaje(reporteIngresos);
}

Ingreso ingresarIngreso(){
	Util.mostrarMensaje("\nIngrese los datos del ingreso:\n");
	var documento = Util.ingresarTexto("\tIngrese el documento   : ");
	var nombre    = Util.ingresarTexto("\tIngrese el nombre      : ");
	var tipo      = Util.ingresarOpcion(
		"\tTIPO DE PERSONAS AUTORIZADAS\n"+
			"\t\t1: Estudiante\n"+
			"\t\t2: Docente\n"+
			"\t\t3: Técnico\n"+
			"\tIngrese tipo de persona: ", 
			TOTAL_TIPO_AUTORIZADOS);

	return new Ingreso(documento, nombre, tipo);
}

String ingresarSalidas(){
	Util.mostrarMensaje("\nIngrese los datos de la salida:\n");
	var documento = Util.ingresarTexto("\tIngrese el documento   : ");

	return documento;
}

List<Ingreso> obtenerPersonasLaboratorio(List<Ingreso> ingresos, List<String> salidas) {
    var salidasSet = new HashSet<String>(salidas);

    return ingresos.stream()
        .filter(ingreso -> !salidasSet.contains(ingreso.documento()))
        .map(ingreso -> new Ingreso(ingreso.documento(), ingreso.nombre(), ingreso.tipo())) 
        .collect(Collectors.toList());
}

boolean tenerTipo(Ingreso ingresos,int tipo)  {
	return ingresos.tipo() == tipo;
}


String generarReporteIngresos(List<Ingreso> ingresos,
						   List<String> salidas,
						   List<Ingreso> personasLaboratorio,
						   int cantidadEstudiantes,
						   int cantidadDocentes,
						   int cantidadTecnicos){

	var listadoIngresos = Util.convertirColeccionCadena(
			"LISTADO DE INGRESOS", 
			ingresos,
			(ingreso) -> convertirIngresoCadena(ingreso));

	var listadoSalidas = Util.convertirColeccionCadena(
			"LISTADO DE SALIDAS", 
			salidas,
			(salida) -> convertirSalidaCadena(salida));

	var listadoPersonasLaboratorio = Util.convertirColeccionCadena(
			"LISTADO DE PERSONAS EN EL LABORATORIO", 
			personasLaboratorio,
			(ingreso) -> convertirIngresoCadena(ingreso));			

	return listadoIngresos + "\n" 
			+ listadoSalidas + "\n" 
			+ listadoPersonasLaboratorio + "\n" 
			+ "CANTIDAD DE PERSONAS AÚN EN EL LABORATORIO\n"
			+ "Cantidad de Estudiantes     : " + cantidadEstudiantes + "\n" 
			+ "Cantidad de Docentes        : " + cantidadDocentes + "\n"
			+ "Cantidad de Técnicos        : " + cantidadTecnicos + "\n";
}

String convertirIngresoCadena(Ingreso ingreso){
	var etiqueta = ETIQUETAS.get(ingreso.tipo());

	return String.format("%-" + ANCHO + "s %-" + ANCHO + "s %s\n", 
	ingreso.documento(), ingreso.nombre(), etiqueta);
}

String convertirSalidaCadena(String salida){
	return salida;
}