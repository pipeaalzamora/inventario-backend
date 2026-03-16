package externalservices

import (
	"bytes"
	"embed"
	"html/template"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

//go:embed templates/*.html
var templatesFS embed.FS

type RenderService struct {
}

func NewRendererService() ports.PortRender {
	return &RenderService{}
}

func (r *RenderService) Render(filePath string, data any) (string, error) {
	tmpl, err := template.ParseFS(templatesFS, filePath)
	if err != nil {
		return "", types.ThrowData("Error al procesar la plantilla")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", types.ThrowData("Error al ejecutar la plantilla")
	}

	return buf.String(), nil
}
