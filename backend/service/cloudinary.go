package service

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)


type Uploader struct {
	cld *cloudinary.Cloudinary
}


func NewUploader() (*Uploader, error) {
    	
    cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
    apiKey := os.Getenv("CLOUDINARY_API_KEY")
    apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

    if cloudName == "" || apiKey == "" || apiSecret == "" {
        return nil, fmt.Errorf("cloudinary environment variables are not set")
    }

    cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
    if err != nil {
        return nil, fmt.Errorf("failed to create cloudinary client: %w", err)
    }

	return &Uploader{cld: cld}, nil
}


func (u *Uploader) UploadBase64Image(ctx context.Context, base64Image string) (string, error) {
    
    parts := strings.Split(base64Image, ",")
    if len(parts) != 2 {
        return "", fmt.Errorf("invalid base64 format")
    }
    
    
    uploadData := base64Image
    
    
	uploadParams := uploader.UploadParams{
		
		PublicID: fmt.Sprintf("casestudy/%d", time.Now().UnixNano()),
		Folder: "depixen-casestudy", 
	}

    
	resp, err := u.cld.Upload.Upload(ctx, uploadData, uploadParams)
	if err != nil {
		return "", fmt.Errorf("cloudinary upload failed: %w", err)
	}

	return resp.SecureURL, nil
}