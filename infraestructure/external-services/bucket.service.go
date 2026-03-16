package externalservices

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"sofia-backend/config"
	"sofia-backend/domain/ports"
	"sofia-backend/shared"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type BucketService struct {
	BucketName  string
	UploadURL   string
	UploadPath  string
	BaseFolder  string
	AuthURL     string
	ClientEmail string
	PrivateKey  string
}

func NewBucketService(config *config.Config) ports.PortBucket {
	return &BucketService{
		BucketName:  config.Bucket.GCPBucketName,
		UploadPath:  config.Bucket.GCPUploadPath,
		BaseFolder:  config.Bucket.GCPBaseFolder,
		AuthURL:     config.Bucket.GCPAuthURL,
		ClientEmail: config.Bucket.GCPClientEmail,
		PrivateKey:  config.Bucket.GCPPrivateKey,
	}
}

// async save(content: Buffer, name: string, relativePath?: string): Promise<FileUrl> {
//     const token = await this.#getAccessToken();

//     const fileName = relativePath ? `${relativePath}/${name}` : name;
//     const url = `https://storage.googleapis.com/upload/storage/v1/b/${this.bucketName}/o?uploadType=media&name=${fileName}`;

//     const mimeType = this.#getMimeTypeFromExtension(path.extname(name));

//     await fetch(url, {
//         method: "POST",
//         headers: {
//             Authorization: `Bearer ${token}`,
//             "Content-Type": mimeType,
//         },
//         body: content,
//     });

//     const fileUrl = {
//         uri: `${this.uploadUrl}${fileName}`,
//         relativeUri: `${this.uploadUrl}${relativePath || ''}`,
//         name: name,
//         path: `${this.uploadPath}/${fileName}`
//     };

//     return fileUrl;
// }

func (b *BucketService) UploadFile(ctx context.Context, file multipart.File, name string) (string, error) {
	accessToken, err := b.getAccessToken()
	if err != nil {
		return "", err
	}

	now := time.Now().Format("20060102150405")
	ext := path.Ext(name)
	newName := strings.ReplaceAll(shared.GenerateUUID(), "-", "") + "_" + now

	fileName := b.BaseFolder + "/" + newName + ext
	uploadStorageUrl := fmt.Sprintf("https://storage.googleapis.com/upload/storage/v1/b/%s/o?uploadType=media&name=%s", b.BucketName, fileName)
	mimeType := shared.GetMimeType(ext)

	// Read file content into buffer
	fileBuffer := &bytes.Buffer{}
	_, err = io.Copy(fileBuffer, file)
	if err != nil {
		return "", err
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uploadStorageUrl, fileBuffer)
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", mimeType)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("failed to upload file: received status code %d", resp.StatusCode)
	}

	// Build and return file URL
	fileURL := fmt.Sprintf("%s/%s", b.UploadPath, fileName)

	return fileURL, nil
}

func (b *BucketService) DeleteFile(ctx context.Context, fileURL string) error {
	// Implement file deletion logic here
	token, err := b.getAccessToken()
	if err != nil {
		return err
	}

	urlTo := fmt.Sprintf("%s/", b.UploadPath)
	fileName := strings.TrimPrefix(fileURL, urlTo)
	encodedFileName := url.PathEscape(fileName)

	url := fmt.Sprintf("https://storage.googleapis.com/storage/v1/b/%s/o/%s", b.BucketName, encodedFileName)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Delete response status: %+v\n", resp) // Log the response status for debugging

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete file: received status code %d", resp.StatusCode)
	}

	return nil
}

func (b *BucketService) generateJWT() (string, error) {
	now := time.Now()
	exp := now.Add(time.Minute * 60)
	claims := jwt.MapClaims{
		"iss":   b.ClientEmail,
		"scope": "https://www.googleapis.com/auth/devstorage.full_control",
		"aud":   b.AuthURL,
		"exp":   exp.Unix(),
		"iat":   now.Unix(),
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(b.PrivateKey))
	if err != nil {
		return "", err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (b *BucketService) getAccessToken() (string, error) {
	token, err := b.generateJWT()
	if err != nil {
		return "", err
	}

	data := url.Values{}
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	data.Set("assertion", token)
	response, err := http.PostForm(b.AuthURL, data)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var result map[string]any
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if _, ok := result["access_token"]; !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	return result["access_token"].(string), nil
}
