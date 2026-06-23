package shared

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Código creado a partir de la implementación del siguiente enlace:
// https://github.com/Pyxis-GMS/project-nikki-admin-api/blob/main/src/core/utils/crafter.ts

func CreatePassword(phrase, secret string) string {
	digest := legacyPasswordDigest(phrase, secret)
	hash, err := bcrypt.GenerateFromPassword([]byte(digest), bcrypt.DefaultCost)
	if err != nil {
		return digest
	}
	return string(hash)
}

func CheckPassword(phrase, hash, secret string) bool {
	digest := legacyPasswordDigest(phrase, secret)
	if IsBcryptPasswordHash(hash) {
		return bcrypt.CompareHashAndPassword([]byte(hash), []byte(digest)) == nil
	}
	return digest == hash
}

func IsBcryptPasswordHash(hash string) bool {
	return strings.HasPrefix(hash, "$2a$") || strings.HasPrefix(hash, "$2b$") || strings.HasPrefix(hash, "$2y$")
}

func PasswordNeedsRehash(hash string) bool {
	return !IsBcryptPasswordHash(hash)
}

func legacyPasswordDigest(phrase, secret string) string {
	combined := phrase + secret
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
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
