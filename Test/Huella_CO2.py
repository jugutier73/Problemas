FACTOR_CARRO = 121.0
FACTOR_MOTO = 53.0
FACTOR_BICICLETA = 3.0

FACTOR_BUS = 49.0

FACTOR_CARNE_ROJA = 7.19
FACTOR_PESCADO = 3.91
FACTOR_VEGETARIANA = 3.81


def main():
    mostrar_mensaje("KILÓMETROS RECORRIDOS EN VEHÍCULOS PARTICULARES")
    km_carro = ingresar_real("- carro: ")
    km_moto = ingresar_real("- moto: ")
    km_bicicleta = ingresar_real("- bicicleta eléctrica: ")

    mostrar_mensaje("\nKILÓMETROS RECORRIDOS EN TRANSPORTE URBANO")
    trayectos_bus = ingresar_real("- bus: ")

    mostrar_mensaje("\nCANTIDAD COMIDAS CON")
    comidas_carne_roja = ingresar_entero("- carne roja: ")
    comidas_pescado = ingresar_entero("- pescado: ")
    comidas_vegetarianas = ingresar_entero("- vegetarianas: ")

    kg_co2_recorridos = calcular_kg_co2_recorridos(km_carro,
                                                   km_moto,
                                                   km_bicicleta)

    kg_co2_trayectos = calcular_kg_co2_trayectos(trayectos_bus)

    kg_co2_comidas = calcular_kg_co2_comidas(comidas_carne_roja,
                                             comidas_pescado,
                                             comidas_vegetarianas)

    kg_co2_total = calcular_kg_co2_total(kg_co2_recorridos,
                                         kg_co2_trayectos,
                                         kg_co2_comidas)

    mensaje_salida = generar_mensaje_salida(kg_co2_recorridos,
                                            kg_co2_trayectos,
                                            kg_co2_comidas,
                                            kg_co2_total)

    mostrar_mensaje(mensaje_salida)


def mostrar_mensaje(mensaje):
    print(mensaje)


def ingresar_texto(pregunta):
    return input(pregunta)


def ingresar_entero(pregunta):
    return int(ingresar_texto(pregunta))


def ingresar_real(pregunta):
    return float(ingresar_texto(pregunta))


def calcular_kg_co2_recorridos(km_carro,
                               km_moto,
                               km_bicicleta):
    return (
        km_carro * FACTOR_CARRO +
        km_moto * FACTOR_MOTO +
        km_bicicleta * FACTOR_BICICLETA
    )


def calcular_kg_co2_trayectos(trayectos_bus):
    return trayectos_bus * FACTOR_BUS


def calcular_kg_co2_comidas(comidas_carne_roja,
                            comidas_carne_blanca,
                            comidas_vegetarianas):
    return (
        comidas_carne_roja * FACTOR_CARNE_ROJA +
        comidas_carne_blanca * FACTOR_PESCADO +
        comidas_vegetarianas * FACTOR_VEGETARIANA
    )


def calcular_kg_co2_total(kg_co2_recorridos,
                          kg_co2_trayectos,
                          kg_co2_comidas):
    return kg_co2_recorridos + kg_co2_trayectos + kg_co2_comidas


def generar_mensaje_salida(kg_co2_recorridos,
                           kg_co2_trayectos,
                           kg_co2_comidas,
                           kg_co2_total):
    return (
        f"\nSu huella de carbono en kg de CO2 por:"
        f"\n- RECORRIDOS {kg_co2_recorridos:.2f} kg"
        f"\n- TRAYECTOS {kg_co2_trayectos:.2f} kg"
        f"\n- COMIDAS {kg_co2_comidas:.2f} kg"
        f"\n\nPara un TOTAL de {kg_co2_total:.2f} "
        f"kg de emisiones de CO2."
    )


main()
