// cloudinary_services.go

package services

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

// CloudinaryService struct represents the Cloudinary service configuration
type CloudinaryService struct {
	CloudinaryURL string
	CloudName     string
	APIKey        string
	APISecret     string
}

// NewCloudinaryService initializes a new Cloudinary service instance
func NewCloudinaryService(cloudinaryURL, cloudName, apiKey, apiSecret string) *CloudinaryService {
	return &CloudinaryService{
		CloudinaryURL: cloudinaryURL,
		CloudName:     cloudName,
		APIKey:        apiKey,
		APISecret:     apiSecret,
	}
}

// UploadImage uploads an image file to Cloudinary
func (cs *CloudinaryService) UploadImage(filePath string) (string, error) {
	// Initialize Cloudinary uploader
	cld, err := cloudinary.NewFromURL(cs.CloudinaryURL)
	if err != nil {
		return "", fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create context
	ctx := context.Background()

	// Upload the file to Cloudinary
	uploadResult, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{Folder: "foods"})
	if err != nil {
		return "", fmt.Errorf("failed to upload image to Cloudinary: %v", err)
	}

	// Return the URL of the uploaded image
	return uploadResult.URL, nil
}
