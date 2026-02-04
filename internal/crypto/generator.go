package crypto

import (
	"crypto/rand"
	"math/big"
)

// Definición de conjuntos de caracteres
const (
	Letters      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers      = "0123456789"
	SpecialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

// GeneratePassword crea una contraseña basada en las preferencias del usuario
func GeneratePassword(length int, includeNumbers bool, includeSpecial bool) (string, error) {
	charset := Letters
	if includeNumbers {
		charset += Numbers
	}
	if includeSpecial {
		charset += SpecialChars
	}

	password := make([]byte, length)
	for i := range password {
		// Seleccionamos un índice aleatorio seguro
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[index.Int64()]
	}

	return string(password), nil
}
