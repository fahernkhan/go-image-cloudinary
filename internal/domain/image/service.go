package image

import "io"

type ImageService interface {
	UploadImage(file io.Reader, filename string) (*Image, error)
}
