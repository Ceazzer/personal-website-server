package domain

type Message struct {
	ID        int64  `json:"id"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	ProfileID int64  `json:"-"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}
