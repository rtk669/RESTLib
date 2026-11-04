package lib

import (
	"time"
)

type Book struct {
	Title      string
	Author     string
	PageAmount int
	Completed  bool
	AddTime    time.Time
	ReadTime   *time.Time
}

func NewBook(title string, author string, pageAmount int) (*Book, error) {
	if title == "" {
		return nil, ErrNoTitle
	}

	if author == "" {
		return nil, ErrNoAuthor
	}

	if pageAmount <= 0 {
		return nil, ErrInvalidPageAmount
	}

	return &Book{
		Title:      title,
		Author:     author,
		PageAmount: pageAmount,
		AddTime:    time.Now(),
	}, nil
}

func (b *Book) CompleteBook() {
	b.Completed = true

	now := time.Now()
	b.ReadTime = &now
}
