package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Serve index.html at /home
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Serve about.html at /about
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/about.html")
	})

	// Serve static files (e.g., SVG files) from the "static" directory
	fs := http.FileServer(http.Dir("static/asset"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Serve the static folder's root files (e.g., index.html, about.html)
	fsRoot := http.FileServer(http.Dir("static"))
	mux.Handle("/", http.StripPrefix("/", fsRoot))

	// Start the server on port 8080
	log.Println("Serving on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
