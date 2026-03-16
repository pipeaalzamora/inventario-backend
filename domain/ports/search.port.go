package ports

type PortSearch interface {
	Search(index string, query string) ([]map[string]interface{}, error)
	AddDocument(index string, document map[string]interface{}) error
	DeleteDocument(index string, documentID string) error
	UpdateDocument(index string, document map[string]interface{}) error
}
