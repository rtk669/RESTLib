package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Handlers *LibHandlers
}

func NewServer(handlers *LibHandlers) *Server {
	return &Server{
		Handlers: handlers,
	}
}

func (s *Server) StartServer() error {
	router := http.NewServeMux()

	router.HandleFunc("POST /books", s.Handlers.AddBook)
	router.HandleFunc("GET /books", s.Handlers.GetBooks)
	router.HandleFunc("GET /books/{title}", s.Handlers.GetBook)
	router.HandleFunc("PATCH /books/{title}", s.Handlers.CompleteBook)
	router.HandleFunc("DELETE /books/{title}", s.Handlers.DeleteBook)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", router); err != nil {
		return err
	}

	return nil
}
