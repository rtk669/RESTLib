package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"lib/lib"
	"net/http"
	"strconv"
)

type LibHandlers struct {
	library *lib.Library
}

func NewHandlers(library *lib.Library) *LibHandlers {
	return &LibHandlers{
		library: library,
	}
}

/*
	/books
	POST
	json DTO
*/

func (h *LibHandlers) AddBook(w http.ResponseWriter, r *http.Request) {
	var bookDto BookDTO
	encdr := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&bookDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := encdr.Encode(err.Error()); err != nil {
			fmt.Println(err)
		}
		return
	}
	book, err := h.library.AddBook(bookDto.Title, bookDto.Author, bookDto.PageAmount)
	if err != nil {
		switch {
		case errors.Is(err, lib.ErrNoTitle) ||
			errors.Is(err, lib.ErrNoAuthor) ||
			errors.Is(err, lib.ErrInvalidPageAmount):
			w.WriteHeader(http.StatusBadRequest)
			if err := encdr.Encode(err.Error()); err != nil {
				fmt.Println(err)
			}
			return
		case errors.Is(err, lib.ErrBookAlreadyExist):
			w.WriteHeader(http.StatusConflict)
			if err := encdr.Encode(err.Error()); err != nil {
				fmt.Println(err)
			}
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		if err := encdr.Encode(err.Error()); err != nil {
			fmt.Println(err)
		}
		return
	}

}

func (h *LibHandlers) GetBook(w http.ResponseWriter, r *http.Request) {
	encdr := json.NewEncoder(w)
	title := r.PathValue("title")

	book, err := h.library.GetBook(title)

	if err != nil {
		if errors.Is(err, lib.ErrBookNotFound) {
			w.WriteHeader(http.StatusNotFound)
			if err := encdr.Encode(err.Error()); err != nil {
				fmt.Println(err)
			}
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
	if err := json.NewEncoder(w).Encode(book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *LibHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	authorExist := r.URL.Query().Has("author")
	author := r.URL.Query().Get("author")

	completedExist := r.URL.Query().Has("completed")
	completedStr := r.URL.Query().Get("completed")
	var completed bool

	if completedExist {
		var err error
		completed, err = strconv.ParseBool(completedStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	bookMap := h.library.GetBooks(authorExist, completedExist, author, completed)

	if err := json.NewEncoder(w).Encode(bookMap); err != nil {
		fmt.Println(err)
	}

}

func (h *LibHandlers) CompleteBook(w http.ResponseWriter, r *http.Request) {
	encdr := json.NewEncoder(w)
	title := r.PathValue("title")

	book, err := h.library.CompleteBook(title)
	if err != nil {
		switch {
		case errors.Is(err, lib.ErrBookNotFound):
			w.WriteHeader(http.StatusNotFound)
			if err := encdr.Encode(err.Error()); err != nil {
				fmt.Println(err)
			}
			return
		case errors.Is(err, lib.ErrBookAreadyCompleted):
			w.WriteHeader(http.StatusConflict)
			if err := encdr.Encode(err.Error()); err != nil {
				fmt.Println(err)
			}
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	b, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h *LibHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	encdr := json.NewEncoder(w)
	title := r.PathValue("title")

	err := h.library.DeleteBook(title)
	if err != nil {
		switch {
		case errors.Is(err, lib.ErrBookNotFound):
			w.WriteHeader(http.StatusNotFound)
			if err := encdr.Encode(err.Error()); err != nil {
				fmt.Println(err)
			}
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
