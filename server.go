package main

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

// Servidor manipulable desde el main, * para evitar copias.
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRounter(),
	}

}

//Método para iniciar el servidor:

func (s *Server) Listen() error {

	// Definición del Endpoint principal:
	http.Handle("/", s.router)

	// Control de errores:
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
