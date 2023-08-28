package helpers

import (
	"Lara/models"
	"errors"
	"fmt"
)

func GetContentInstance(contentType string) (models.Reviewable, error) {
	switch contentType {
	case "books":
		return &models.Book{}, nil
	case "movies":
		return &models.Movie{}, nil
	case "games":
		return &models.Game{}, nil
	case "series":
		return &models.Series{}, nil
	default:
		return nil, fmt.Errorf("Invalid content type")
	}
}

func ValidateRequiredFields(content models.Reviewable) error {
	switch c := content.(type) {
	case *models.Book:
		if c.Title == "" || c.Sinopsis == "" || c.Author == "" || c.ReleaseDate == "" || c.ISBN == "" {
			return errors.New("Campos obrigatórios para o livro não estão preenchidos")
		}
	case *models.Movie:
		if c.Title == "" || c.Sinopsis == "" || c.Director == "" || c.ReleaseDate == "" {
			return errors.New("Campos obrigatórios para o filme não estão preenchidos")
		}
	case *models.Game:
		if c.Title == "" || c.Sinopsis == "" || c.Developer == "" || c.ReleaseDate == "" {
			return errors.New("Campos obrigatórios para o jogo não estão preenchidos")
		}
	default:
		return errors.New("Tipo de conteúdo inválido")
	}

	return nil
}
