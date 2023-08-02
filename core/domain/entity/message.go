package entity

type Message struct {
	ID        int64  `json:"id"`
	Body      string `json:"body"`
	ProfileID int64  `json:"-"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}
