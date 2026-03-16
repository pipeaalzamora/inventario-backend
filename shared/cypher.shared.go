package shared

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	// "golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"errors"
	"math/big"
)

// Código creado a partir de la implementación del siguiente enlace:
// https://github.com/Pyxis-GMS/project-nikki-admin-api/blob/main/src/core/utils/crafter.ts

// Dejo las 2 versiones de CreatePassword y CheckPassword comentadas para que puedas elegir la que prefieras.
// La primera usa bcrypt y la segunda usa sha256 para generar un hash de la contraseña.
// Ambas son seguras, pero tienen diferentes características de rendimiento y almacenamiento.
// Si decides usar bcrypt, asegúrate de importar el paquete "golang.org/x/crypto/bcrypt".
// Si decides usar sha256, no necesitas importar ningún paquete adicional.
// Si decides usar bcrypt, puedes descomentar las siguientes dos funciones y comentar las de abajo
// que usan sha256.

// Para claves de acceso, bcrypt es generalmente preferido por su resistencia a ataques de fuerza bruta.

// func CreatePassword(phrase, secret string) (string, error) {
// 	combined := phrase + secret
// 	hash, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hash), nil
// }

// func CheckPassword(phrase, hash, secret string) bool {
// 	combined := phrase + secret
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(combined))
// 	return err == nil
// }

func CreatePassword(phrase, secret string) string {
	combined := phrase + secret
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
}

func CheckPassword(phrase, hash, secret string) bool {
	expected := CreatePassword(phrase, secret)
	return expected == hash
}

func CreateRandomNumber(length int) (string, error) {
	const ref = "0123456789"
	return randomizeString(ref, length)
}

func CreateRandomString(length int) (string, error) {
	const ref = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHYJKLMNOPQRSTUVWXYZ!_?.-*[]{}"
	return randomizeString(ref, length)
}

func CreateRandomURLString(length int) (string, error) {
	const ref = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHYJKLMNOPQRSTUVWXYZ-"
	return randomizeString(ref, length)
}

func CreateSaltyHash(phrase string) string {
	combined := phrase + "sofia-backend"
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
}

func CompareSaltyHash(phrase, hash string) bool {
	expected := CreateSaltyHash(phrase)
	return expected == hash
}

func randomizeString(ref string, length int) (string, error) {
	if length <= 0 {
		length = len(ref)
	}
	if len(ref) == 0 {
		return "", errors.New("reference string cannot be empty")
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(ref))))
		if err != nil {
			return "", err
		}
		result[i] = ref[index.Int64()]
	}

	return string(result), nil
}

func HashMapGeneric(data map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(jsonBytes)
	return hex.EncodeToString(hash[:]), nil
}
