package service

import (
	"io"
)

type ImageStorage interface {
	GetImg(imageId string) (io.Reader, error)
	SaveImg(imageId string, image io.Reader) error
}

type ImageService struct {
	storage ImageStorage
}

func NewImageService(storage ImageStorage) *ImageService {
	return &ImageService{
		storage: storage,
	}
}

func (s *ImageService) GetImg(imageId string) (io.Reader, error) {
	return s.storage.GetImg(imageId)
}

func (s *ImageService) SaveImg(imageId string, image io.Reader) error {
	return s.storage.SaveImg(imageId, image)
}
