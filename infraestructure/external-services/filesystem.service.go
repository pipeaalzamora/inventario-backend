package externalservices

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sofia-backend/config"
	"sofia-backend/domain/ports"
	"strings"
)

type FileSystemService struct {
	relativePath string
}

func NewFileSystemService(config *config.Config) ports.PortBucket {
	return &FileSystemService{
		relativePath: config.FileSys.RelativePath,
	}
}

func (f *FileSystemService) UploadFile(ctx context.Context, file multipart.File, fileName string) (string, error) {
	// Create the uploads directory if it doesn't exist
	uploadDir := f.relativePath
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Create the full file path
	filePath := filepath.Join(uploadDir, fileName)

	// Create the destination file
	destFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer destFile.Close()

	// Reset file pointer to beginning
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("failed to reset file pointer: %w", err)
	}

	// Copy the uploaded file content to the destination file
	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}

	// Return the relative file path
	return filePath, nil
}

func (f *FileSystemService) DeleteFile(ctx context.Context, fileURL string) error {
	// Clean the file path and ensure it's within our upload directory
	cleanPath := filepath.Clean(fileURL)

	// Security check: ensure the file is within our upload directory
	if !strings.HasPrefix(cleanPath, f.relativePath) {
		return fmt.Errorf("file path is outside upload directory")
	}

	// Check if file exists
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", cleanPath)
	}

	// Delete the file
	err := os.Remove(cleanPath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}
