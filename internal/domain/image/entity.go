package image

import "github.com/google/uuid"

type Image struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	PublicID string `json:"public_id"`
}

// NewImage adalah constructor function untuk membuat instance Image
func NewImage(url, publicID string) *Image {
	return &Image{
		ID:       uuid.New().String(), // Generate UUID sebagai ID
		URL:      url,
		PublicID: publicID,
	}
}
