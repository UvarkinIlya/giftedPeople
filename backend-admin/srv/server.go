package srv

import (
	"github.com/labstack/echo/v4"

	"backend-admin/models"
)

type personService interface {
	GetAllPersons() ([]models.Person, error)
	Save(person models.Person) error
	Delete(personId string) error
}

type imagesService interface {
	GetImg(imageId string) ([]byte, error)
	SaveImg(imageId string, image []byte) error
}

type Server struct {
	personsService personService
	imagesService  imagesService
	port           string
}

func NewServer(personsService personService, imagesService imagesService, port string) *Server {
	return &Server{
		personsService: personsService,
		imagesService:  imagesService,
		port:           port,
	}
}

func (s *Server) Run() {
	e := echo.New()

	e.GET("/person", getAllPersons)
	e.GET("/person/:roadID", getPersonByRoad)
	e.POST("/person", createPerson)
	e.PUT("/person", updatePerson)
	e.DELETE("/person", deletePerson)

	// TODO send file

	e.Logger.Fatal(e.Start(":" + s.port))
}

func getAllPersons(c echo.Context) error {
	panic("implement me")
	return nil
}

func createPerson(c echo.Context) error {
	panic("implement me")
	return nil
}

func updatePerson(c echo.Context) error {
	panic("implement me")
	return nil
}

func deletePerson(c echo.Context) error {
	panic("implement me")
	return nil
}

func getPersonByRoad(c echo.Context) error {
	panic("implement me")
	return nil
}
