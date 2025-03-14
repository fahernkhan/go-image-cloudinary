package utils

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file io.Reader, filename string) (string, string, error) {
	ctx := context.Background()
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", "", err
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: filename,
	})
	if err != nil {
		return "", "", err
	}

	return uploadResult.SecureURL, uploadResult.PublicID, nil
}
