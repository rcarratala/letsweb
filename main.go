package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	// Raise an 404 if the url not matches
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from the Webserver"))
}

// Function to manage the /template?id=XXX
func showTemplate(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// Check if the id can't be converted to integer and if it's negative
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// use the Fprintf to write to the variable of the w ResponseWriter
	fmt.Fprintf(w, "Display a specific template with ID %d...", id)
}

func createTemplate(w http.ResponseWriter, r *http.Request) {
	// Raise a Method Not Allowed or 405
	if r.Method != "POST" {
		// Use the Header().Set() method to add an 'Allow: POST' header to the response header map
		w.Header().Set("Allow", "POST")

		// Use instead the http.Error()
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))

		//Use the http.Error() function to send a 405 status code
		http.Error(w, "Method Not Allowed. Please check Allowed Methods", 405)
		return
	}

	w.Write([]byte("Create a new template"))
}

func main() {
	// Initialize a new servemux
	mux := http.NewServeMux()
	// Register a / as the index
	mux.HandleFunc("/", index)

	mux.HandleFunc("/template", showTemplate)

	mux.HandleFunc("/template/create", createTemplate)

	log.Println("Starting Server on :3000")
	// Starting a new web server on 3000 port and log the errors if any
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
