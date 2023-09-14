package helpers

import (
	"Lara/models/contents"
	"Lara/models/reviewable"
	"errors"
	"fmt"
)

func GetContentInstance(contentType string) (reviewable.Reviewable, error) {
	switch contentType {
	case "books":
		return &contents.Book{}, nil
	case "movies":
		return &contents.Movie{}, nil
	case "games":
		return &contents.Game{}, nil
	case "series":
		return &contents.Series{}, nil
	default:
		return nil, fmt.Errorf("Invalid content type")
	}
}

func ValidateRequiredFields(content reviewable.Reviewable) error {
	switch c := content.(type) {
	case *contents.Book:
		if c.Title == "" || c.Sinopsis == "" || c.Author == "" || c.ReleaseDate == "" || c.ISBN == "" {
			return errors.New("Campos obrigatórios para o livro não estão preenchidos")
		}
	case *contents.Movie:
		if c.Title == "" || c.Sinopsis == "" || c.Director == "" || c.ReleaseDate == "" {
			return errors.New("Campos obrigatórios para o filme não estão preenchidos")
		}
	case *contents.Game:
		if c.Title == "" || c.Sinopsis == "" || c.Developer == "" || c.ReleaseDate == "" {
			return errors.New("Campos obrigatórios para o jogo não estão preenchidos")
		}
	default:
		return errors.New("Tipo de conteúdo inválido")
	}

	return nil
}
