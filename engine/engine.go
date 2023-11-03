package engine

import (
	"encoding/json"
	"fmt"
)

var funcionesDisponibles = map[string]func(interface{}) map[string]interface{}{
	"Edad_Cliente":    obtenerEdadCliente,
	"Nombre_Usuario":  obtenerNombreUsuario,
	"Numero_Telefono": obtenerNumeroTelefono,
}

func EvaluarJSON(jsonString string) (bool, string) {
	var config Config
	var valorObtenido map[string]interface{}

	// Decodificar el JSON en la estructura Config
	if err := json.Unmarshal([]byte(jsonString), &config); err != nil {
		return false, "Error al decodificar el JSON"
	}

	// Mapa para almacenar valores intermedios
	valores := make(map[string]interface{})

	// Mapa para registrar si una función ya ha sido ejecutada
	ejecutado := make(map[string]map[string]interface{})

	// Iterar sobre cada paso
	for _, paso := range config.Steps {
		if paso.ObtenerDato.Accion != "" {
			funcion, existe := funcionesDisponibles[paso.ObtenerDato.Accion]
			if !existe {
				return false, "La función solicitada no está disponible"
			}

			var valorEntrada interface{}

			if paso.ObtenerDato.Param.Value == "Parametro_de_entrada" {
				// Obtener valor de config.ParametroDeEntrada.Value
				valorEntrada = config.ParametroDeEntrada.Value
			} else {
				// Obtener valor del mapa de valores intermedios
				fmt.Println("Obteniendo valor de entrada")
				valorEntrada = valores[paso.ObtenerDato.Param.Name]
			}

			if val, ok := ejecutado[paso.ObtenerDato.Accion]; ok {
				fmt.Println("Obteniendo del caché: ", paso.ObtenerDato.Accion)
				valorObtenido = val
			} else {
				fmt.Println("Ejecutando funcion: ", paso.ObtenerDato.Accion)
				valorObtenido = funcion(valorEntrada)
			}

			for _, filtro := range paso.ActivarFiltro {
				resultado, detalle := evaluarFiltro(valorObtenido, filtro.Param.Name, filtro)

				if !resultado {
					return resultado, detalle
				}
			}

			// Almacenar el valor obtenido en el mapa de valores intermedios
			for key, _ := range paso.ObtenerDato.Retorno {
				valores[key] = valorObtenido[key]
			}

			ejecutado[paso.ObtenerDato.Accion] = valorObtenido
		}
	}

	return true, "Todas las acciones fueron exitosas"
}
