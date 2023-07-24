package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	app := NewApplication()

	mux.HandleFunc("/", app.Index)
	mux.HandleFunc("/create", app.CreatePost)
	mux.HandleFunc("/read", app.ReadPost)
	mux.HandleFunc("/delete", app.DeletePost)

	fmt.Println("listen on port 8081")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Panic(err)
	}
}
