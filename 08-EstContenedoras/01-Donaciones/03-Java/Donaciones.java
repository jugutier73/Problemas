/*
   Crear un programa para registrar y consultar de manera ágil
   todas las donaciones recibidas. Esta debería ofrecer el total 
   recaudado, listados alfabéticos y descendentes por valor, 
   identificar a los aportantes que superen un umbral específico y 
   señalar al mayor donante.

   Autor: Julián Esteban Gutiérrez Posada
   Fecha: Diciembre 2025
   Licencia: GNU GPL v3
*/

import modulo.Util;

final int UMBRAL = 10000;

record Donante(String nombre, int donacion) {}

void main() {
    var donantes = ingresarColeccion(() -> ingresarDonante());

    var donantesPorNombre = ordenarColeccion(donantes, 
        compararNombre(), false);
    var donantesPorDonacion = ordenarColeccion(donantes,
        compararDonacion(), true);

    var mayoresDonantes = obtenerMayoresDonaciones(donantes, UMBRAL);
    var mayorDonante = obtenerMayorDonante(donantes);
    var sumaDonaciones = obtenerSumaDonaciones(donantes);

    var reporte_donantes  = generarReporteDonaciones(
        donantesPorNombre, 
        donantesPorDonacion, 
        mayoresDonantes, 
        mayorDonante, 
        sumaDonaciones);

    Util.mostrarMensaje(reporte_donantes);
}

<Tipo> List<Tipo> ingresarColeccion(Supplier<Tipo> ingresarElemento){
    var coleccion = new ArrayList<Tipo>();
    var hayMas = true;

    while (hayMas){
        coleccion.add(ingresarElemento.get());

        hayMas = Util.ingresarLogico( "Hay más datos (s/n): ");
    }

    return coleccion;
}

Donante ingresarDonante(){
    Util.mostrarMensaje("\nIngrese los datos de un donante:\n");
    var nombre    = Util.ingresarTexto("\tIngrese el nombre: ");
    var donacion  = Util.ingresarEntero("\tIngrese donación : ");

    return new Donante(nombre, donacion);
}

<Tipo> List<Tipo>  ordenarColeccion(List<Tipo> coleccion, 
                                    Comparator<Tipo> comparador, 
                                    boolean descendente){
    return coleccion.stream()
            .sorted(descendente ? comparador.reversed() : comparador)
            .toList();
}

Comparator<Donante> compararNombre() {
    return Comparator.comparing(Donante::nombre);
}

Comparator<Donante> compararDonacion() {
    return Comparator.comparingInt(Donante::donacion);
}

List<Donante> obtenerMayoresDonaciones(List<Donante> donantes, 
                                       int valorLimite){
    return donantes.stream()
            .filter(d -> d.donacion() > valorLimite)
            .toList();
}

Donante obtenerMayorDonante(List<Donante> donantes){
    return Collections.max(donantes, 
        Comparator.comparingInt(Donante::donacion));
}

int obtenerSumaDonaciones(List<Donante> donantes){
    return donantes.stream()
        .mapToInt(Donante::donacion)
        .sum();
}

String generarReporteDonaciones(List<Donante> donantesPorNombre, 
                                List<Donante> donantesPorDonacion, 
                                List<Donante> mayoresDonantes, 
                                Donante mayorDonante, 
                                int sumaDonaciones){
    var listadoPorNombre = convertirColeccionCadena(
        "LISTADO EN ORDEN ALFABÉTICO", 
        donantesPorNombre,
        (donante) -> convertirDonanteCadena(donante));

    var listadoPorDonacion = convertirColeccionCadena(
        "LISTADO ORDENADO POR DONACIÓN", 
        donantesPorDonacion,
        (donante) -> convertirDonanteCadena(donante));

    var listadoMayores = convertirColeccionCadena(
        "LISTADO DONACIONES MAYORES A $" + UMBRAL, 
        mayoresDonantes,
        (donante) -> convertirDonanteCadena(donante));

    return listadoPorNombre + "\n" 
            + listadoPorDonacion + "\n" 
            + listadoMayores + "\n" 
            + "El mayor donante: " + mayorDonante.nombre() + "\n" 
            + "Total de donantes $" + sumaDonaciones + "\n";
}

<Tipo> String convertirColeccionCadena(String titulo, 
                                       List<Tipo> coleccion, 
                                       Function<Tipo,String> convertirElementoCadena){
    var mensaje = "\n" + titulo + "\n";

    for (var elemento : coleccion) {
        mensaje += "\t" + convertirElementoCadena.apply(elemento);
    }

    return mensaje;
}

String convertirDonanteCadena(Donante donante){
    return String.format("$%10d \t %s%n", donante.donacion(), donante.nombre());
}