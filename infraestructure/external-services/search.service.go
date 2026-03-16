package externalservices

import (
	"sofia-backend/config"
	"sofia-backend/domain/ports"
	"sofia-backend/types"

	"github.com/meilisearch/meilisearch-go"
)

type SearchService struct {
	melisearchClient meilisearch.ServiceManager
}

func NewSearchService(cfg *config.Config) ports.PortSearch {
	client := meilisearch.New(cfg.Search.Host, meilisearch.WithAPIKey(cfg.Search.ApiKey))

	return &SearchService{melisearchClient: client}
}
func (s *SearchService) Search(index string, query string) ([]map[string]interface{}, error) {
	searchResult, err := s.melisearchClient.Index(index).Search(query, &meilisearch.SearchRequest{})

	if err != nil {
		return nil, types.ThrowData("Error al buscar en el índice")
	}

	results := make([]map[string]interface{}, len(searchResult.Hits))
	for i, hit := range searchResult.Hits {
		results[i] = hit.(map[string]interface{})
	}

	return results, nil
}

func (s *SearchService) AddDocument(index string, document map[string]interface{}) error {
	_, err := s.melisearchClient.Index(index).AddDocuments([]map[string]interface{}{document})
	if err != nil {
		return types.ThrowData("Error al agregar el documento al índice")
	}
	return nil
}

func (s *SearchService) DeleteDocument(index string, documentID string) error {
	_, err := s.melisearchClient.Index(index).DeleteDocument(documentID)
	if err != nil {
		return err
	}
	return nil
}

func (s *SearchService) UpdateDocument(index string, document map[string]interface{}) error {
	_, err := s.melisearchClient.Index(index).UpdateDocuments([]map[string]interface{}{document})
	if err != nil {
		return types.ThrowData("Error al actualizar el documento en el índice")
	}
	return nil
}
