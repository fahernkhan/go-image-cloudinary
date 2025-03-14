package cloudinary

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func NewCloudinaryClient() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return nil, err
	}
	return cld, nil
}
