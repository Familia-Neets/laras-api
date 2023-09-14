package users

import (
	"Lara/models/reviewable"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"unique; not null"`
	Email    string `json:"email" gorm:"unique; not null"`
	Password string `json:"password" gorm:"not null"`

	Reviews  []*reviewable.Review `json:"reviews"`
	PlanTo   []*UserReviewable    `json:"plan_to" gorm:"foreignKey:UserID"`
	Current  []*UserReviewable    `json:"current" gorm:"foreignKey:UserID"`
	Finished []*UserReviewable    `json:"finished" gorm:"foreignKey:UserID"`

	Likes []*reviewable.Like `json:"likes" gorm:"foreignKey:UserID"`
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetReviews() []*reviewable.Review {
	return u.Reviews
}

func (u *User) GetStats() map[string]int {
	dict := map[string]int{
		"Book":   0,
		"Movie":  0,
		"Series": 0,
		"Game":   0,
	}

	for _, review := range u.Finished {
		switch review.ReviewableType {
		case "Book":
			dict["Book"] = dict["Book"] + 1
		case "Movie":
			dict["Movie"] = dict["Movie"] + 1
		case "Series":
			dict["Series"] = dict["Series"] + 1
		case "Game":
			dict["Game"] = dict["Game"] + 1
		}
	}

	return dict
}

func (u *User) GetLikes() []*reviewable.Like {
	return u.Likes
}
