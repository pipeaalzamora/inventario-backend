package shared

type PaginationMetadata struct {
	Page            int  `json:"page"`
	Size            int  `json:"size"`
	Total           int  `json:"total"`
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
}

type PaginationResponse[T any] struct {
	Items    []T                `json:"items"`
	Metadata PaginationMetadata `json:"metadata"`
}

func NewPagination[T any](items []T, total, page, size int) PaginationResponse[T] {
	hasNextPage := (page * size) < total
	hasPreviousPage := page > 1

	return PaginationResponse[T]{
		Items: items,
		Metadata: PaginationMetadata{
			Page:            page,
			Size:            size,
			Total:           total,
			HasNextPage:     hasNextPage,
			HasPreviousPage: hasPreviousPage,
		},
	}
}

type PageQueryParams struct {
	Page   int                     `form:"page" binding:"omitempty,min=1"`
	Size   int                     `form:"size" binding:"omitempty,min=1,max=30"`
	Filter *map[string]interface{} `form:"filter" binding:"omitempty"`
}
