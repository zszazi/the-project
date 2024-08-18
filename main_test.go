package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	// Create a new ServeMux and add handlers as in main()
	mux := http.NewServeMux()
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/about.html")
	})
	fsRoot := http.FileServer(http.Dir("static"))
	mux.Handle("/", http.StripPrefix("/", fsRoot))

	// Create a test server
	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}

	// Test /home route
	resp, err := client.Get(server.URL + "/home")
	if err != nil {
		t.Fatalf("Failed to get /home: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK for /home, got %v", resp.Status)
	}

	// Test /about route
	resp, err = client.Get(server.URL + "/about")
	if err != nil {
		t.Fatalf("Failed to get /about: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK for /about, got %v", resp.Status)
	}

	// Test static files served from root
	resp, err = client.Get(server.URL + "/index.html")
	if err != nil {
		t.Fatalf("Failed to get /index.html: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK for /index.html, got %v", resp.Status)
	}
}
