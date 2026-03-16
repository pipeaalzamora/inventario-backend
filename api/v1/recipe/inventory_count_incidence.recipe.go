package recipe

import (
	"encoding/json"
)

// NullableString permite distinguir entre campo omitido, null explícito y valor presente
type NullableString struct {
	Value          *string
	IsExplicitNull bool
}

// UnmarshalJSON implementa json.Unmarshaler para distinguir entre omitido, null y valor
func (n *NullableString) UnmarshalJSON(data []byte) error {
	// Si el campo no está presente, no se llama este método
	// Si el campo es null explícito, data será "null"
	if string(data) == "null" {
		n.Value = nil
		n.IsExplicitNull = true
		return nil
	}

	// Si el campo es un string, deserializarlo
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	n.Value = &str
	n.IsExplicitNull = false
	return nil
}

type RecipeInventoryCountIncidence struct {
	Items []RecipeIncidenceItem `json:"items" binding:"required"`
}

// UnmarshalJSON permite aceptar tanto array directo como objeto con campo items
func (r *RecipeInventoryCountIncidence) UnmarshalJSON(data []byte) error {
	// Primero intentar como array directo
	var items []RecipeIncidenceItem
	if err := json.Unmarshal(data, &items); err == nil {
		r.Items = items
		return nil
	}

	// Si falla, intentar como objeto con campo items
	var obj struct {
		Items []RecipeIncidenceItem `json:"items"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	r.Items = obj.Items
	return nil
}

type RecipeIncidenceItem struct {
	IDProducto  string              `json:"idProducto" binding:"required,uuid"`
	Incidencias RecipeIncidenceData `json:"incidencias" binding:"required"`
	Cantidades  []RecipeQuantity    `json:"cantidades" binding:"required"`
}

type RecipeIncidenceData struct {
	Observaciones *string       `json:"observaciones,omitempty"` // Opcional - puede ser vacío o nil
	Imagen        NullableString `json:"imagen,omitempty"`        // Opcional - null explícito indica eliminar imagen
	MimeType      *string       `json:"mimeType,omitempty"`      // Opcional - solo requerido si se envía imagen
}

type RecipeQuantity struct {
	UnitID  int     `json:"unitId" binding:"required"`
	UnitAbv string  `json:"unitAbv" binding:"required"`
	Count   float32 `json:"count" binding:"required"`
	Factor  float32 `json:"factor" binding:"required"`
}
