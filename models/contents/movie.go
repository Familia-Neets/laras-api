package contents

import (
	"Lara/models/reviewable"
	"Lara/models/users"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string  `json:"title" gorm:"not null"`
	Sinopsis    string  `json:"sinopsis" gorm:"not null"`
	Director    string  `json:"director" gorm:"not null"`
	ReleaseDate string  `json:"release_date" gorm:"not null"`
	Rating      float32 `json:"rating" gorm:"not null;default:0"`

	UsersStatus []*users.UserReviewable `json:"user_interactions" gorm:"polymorphic:Reviewable;"`

	Genres  []*Genre             `json:"genres" gorm:"many2many:movie_genres;"`
	Reviews []*reviewable.Review `json:"reviews" gorm:"polymorphic:Reviewable;"`
}

func (m *Movie) GetID() uint {
	return m.ID
}

func (m *Movie) GetType() string {
	return "movies"
}

func (m *Movie) GetReviews() []*reviewable.Review {
	return m.Reviews
}

func (m *Movie) AppendReview(review *reviewable.Review) {
	m.Reviews = append(m.Reviews, review)
}

func (m *Movie) GetRating() float32 {
	return m.Rating
}

func (s *Movie) UpdateRating() float32 {
	var rating float32
	for _, review := range s.Reviews {
		rating += review.Rating
	}
	s.Rating = rating / float32(len(s.Reviews))
	return s.Rating
}
