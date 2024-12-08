package service

import "mime/multipart"

type ImageStorage interface {
	GetImg(imageId string) ([]byte, error)
	SaveImg(imageId string, image multipart.File) error
}

type ImageService struct {
	storage ImageStorage
}

func NewImageService(storage ImageStorage) *ImageService {
	return &ImageService{
		storage: storage,
	}
}

func (s *ImageService) GetImg(imageId string) ([]byte, error) {
	panic("implement me")
}

func (s *ImageService) SaveImg(imageId string, image multipart.File) error {
	panic("implement me")
}
