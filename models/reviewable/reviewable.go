package reviewable

type Reviewable interface {
	GetID() uint
	GetType() string
	GetReviews() []*Review
	AppendReview(review *Review)
	GetRating() float32
	UpdateRating() float32
}
