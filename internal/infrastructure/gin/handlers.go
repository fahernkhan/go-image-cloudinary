package gin

import (
	"ecommerce-app/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(imageHandler *http.ImageHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/upload", imageHandler.UploadImage)

	return r
}
