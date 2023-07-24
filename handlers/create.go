package handlers

import (
	"net/http"
)

func (app *Application) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := "title"
		body := "body"

		err := app.Model.CreatePost(title, body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("ERROR: can't Create Post"))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("Post Created"))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed"))
	return
}
