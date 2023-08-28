package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Title       string  `json:"title" gorm:"not null"`
	Developer   string  `json:"developer" gorm:"not null"`
	Sinopsis    string  `json:"sinopsis" gorm:"not null"`
	ReleaseDate string  `json:"release_date" gorm:"not null"`
	Rating      float32 `json:"rating" gorm:"default:0"`

	UsersStatus []*UserReviewable `json:"user_interactions" gorm:"polymorphic:Reviewable;"`

	Genres  []*Genre  `json:"genres" gorm:"many2many:game_genres;"`
	Reviews []*Review `json:"reviews" gorm:"polymorphic:Reviewable;"`
}

func (g *Game) GetID() uint {
	return g.ID
}

func (g *Game) GetType() string {
	return "games"
}

func (g *Game) GetReviews() []*Review {
	return g.Reviews
}

func (g *Game) AppendReview(review *Review) {
	g.Reviews = append(g.Reviews, review)
}

func (g *Game) GetRating() float32 {
	return g.Rating
}

func (s *Game) UpdateRating() float32 {
	var rating float32
	for _, review := range s.Reviews {
		rating += review.Rating
	}
	s.Rating = rating / float32(len(s.Reviews))
	return s.Rating
}
