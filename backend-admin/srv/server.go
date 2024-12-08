package srv

import (
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"backend-admin/models"
)

type personService interface {
	GetAllPersons() ([]models.Person, error)
	Create(person models.Person) error
	Delete(personId string) error
}

type imagesService interface {
	GetImg(imageId string) (io.Reader, error)
	SaveImg(imageId string, image io.Reader) error
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/person", s.getAllPersons)
	e.GET("/person/:roadID", s.getPersonByRoad)
	e.POST("/person", s.createPerson)
	e.PUT("/person", s.updatePerson)
	e.DELETE("/person", s.deletePerson)

	// TODO send file

	e.Logger.Fatal(e.Start(":" + s.port))
}

func (s *Server) getAllPersons(c echo.Context) error {
	persons, err := s.personsService.GetAllPersons()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, persons)
}

func (s *Server) createPerson(c echo.Context) error {
	img, person, err := getImgAndPerson(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = s.imagesService.SaveImg(img.ID, img.File)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = s.personsService.Create(person)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) updatePerson(c echo.Context) error {
	img, person, err := getImgAndPerson(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = s.imagesService.SaveImg(img.ID, img.File)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = s.personsService.Create(person)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) deletePerson(c echo.Context) error {
	panic("implement me")
	return nil
}

func (s *Server) getPersonByRoad(c echo.Context) error {
	panic("implement me")
	return nil
}

func getImgAndPerson(c echo.Context) (img models.ImgWithID, person models.Person, err error) {
	imgID := uuid.NewString()

	file, err := c.FormFile("file")
	if err != nil {
		return models.ImgWithID{}, models.Person{}, err
	}
	src, err := file.Open()
	if err != nil {
		return models.ImgWithID{}, models.Person{}, err
	}
	defer src.Close()

	img = models.ImgWithID{
		ID:   imgID,
		File: src,
	}

	person = models.Person{
		ID:          uuid.NewString(),
		Name:        c.FormValue("name"),
		Road:        c.FormValue("roadID"),
		Description: c.FormValue("description"),
		Img:         imgID,
	}

	return img, person, nil
}
