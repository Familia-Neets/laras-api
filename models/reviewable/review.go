package reviewable

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID         uint    `json:"user_id" gorm:"not null"`
	Review         string  `json:"review" gorm:"not null"`
	Rating         float32 `json:"rating" gorm:"not null"`
	ReviewableID   uint    `json:"reviewable_id" gorm:"not null"`
	ReviewableType string  `json:"reviewable_type" gorm:"not null"`

	Likes []*Like `json:"likes" gorm:"foreignKey:ReviewID"`
}
