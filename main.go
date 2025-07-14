package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	staticContent, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Println("[ERROR] Error while loading static files")
		return
	}
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticContent))))

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
  log.Println("[INFO] Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
