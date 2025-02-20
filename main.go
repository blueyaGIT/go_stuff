package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/russross/blackfriday/v2"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the content of README.md
	content, err := ioutil.ReadFile("README.md")
	if err != nil {
		http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the Markdown to HTML
	htmlContent := blackfriday.Run(content)

	// Add CSS for dark mode
	darkModeCSS := `
		<style>
			body {
				background-color: #121212;
				color: #ffffff;
				font-family: Arial, sans-serif;
				padding: 20px;
			}
			h1, h2, h3, h4, h5, h6 {
				color: #bb86fc;
			}
			a {
				color: #03dac6;
				text-decoration: none;
			}
			a:hover {
				text-decoration: underline;
			}
			pre {
				background-color: #1e1e1e;
				color: #ffffff;
				padding: 10px;
				border-radius: 5px;
			}
			code {
				background-color: #1e1e1e;
				color: #ffffff;
				padding: 2px 4px;
				border-radius: 3px;
			}
		</style>
	`

	// Combine the CSS and the rendered Markdown content
	finalHTML := "<html><head>" + darkModeCSS + "</head><body>" + string(htmlContent) + "</body></html>"

	// Set the content type to HTML and write the response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(finalHTML))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
