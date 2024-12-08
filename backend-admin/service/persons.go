package service

import "backend-admin/models"

type storage interface {
	GetAllPersons() ([]models.Person, error)
	Create(person models.Person) error
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
	return s.storage.GetAllPersons()
}

func (s *PersonService) Create(person models.Person) error {
	return s.storage.Create(person)
}

func (s *PersonService) Update(person models.Person) error {
	err := s.storage.Delete(person.ID)
	if err != nil {
		return err
	}

	return s.storage.Create(person)
}

func (s *PersonService) Delete(personId string) error {
	return s.storage.Delete(personId)
}
