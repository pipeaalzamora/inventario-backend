package ports

type PortRender interface {
	Render(filePath string, data any) (string, error)
}
