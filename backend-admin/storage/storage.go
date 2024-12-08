package storage

import (
	"context"
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend-admin/models"
)

type Storage struct {
	personsCollection *mongo.Collection
	imagesCollection  *mongo.Collection
}

func New() (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	personsCollection := client.Database("persons").Collection("persons")
	imagesCollection := client.Database("persons").Collection("images")

	return &Storage{
		personsCollection: personsCollection,
		imagesCollection:  imagesCollection,
	}, nil
}

func (s *Storage) GetAllPersons() ([]models.Person, error) {
	filter := bson.M{}
	cursor, err := s.personsCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	persons := make([]models.Person, 0)
	err = cursor.All(context.Background(), &persons)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (s *Storage) Create(person models.Person) error {
	_, err := s.personsCollection.InsertOne(context.Background(), person)
	return err
}

func (s *Storage) Delete(personId string) error {
	_, err := s.personsCollection.DeleteOne(context.Background(), personId)
	return err
}

func (s *Storage) GetImg(imageId string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) SaveImg(imageId string, image multipart.File) error {
	// TODO implement me
	panic("implement me")
}
