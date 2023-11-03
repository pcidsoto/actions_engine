package engine

func obtenerNombreUsuario(param interface{}) map[string]interface{} {
	// Implementa la lógica para obtener el nombre del usuario aquí
	nombre := "Antonio" // Cambia esto según tu lógica real
	return map[string]interface{}{
		"nombre_usuario": nombre,
	}
}

func obtenerNumeroTelefono(param interface{}) map[string]interface{} {
	// Implementa la lógica para obtener el número de teléfono aquí
	telefono := "+569666666665" // Cambia esto según tu lógica real
	return map[string]interface{}{
		"numero_telefono": telefono,
	}
}

func obtenerEdadCliente(param interface{}) map[string]interface{} {
	edad := 30 // Simulación de obtener la edad del cliente
	return map[string]interface{}{
		"edad_cliente": edad,
		"nombre":       "Pedro",
	}
}
