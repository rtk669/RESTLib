package server

type BookDTO struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	PageAmount int    `json:"pageAmount"`
}
