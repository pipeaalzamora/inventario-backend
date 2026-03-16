package shared

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

const (
	MaxBase64Size = 2831160 // 2.7MB en bytes
)

// Base64ToReader convierte un string base64 a un io.Reader
// Valida el tamaño máximo y el formato base64
func Base64ToReader(base64String string) (io.Reader, error) {
	// Validar tamaño
	if len(base64String) > MaxBase64Size {
		return nil, errors.New("la imagen excede el tamaño máximo permitido (2.7MB)")
	}

	// Validar que no esté vacío
	if len(base64String) == 0 {
		return nil, errors.New("el string base64 está vacío")
	}

	// Decodificar base64
	decoded, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, errors.New("formato base64 inválido")
	}

	return bytes.NewReader(decoded), nil
}

// GetExtensionFromMimeType obtiene la extensión de archivo desde un mimeType
func GetExtensionFromMimeType(mimeType string) string {
	switch mimeType {
	case "image/jpeg":
		return "jpg"
	case "image/png":
		return "png"
	case "image/webp":
		return "webp"
	default:
		return "jpg" // default
	}
}

// ValidateMimeType valida que el mimeType sea uno de los permitidos
func ValidateMimeType(mimeType string) bool {
	allowedTypes := []string{"image/jpeg", "image/png", "image/webp"}
	for _, allowed := range allowedTypes {
		if mimeType == allowed {
			return true
		}
	}
	return false
}

// CleanBase64String limpia el string base64 removiendo el prefijo data:image/...;base64, si existe
func CleanBase64String(base64String string) string {
	// Remover prefijo data:image/...;base64, si existe
	if idx := strings.Index(base64String, ","); idx != -1 {
		return base64String[idx+1:]
	}
	return base64String
}
