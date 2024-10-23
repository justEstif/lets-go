package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// use the Header().Add() method to add a 'Server: Go' header to
	// the response header map. The first parameter is the header name,
	// the second parameter is the header value
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

// add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// extract the value of the id wild card from the request using r.PathValue()
	// and try to convert it to an integer using the strconv.Atoi() function. If
	// it can't be converted to an integer, or the value is less than 1, we return
	// a 404 page not found response

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	// w.Write([]byte(msg))
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// add a snippetCreatePost handler function
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// use the w.WriteHeader() method to send a 201 status code.
	w.WriteHeader(http.StatusCreated)
	// then w.Write() method write the response body
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
