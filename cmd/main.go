package main

import (
	"fmt"
	"os"
	// Usaremos una librería estándar para leer .env más adelante,
	// por ahora solo probamos el flujo básico.
)

func main() {
	fmt.Println("--- Password Manager Core ---")

	// Simulacro de carga de configuración
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development (default)"
	}

	fmt.Printf("Entorno actual: %s\n", appEnv)
	fmt.Println("Estado: Inicializado correctamente.")

	// Aquí es donde llamarás a tus funciones de internal/crypto e internal/ui
	fmt.Println("\nPróximo paso: Implementar el generador de contraseñas...")
}
