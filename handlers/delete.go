package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
)

func (app *Application) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("postID")

		postID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("ERROR: Bad input"))
			return
		}

		err = app.Model.DeletePost(postID)
		if err != nil {
			if err == sql.ErrNoRows || err.Error() == "enr" {
				w.WriteHeader(404)
				w.Write([]byte("ERROR: 404 post not found"))
				return
			}

			w.WriteHeader(500)
			w.Write([]byte("ERROR: can't delete post"))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("Post successfully deleted"))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("ERROR: Method Not Allowed"))
	return
}
