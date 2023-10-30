package internal

import (
	"DiskusiTugas/config"
	"context"
	"net/url"
	"path/filepath"
	"strings"

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

func DeleteFromCloudinary(input string) error {
	config := config.InitConfig()
	ctx := context.Background()
	cld, err := SetupCloudinary(config.CloudinaryAPISecret, config.CloudinaryCloudName, config.CloudinaryAPIKey)
	if err != nil {
		return err
	}
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: input})
	if err != nil {
		return err
	}
	return nil
}

func GetPublicIDFromURL(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	path := parsedURL.Path
	path = strings.TrimPrefix(path, "/")
	parts := strings.Split(path, "/")
	result := strings.Join(parts[len(parts)-2:], "/")
	result = strings.TrimSuffix(result, filepath.Ext(result))
	return result, nil
}
