package storage

import (
	"bytes"
	"context"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend-admin/models"
)

type Storage struct {
	personsCollection *mongo.Collection
	imagesCollection  *gridfs.Bucket
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

	db := client.Database("persons")

	personsCollection := db.Collection("persons")
	imagesCollection, err := gridfs.NewBucket(db, options.GridFSBucket().SetName("images"))
	if err != nil {
		return nil, err
	}

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

func (s *Storage) Update(person models.Person) error {
	_, err := s.personsCollection.ReplaceOne(context.Background(), bson.D{{"id", person.ID}}, person)
	return err
}

func (s *Storage) Delete(personId string) error {
	_, err := s.personsCollection.DeleteOne(context.Background(), bson.D{{"id", personId}})
	return err
}

func (s *Storage) GetPersonsByRoad(roadId string) ([]models.Person, error) {
	cursor, err := s.personsCollection.Find(context.Background(), bson.D{{"road", roadId}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	persons := make([]models.Person, 0)
	err = cursor.All(context.Background(), &persons)

	return persons, nil
}

func (s *Storage) GetImg(imageId string) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	_, err := s.imagesCollection.DownloadToStreamByName(imageId, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *Storage) SaveImg(imageId string, image io.Reader) error {
	_, err := s.imagesCollection.UploadFromStream(imageId, image)
	return err
}
