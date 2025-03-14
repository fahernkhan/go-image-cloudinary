package cloudinary

import (
	"context"
	"ecommerce-app/internal/domain/image"
	"ecommerce-app/pkg/logger"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.uber.org/zap"
)

// CloudinaryAdapter mengimplementasikan ImageService
type CloudinaryAdapter struct {
	cld *cloudinary.Cloudinary
}

// NewCloudinaryAdapter membuat instance baru dari CloudinaryAdapter
func NewCloudinaryAdapter(cld *cloudinary.Cloudinary) *CloudinaryAdapter {
	return &CloudinaryAdapter{cld: cld}
}

// UploadImage mengimplementasikan method UploadImage dari ImageService
func (ca *CloudinaryAdapter) UploadImage(file io.Reader, filename string) (*image.Image, error) {
	ctx := context.Background()

	logger.Logger.Info("Uploading image to Cloudinary", zap.String("filename", filename))

	uploadResult, err := ca.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: filename,
	})
	if err != nil {
		logger.Logger.Error("Failed to upload image to Cloudinary", zap.Error(err))
		return nil, err
	}

	logger.Logger.Info("Image uploaded to Cloudinary", zap.String("url", uploadResult.SecureURL))
	return &image.Image{
		URL:      uploadResult.SecureURL,
		PublicID: uploadResult.PublicID,
	}, nil
}
