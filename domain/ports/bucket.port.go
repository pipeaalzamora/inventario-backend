package ports

import (
	"context"
	"mime/multipart"
)

type PortBucket interface {
	UploadFile(ctx context.Context, file multipart.File, fileName string) (string, error)
	DeleteFile(ctx context.Context, fileURL string) error
}
