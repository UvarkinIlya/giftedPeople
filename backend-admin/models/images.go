package models

import "mime/multipart"

type ImgWithID struct {
	ID   string
	File multipart.File
}
