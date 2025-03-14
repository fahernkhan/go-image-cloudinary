package main

import (
	"ecommerce-app/internal/infrastructure/cloudinary"
	"ecommerce-app/internal/infrastructure/gin"
	"ecommerce-app/internal/interfaces/http"
	"ecommerce-app/internal/usecases/image"
	"ecommerce-app/pkg/logger"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Inisialisasi logger
	logger.InitLogger()
	logger.Logger.Info("Logger initialized")

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("Error loading .env file", zap.Error(err))
	}

	// Setup Cloudinary
	cld, err := cloudinary.NewCloudinaryClient()
	if err != nil {
		logger.Logger.Fatal("Failed to initialize Cloudinary", zap.Error(err))
	}

	// Buat adapter untuk Cloudinary
	cldAdapter := cloudinary.NewCloudinaryAdapter(cld)

	// Setup use case
	uploadImageUseCase := image.NewUploadImageUseCase(cldAdapter)

	// Setup handler
	imageHandler := http.NewImageHandler(uploadImageUseCase)

	// Setup router
	r := gin.SetupRouter(imageHandler)

	// Run server
	logger.Logger.Info("Server started on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
