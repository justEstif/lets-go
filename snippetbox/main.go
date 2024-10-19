package main

import (
	"log"
	"net/http"
)

// define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// print a log message to say that the server is starting
	log.Print("starting server on :4000")

	// use the http.listenAndServe() function to start a web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.listenAndServe() returns an error
	// that any error return by http.listenAndServe() is always non-nil
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
