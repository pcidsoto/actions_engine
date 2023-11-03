package engine

import (
	"fmt"
)

func evaluarFiltro(valorObtenido map[string]interface{}, name string, filtro ActivarFiltro) (bool, string) {
	valorFloat, err := normalizeToFloat64(valorObtenido[name])
	if err != nil {
		// Si no se puede normalizar a float64, trata de comparar como strings
		valorStr, ok := valorObtenido[name].(string)
		if !ok {
			return false, "No se puede normalizar en Float64 y no es un string"
		}

		return compararCadena(valorStr, filtro.Param.Value.(string), filtro.Accion)
	}

	switch filtro.Accion {
	case "Igual que":
		return compararNumeros(valorFloat, filtro.Param.Value.(float64), "==")
	case "Mayor que":
		return compararNumeros(valorFloat, filtro.Param.Value.(float64), ">")
	case "Menor que":
		return compararNumeros(valorFloat, filtro.Param.Value.(float64), "<")
	case "Mayor o igual":
		return compararNumeros(valorFloat, filtro.Param.Value.(float64), ">=")
	case "Menor o igual":
		return compararNumeros(valorFloat, filtro.Param.Value.(float64), "<=")
	case "Diferente":
		return compararNumeros(valorFloat, filtro.Param.Value.(float64), "!=")
	default:
		return false, "Operador no soportado"
	}
}

func compararCadena(valorObtenido string, valorEsperado string, operador string) (bool, string) {
	switch operador {
	case "Igual que":
		if valorObtenido == valorEsperado {
			return true, "El valor es igual (comparación de strings)"
		}
	case "Diferente que":
		if valorObtenido != valorEsperado {
			return true, "El valor es diferente (comparación de strings)"
		}
	default:
		fmt.Println("valorObtenido", valorEsperado, " valorEsperado", valorEsperado, " operador", operador)
		return false, "Operador de cadena no soportado"
	}

	return false, fmt.Sprintf("El valor no cumple la condición. Valor esperado: %v, Valor obtenido: %v", valorEsperado, valorObtenido)
}

func compararNumeros(valorObtenido float64, valorEsperado float64, operador string) (bool, string) {
	switch operador {
	case "==":
		if valorObtenido == valorEsperado {
			return true, "El valor es igual (comparación de números)"
		}
	case "!=":
		if valorObtenido != valorEsperado {
			return true, "El valor es diferente (comparación de números)"
		}
	case ">":
		if valorObtenido > valorEsperado {
			return true, "El valor es mayor (comparación de números)"
		}
	case "<":
		if valorObtenido < valorEsperado {
			return true, "El valor es menor (comparación de números)"
		}
	case ">=":
		if valorObtenido >= valorEsperado {
			return true, "El valor es mayor o igual (comparación de números)"
		}
	case "<=":
		if valorObtenido <= valorEsperado {
			return true, "El valor es menor o igual (comparación de números)"
		}
	default:
		return false, "Operador numérico no soportado"
	}

	return false, fmt.Sprintf("El valor no cumple la condición. Valor esperado: %v, Valor obtenido: %v", valorEsperado, valorObtenido)
}

func normalizeToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("No se puede normalizar el tipo %T a float64", value)
	}
}
