package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World from Golang serverless</h1>")

	// Display current directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(w, "Error getting current directory: %v", err)
		return
	}
	fmt.Fprintf(w, "<p>Current Directory: %s</p>", currentDir)

	// List files in current directory
	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		fmt.Fprintf(w, "Error listing files: %v", err)
		return
	}

	fmt.Fprintf(w, "<h2>Files:</h2>")
	fmt.Fprintf(w, "<ul>")
	for _, file := range files {
		fmt.Fprintf(w, "<li>%s</li>", file.Name())
	}
	fmt.Fprintf(w, "</ul>")
}

