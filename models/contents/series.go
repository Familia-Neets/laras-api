package contents

import (
	"Lara/models/reviewable"
	"Lara/models/users"

	"gorm.io/gorm"
)

type Series struct {
	gorm.Model
	Title       string  `json:"title" gorm:"not null"`
	Sinopsis    string  `json:"sinopsis"`
	Seasons     uint    `json:"seasons"`
	Episodes    uint    `json:"episodes"`
	ReleaseDate string  `json:"release_date"`
	Rating      float32 `json:"rating" gorm:"default:0"`

	UsersStatus []*users.UserReviewable `json:"user_interactions" gorm:"polymorphic:Reviewable;"`

	Genres  []*Genre             `json:"genres" gorm:"many2many:series_genres;"`
	Reviews []*reviewable.Review `json:"reviews" gorm:"polymorphic:Reviewable;"`
}

func (s *Series) GetModel() interface{} {
	return s
}

func (s *Series) GetType() string {
	return "series"
}

func (s *Series) GetID() uint {
	return s.ID
}

func (s *Series) GetReviews() []*reviewable.Review {
	return s.Reviews
}

func (s *Series) AppendReview(review *reviewable.Review) {
	s.Reviews = append(s.Reviews, review)
}

func (s *Series) GetRating() float32 {
	return s.Rating
}

func (s *Series) UpdateRating() float32 {
	var rating float32
	for _, review := range s.Reviews {
		rating += review.Rating
	}
	s.Rating = rating / float32(len(s.Reviews))
	return s.Rating
}
