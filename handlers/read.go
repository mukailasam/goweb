package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func (app *Application) ReadPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("postID")

		postID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("ERROR: Bad input"))
			return
		}

		p, err := app.Model.ReadPost(postID)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(404)
				w.Write([]byte("ERROR: 404 post not found"))
				return
			}

			w.WriteHeader(500)
			w.Write([]byte("ERROR: can't read post"))
			return
		}

		w.WriteHeader(200)
		blog := fmt.Sprintf("Blog Title: %s\nBlog Body: %s \n", p.Title, p.Body)
		w.Write([]byte(blog))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed"))
	return
}
