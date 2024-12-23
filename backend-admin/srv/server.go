package srv

import (
	"bytes"
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
	Update(person models.Person) error
	Delete(personId string) error
	GetPersonByRoad(roadID string) (models.Person, error)
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
	e.Use(middleware.BodyLimit("10M"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Разрешить все источники (или замените на нужные)
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	}))
	e.GET("/person", s.getAllPersons)
	e.GET("/person/by-road-id/:roadID", s.getPersonByRoad)
	e.POST("/person", s.createPerson)
	e.PUT("/person/:id", s.updatePerson)
	e.DELETE("/person/:id", s.deletePerson)

	e.GET("/person-image/:id", s.getPersonImg)

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
	_, person, err := getImgAndPerson(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// err = s.imagesService.SaveImg(img.ID, img.File)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	err = s.personsService.Create(person)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) updatePerson(c echo.Context) error {
	personId := c.Param("id")
	// img, err := getImg(c)
	// if err == nil {
	// 	err = s.imagesService.SaveImg(img.ID, img.File)
	// 	if err != nil {
	// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}
	// }

	person := models.Person{
		ID:          uuid.NewString(),
		Name:        c.FormValue("name"),
		Road:        c.FormValue("road"),
		Description: c.FormValue("description"),
		// Img:         img.ID,
		Img: "",
	}

	person.ID = personId
	err := s.personsService.Update(person)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) deletePerson(c echo.Context) error {
	id := c.Param("id")
	err := s.personsService.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) getPersonByRoad(c echo.Context) error {
	roadId := c.Param("roadID")
	person, err := s.personsService.GetPersonByRoad(roadId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, person)
}

func (s *Server) getPersonImg(c echo.Context) error {
	id := c.Param("id")
	img, err := s.imagesService.GetImg(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(img)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(http.StatusOK, "image/jpg", buf.Bytes())
}

func getImgAndPerson(c echo.Context) (img models.ImgWithID, person models.Person, err error) {
	// img, err = getImg(c)
	// if err != nil {
	// 	return models.ImgWithID{}, models.Person{}, err
	// }

	person = models.Person{
		ID:          uuid.NewString(),
		Name:        c.FormValue("name"),
		Road:        c.FormValue("road"),
		Description: c.FormValue("description"),
		// Img:         img.ID,
		Img: "",
	}

	return img, person, nil
}

func getImg(c echo.Context) (models.ImgWithID, error) {
	imgID := uuid.NewString()

	file, err := c.FormFile("file[]")
	if err != nil {
		return models.ImgWithID{}, err
	}
	src, err := file.Open()
	if err != nil {
		return models.ImgWithID{}, err
	}
	defer src.Close()

	return models.ImgWithID{
		ID:   imgID,
		File: src,
	}, nil
}
