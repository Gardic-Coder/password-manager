package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gardic-coder/password-manager/internal/crypto"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Cargar .env
	if err := godotenv.Load(); err != nil {
		// Solo imprime si el error NO es que el archivo no existe
		if !os.IsNotExist(err) {
			log.Printf("Error cargando .env: %v", err)
		}
	}

	fmt.Println("--- Password Manager Core ---")

	// 2. Probar Generador
	// Generar contraseña de 16 caracteres, con números y símbolos
	nuevaPass, err := crypto.GeneratePassword(16, true, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contraseña generada: %s\n", nuevaPass)
	fmt.Printf("DB Path desde .env: %s\n", os.Getenv("DB_PATH"))
}
