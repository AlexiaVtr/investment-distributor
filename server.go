package main

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

// Obtiene el puerto y lo retorna con la ruta:
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

// Agrega la ruta y método al Handler del main:

func (s *Server) Handle(path string, method string,
	handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}
