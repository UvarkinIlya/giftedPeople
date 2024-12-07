package storage

import "backend-admin/models"

type Storage struct{}

func New() (*Storage, error) {
	return &Storage{}, nil
}

func (s *Storage) GetAllPersons() ([]models.Person, error) {
	panic("implement me")
	return []models.Person{}, nil
}

func (s *Storage) Save(person models.Person) error {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) Delete(personId string) error {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) GetImg(imageId string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) SaveImg(imageId string, image []byte) error {
	// TODO implement me
	panic("implement me")
}
