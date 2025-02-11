final double FACTOR_CARRO = 121.0;
final double FACTOR_MOTO = 53.0;
final double FACTOR_BICICLETA = 3.0;

final double FACTOR_BUS = 49.0;

final double FACTOR_CARNE_ROJA = 7.19;
final double FACTOR_PESCADO = 3.91;
final double FACTOR_VEGETARIANA = 3.81;

public void main() throws Exception {
  mostrarMensaje("KILÓMETROS RECORRIDOS EN VEHÍCULOS PARTICULARES\n");
  double kmCarro = ingresarReal("- carro: ");
  double kmMoto = ingresarReal("- moto: ");
  double kmBicicleta = ingresarReal("- bicicleta: ");

  mostrarMensaje("\nKILÓMETROS RECORRIDOS EN TRANSPORTE URBANO\n");
  double trayectosBus = ingresarReal("- bus: ");

  mostrarMensaje("\nCANTIDAD COMIDAS CON\n");
  int comidasCarneRoja = ingresarEntero("- carne roja: ");
  int comidasPescado = ingresarEntero("- pescado: ");
  int comidasVegetarianas = ingresarEntero("- vegetarianas: ");

  double kgCO2Recorridos = calcularKgCO2Recorridos(
    kmCarro, kmMoto, kmBicicleta);

  double kgCO2Trayectos = calcularKgCO2Trayectos(trayectosBus);

  double kgCO2Comidas = calcularKgCO2Comidas(
    comidasCarneRoja, comidasPescado, comidasVegetarianas);

  double kgCO2Total = calcularKgCo2Total(
    kgCO2Recorridos, kgCO2Trayectos, kgCO2Comidas);

  String mensajeSalida = generarMensajeSalida(
    kgCO2Recorridos, kgCO2Trayectos, kgCO2Comidas, kgCO2Total);

  mostrarMensaje(mensajeSalida);

  System.exit(0);
}

void mostrarMensaje(String mensaje) {
  System.out.print(mensaje);
}

String ingresarTexto(String mensaje) {
  var scanner = new Scanner(System.in);

  mostrarMensaje(mensaje);
  return scanner.nextLine();
}

int ingresarEntero(String mensaje) {
  return Integer.parseInt(ingresarTexto(mensaje));
}

double ingresarReal(String mensaje) {
  return Double.parseDouble(ingresarTexto(mensaje));
}

double calcularKgCO2Recorridos(
    double kmCarro,
    double kmMoto,
    double kmBicicleta) {
  return (kmCarro * FACTOR_CARRO +
      kmMoto * FACTOR_MOTO +
      kmBicicleta * FACTOR_BICICLETA);
}

double calcularKgCO2Trayectos(double trayectosBus) {
  return trayectosBus * FACTOR_BUS;
}

double calcularKgCO2Comidas(
    int comidasCarneRoja,
    int comidasPescado,
    int comidasVegetarianas) {
  return (
      comidasCarneRoja * FACTOR_CARNE_ROJA +
      comidasPescado * FACTOR_PESCADO +
      comidasVegetarianas * FACTOR_VEGETARIANA
  );
}

double calcularKgCo2Total(
    double kgCO2Recorridos,
    double kgCO2Trayectos,
    double kgCO2Comidas) {
  return (kgCO2Recorridos + kgCO2Trayectos + kgCO2Comidas);
}

String generarMensajeSalida(
    double kgCO2Recorridos,
    double kgCO2Trayectos,
    double kgCO2Comidas,
    double kgCO2Total) {
  return String.format(
      "\nSu huella de carbono en kg de CO2 por:\n" +
          "- RECORRIDOS %.2f kg\n" +
          "- TRAYECTOS %.2f kg\n" +
          "- COMIDAS %.2f kg\n\n" +
          "Para un TOTAL de %.2f kg de emisiones de CO2.",
      kgCO2Recorridos, kgCO2Trayectos, kgCO2Comidas, kgCO2Total);
}

main();
