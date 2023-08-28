package models

type UserReviewable struct {
	UserID         uint
	ReviewableID   uint
	ReviewableType string
	PlanTo         bool
	Current        bool
	Finished       bool
}
