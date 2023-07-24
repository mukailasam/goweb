package handlers

import (
	"net/http"
)

func (app *Application) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(200)
		w.Write([]byte("A simple Go program that implement Dependency Injection\n Depend on the methods not the object"))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed"))
	return
}
