package image

import (
	"ecommerce-app/internal/domain/image"
	"ecommerce-app/pkg/utils"
	"io"

	"github.com/google/uuid"
)

type UploadImageUseCase struct {
	ImageService image.ImageService
}

func NewUploadImageUseCase(imageService image.ImageService) *UploadImageUseCase {
	return &UploadImageUseCase{ImageService: imageService}
}

func (uc *UploadImageUseCase) Execute(file io.Reader, filename string) (*image.Image, error) {
	url, publicID, err := utils.UploadToCloudinary(file, filename)
	if err != nil {
		return nil, err
	}

	// Buat instance Image dengan ID yang dihasilkan
	return &image.Image{
		ID:       uuid.New().String(), // Generate UUID
		URL:      url,
		PublicID: publicID,
	}, nil
}
