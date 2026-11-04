package lib

import "errors"

var ErrNoTitle = errors.New("No title")
var ErrNoAuthor = errors.New("No author")
var ErrInvalidPageAmount = errors.New("Invalid page amount")
var ErrBookNotFound = errors.New("Book not found")
var ErrBookAlreadyExist = errors.New("Book already exist")
var ErrBookAreadyCompleted = errors.New("Book already completed")
