package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the content of README.md
	content, err := ioutil.ReadFile("README.md")
	if err != nil {
		http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type to plain text (or text/markdown if you prefer)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(content)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
