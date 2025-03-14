package http

import (
	entity "ecommerce-app/internal/domain/image"
	"ecommerce-app/internal/usecases/image"
	"ecommerce-app/pkg/logger"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImageHandler struct {
	UploadImageUseCase *image.UploadImageUseCase
}

func NewImageHandler(uploadImageUseCase *image.UploadImageUseCase) *ImageHandler {
	return &ImageHandler{UploadImageUseCase: uploadImageUseCase}
}

// removeExtension menghapus ekstensi dari nama file
func removeExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func (h *ImageHandler) UploadImage(c *gin.Context) {
	logger.Logger.Info("Handling image upload request")

	// Ambil file dari request
	file, err := c.FormFile("file")
	if err != nil {
		logger.Logger.Error("Failed to get file from request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buka file
	fileContent, err := file.Open()
	if err != nil {
		logger.Logger.Error("Failed to open file", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer fileContent.Close()

	// Hapus ekstensi dari nama file
	filenameWithoutExt := removeExtension(file.Filename)

	// Upload gambar ke Cloudinary
	uploadedImage, err := h.UploadImageUseCase.Execute(fileContent, filenameWithoutExt)
	if err != nil {
		logger.Logger.Error("Failed to upload image to Cloudinary", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Buat instance Image dengan ID yang dihasilkan
	finalImage := &entity.Image{
		ID:       uploadedImage.ID, // Pastikan ID diisi di use case atau domain
		URL:      uploadedImage.URL,
		PublicID: uploadedImage.PublicID,
	}

	// Kirim respons ke client
	logger.Logger.Info("Image uploaded successfully",
		zap.String("url", finalImage.URL),
		zap.String("public_id", finalImage.PublicID),
	)
	c.JSON(http.StatusOK, finalImage)
}
