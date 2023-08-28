package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string  `json:"title" gorm:"not null"`
	Author      string  `json:"author" gorm:"not null"`
	Sinopsis    string  `json:"sinopsis" gorm:"not null"`
	ReleaseDate string  `json:"release_date" gorm:"not null"`
	ISBN        string  `json:"isbn" gorm:"not null"`
	Rating      float32 `json:"rating" gorm:"default:0"`

	UsersStatus []*UserReviewable `json:"user_interactions" gorm:"polymorphic:Reviewable;"`

	Genres  []*Genre  `json:"genres" gorm:"many2many:book_genres;"`
	Reviews []*Review `json:"reviews" gorm:"polymorphic:Reviewable;"`
}

func (b *Book) GetID() uint {
	return b.ID
}

func (b *Book) GetType() string {
	return "books"
}

func (b *Book) GetReviews() []*Review {
	return b.Reviews
}

func (b *Book) AppendReview(review *Review) {
	b.Reviews = append(b.Reviews, review)
}

func (b *Book) GetRating() float32 {
	return b.Rating
}

func (s *Book) UpdateRating() float32 {
	var rating float32
	for _, review := range s.Reviews {
		rating += review.Rating
	}
	s.Rating = rating / float32(len(s.Reviews))
	return s.Rating
}
