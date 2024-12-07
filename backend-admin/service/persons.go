package service

import "backend-admin/models"

type storage interface {
	GetAllPersons() ([]models.Person, error)
	Save(person models.Person) error
	Delete(personId string) error
}

type PersonService struct {
	storage storage
}

func NewPersonService(storage storage) *PersonService {
	return &PersonService{
		storage: storage,
	}
}

func (s *PersonService) GetAllPersons() ([]models.Person, error) {
	panic("implement me")
}

func (s *PersonService) Save(person models.Person) error {
	panic("implement me")
}

func (s *PersonService) Delete(personId string) error {
	panic("implement me")
}
