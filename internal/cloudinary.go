package internal

import (
	"DiskusiTugas/config"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func SetupCloudinary(secret_api string, cloud_name string, api_key string) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(cloud_name, api_key, secret_api)
	if err != nil {
		return nil, err
	}

	return cld, nil
}

func UploadToCloudinary(input interface{}) (string, error) {
	config := config.InitConfig()
	ctx := context.Background()
	cld, err := SetupCloudinary(config.CloudinaryAPISecret, config.CloudinaryCloudName, config.CloudinaryAPIKey)
	if err != nil {
		return "", err
	}
	result, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.CloudinaryUploadFolder})
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
