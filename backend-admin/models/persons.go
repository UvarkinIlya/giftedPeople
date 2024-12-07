package models

type Person struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Road        string `json:"road"`
	Description string `json:"desc"`
	Img         string `json:"img"`
}
