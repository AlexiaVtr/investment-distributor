package main

import (
	"net/http"
)

// Estructura para los endpoint:
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRounter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}

}

// Asociar Handler con la ruta y crear conexión con handlers.go:

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

// Función para instanciar el Handle en server.go.

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Conexión con el handler si existe:

	handler(w, request)

}
