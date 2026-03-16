package shared

import (
	"strconv"
	"strings"
)

func IsValidRUT(rut string) bool {
	// Limpieza básica
	rut = strings.TrimSpace(rut)
	if rut == "" {
		return false
	}

	// Eliminar puntos y guión
	rut = strings.ReplaceAll(rut, ".", "")
	rut = strings.ReplaceAll(rut, "-", "")
	rut = strings.ToUpper(rut)

	// Debe tener al menos cuerpo + dígito verificador
	if len(rut) < 2 {
		return false
	}

	body := rut[:len(rut)-1]
	dv := rut[len(rut)-1:]

	// Validar que el cuerpo sea numérico
	for _, c := range body {
		if c < '0' || c > '9' {
			return false
		}
	}

	// Cálculo del dígito verificador
	sum := 0
	multiplier := 2

	// Recorremos el cuerpo desde el último dígito hacia atrás
	for i := len(body) - 1; i >= 0; i-- {
		digit := int(body[i] - '0')
		sum += digit * multiplier
		multiplier++
		if multiplier > 7 {
			multiplier = 2
		}
	}

	remainder := sum % 11
	check := 11 - remainder

	var computedDV string
	switch check {
	case 11:
		computedDV = "0"
	case 10:
		computedDV = "K"
	default:
		computedDV = strconv.Itoa(check)
	}

	return computedDV == dv
}
