package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
)

type DomainError struct {
	Message string
}

func (e DomainError) Error() string {
	return e.Message
}

type DataError struct {
	Message string
}

func (e DataError) Error() string {
	return e.Message
}

type PowerError struct {
	Message string
}

func (e PowerError) Error() string {
	return e.Message
}

type ValidationErrorItem struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error, structType any) []ValidationErrorItem {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		// validación estructural (reglas como min, required, etc.)
		t := reflect.TypeOf(structType)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		out := make([]ValidationErrorItem, len(ve))
		for i, fe := range ve {
			param := extractParamName(fe, t)
			message := buildMessage(fe)
			out[i] = ValidationErrorItem{
				Param:   param,
				Message: message,
			}
		}
		return out
	}

	// errores de tipo
	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		return []ValidationErrorItem{
			{
				Param:   ute.Field,
				Message: fmt.Sprintf("Invalid type for %s: expected %v, but got %q", ute.Field, ute.Type, ute.Value),
			},
		}
	}

	// errores de sintaxis
	var se *json.SyntaxError
	if errors.As(err, &se) {
		return []ValidationErrorItem{
			{
				Param:   "",
				Message: "Syntax error: " + err.Error(),
			},
		}
	}

	// fallback general
	return []ValidationErrorItem{}
}

// extrae el nombre del campo desde la tag `form` o `json` si está disponible
func extractParamName(fe validator.FieldError, structType any) string {
	// Obtenemos el tipo reflejado del struct original (no el campo)
	rt := reflect.TypeOf(structType)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if field, ok := rt.FieldByName(fe.StructField()); ok {

		if tag := field.Tag.Get("form"); tag != "" {

			return strings.Split(tag, ",")[0]
		}
		if tag := field.Tag.Get("json"); tag != "" {
			return strings.Split(tag, ",")[0]
		}
	}

	// fallback
	return fe.Field()
}

func buildMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email", fe.Field())
	}
	return fmt.Sprintf("%s is not valid", fe.Field())
}

func SendErrorResponse(c *gin.Context, err error) {
	var domainErr DomainError
	var dataErr DataError
	var powerError PowerError

	switch {
	case errors.As(err, &domainErr):
		c.JSON(http.StatusBadRequest, gin.H{"error": domainErr.Message})
	case errors.As(err, &dataErr):
		c.JSON(http.StatusBadRequest, gin.H{"error": dataErr.Message})
	case errors.As(err, &powerError):
		c.JSON(http.StatusForbidden, gin.H{"error": powerError.Message})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred: " + err.Error()})
	}
}

func SendValidationErrorResponse(c *gin.Context, errors []ValidationErrorItem) {
	c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
}

// ///// SOLUCIÓN CRISTIAN ///////

type customFieldError struct {
	field  string
	tag    string
	errMsg string
}

// Translate implements validator.FieldError.
func (c customFieldError) Translate(ut ut.Translator) string {
	return c.errMsg
}

func (c customFieldError) Tag() string             { return c.tag }
func (c customFieldError) ActualTag() string       { return c.tag }
func (c customFieldError) Namespace() string       { return c.field }
func (c customFieldError) StructNamespace() string { return c.field }
func (c customFieldError) Field() string           { return c.field }
func (c customFieldError) StructField() string     { return c.field }
func (c customFieldError) Param() string           { return "" }
func (c customFieldError) Value() interface{}      { return nil }
func (c customFieldError) Kind() reflect.Kind      { return reflect.String }
func (c customFieldError) Type() reflect.Type      { return reflect.TypeOf("") }
func (c customFieldError) Error() string           { return c.errMsg }

func BindJSON(ctx *gin.Context, obj any) error {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		return _bind(obj, err, "json")
	}

	return nil
}

func BindQuery(ctx *gin.Context, obj any) error {
	if err := ctx.ShouldBindQuery(obj); err != nil {
		return _bind(obj, err, "query")
	}

	return nil
}

func Bind(ctx *gin.Context, obj any) error {
	if err := ctx.ShouldBind(obj); err != nil {
		return _bind(obj, err, "form")
	}

	return nil
}

func _bind(obj any, err error, tagName string) error {

	ve := make(validator.ValidationErrors, 0)

	if err.Error() == "EOF" {
		val := reflect.ValueOf(obj)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		t := val.Type()

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fe := makeFieldError(field, tagName)
			ve = append(ve, fe)
		}

		return ve
	}

	if veOrig, ok := err.(validator.ValidationErrors); ok {
		val := reflect.ValueOf(obj)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		t := val.Type()

		ve = make(validator.ValidationErrors, 0)

		for _, fe := range veOrig {
			//msg := ""
			fieldName := fe.Field()
			msg := fieldName + " inválido"
			if f, ok := t.FieldByName(fe.StructField()); ok {
				fieldName = getFieldName(f, tagName)
				msg = fieldName + " inválido" // fallback
				if tag := f.Tag.Get("errMsg"); tag != "" {
					msg = tag
				}
			}

			customFE := customFieldError{
				field:  fieldName,
				tag:    fe.Tag(),
				errMsg: msg,
			}
			ve = append(ve, customFE)
		}

		return ve
	}

	// Otro tipo de error
	err = errors.New("Error no reconocido: no se pudo procesar la solicitud")

	return err
}

func makeFieldError(field reflect.StructField, tag string) validator.FieldError {
	fieldName := getFieldName(field, tag)
	msg := fieldName + " es obligatorio" // fallback
	if _tag := field.Tag.Get("errMsg"); _tag != "" {
		msg = _tag
	}
	return &customFieldError{
		field:  fieldName,
		tag:    "required",
		errMsg: msg,
	}
}

func getFieldName(field reflect.StructField, tag string) string {
	if _tag := field.Tag.Get(tag); _tag != "" {
		return _tag
	}
	return field.Name
}
