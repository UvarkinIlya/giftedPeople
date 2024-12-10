package service

import (
	"math/rand"

	"backend-admin/models"
)

type storage interface {
	Get(id string) (models.Person, error)
	GetAllPersons() ([]models.Person, error)
	Create(person models.Person) error
	Update(person models.Person) error
	Delete(personId string) error
	GetPersonsByRoad(roadId string) ([]models.Person, error)
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
	oldPerson, err := s.storage.Get(person.ID)
	if err != nil {
		return err
	}

	if person.Img == "" {
		person.Img = oldPerson.Img
	}

	err = s.storage.Update(person)
	if err != nil {
		return err
	}

	return nil
}

func (s *PersonService) Delete(personId string) error {
	return s.storage.Delete(personId)
}

func (s *PersonService) GetPersonByRoad(roadID string) (models.Person, error) {
	persons, err := s.storage.GetPersonsByRoad(roadID)
	if err != nil {
		return models.Person{}, err
	}

	return persons[rand.Intn(len(persons))], nil
}
