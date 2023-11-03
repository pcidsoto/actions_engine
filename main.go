package main

import (
	"engine/engine"
	"fmt"
)

func main() {
	jsonString := `{
    "Parametro_de_entrada": {
        "name": "userID",
        "value": 12345,
        "type": "entero"
    },
    "steps": [
        {
            "ObtenerDato": {
                "accion": "Edad_Cliente",
                "param": {
                    "value": "Parametro_de_entrada"
                },
                "retorno": {
                    "edad_cliente": "edad_cliente",
					"nombre": "nombre"
                }
            },
            "ActivarFiltro": [
				{
					"accion": "Igual que",
					"param": {
						"name": "edad_cliente",
						"type": "entero",
						"value": 30
					}
				}
			]
        },
        {
            "ObtenerDato": {
                "accion": "Edad_Cliente",
                "param": {
                    "value": "nombre"
                },
                "retorno": {
					"edad_cliente": "edad_cliente",
					"nombre": "nombre"
                }
            },
            "ActivarFiltro": [
				{
					"accion": "Igual que",
					"param": {
						"name": "nombre",
						"type": "string",
						"value": "Pedro"
					}
				}
			]
        },
        {
            "ObtenerDato": {
                "accion": "Nombre_Usuario",
                "param": {
                    "type": "string",
                    "value": "userID"
                },
                "retorno": {
                    "nombre_usuario": "nombre_usuario"
                }
            },
            "ActivarFiltro": [
				{
					"accion": "Igual que",
					"param": {
						"name": "nombre_usuario",
						"type": "string",
						"value": "Antonio"
					}
				}
			]
        },
        {
            "ObtenerDato": {
                "accion": "Numero_Telefono",
                "param": {
                    "type": "string",
                    "value": "userID"
                },
                "retorno": {
                    "numero_telefono": "numero_telefono"
                }
            },
            "ActivarFiltro": [
				{
					"accion": "Igual que",
					"param": {
						"name": "numero_telefono",
						"type": "string",
						"value": "+569666666665"
					}
				}
			]
        }
    ]
}
`

	resultado, detalle := engine.EvaluarJSON(jsonString)

	fmt.Printf("Resultado: %v, Detalle: %s\n", resultado, detalle)
}
