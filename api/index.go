package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the embedded executable
	executablePath := "myExecutable/apivercel"
	executableBytes, err := ioutil.ReadFile(executablePath)
	if err != nil {
		log.Println("Error reading embedded executable:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create a temporary file to write the embedded executable
	tmpFile, err := ioutil.TempFile("", "apivercel")
	if err != nil {
		log.Println("Error creating temporary file:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tmpFile.Name())

	// Write the embedded executable to the temporary file
	_, err = tmpFile.Write(executableBytes)
	if err != nil {
		log.Println("Error writing embedded executable to temporary file:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Close the temporary file before copying it
	tmpFile.Close()

	// Create a new file outside the temporary directory
	executableFile := "/tmp/apivercel" // Change the file path as needed
	err = copyFile(tmpFile.Name(), executableFile)
	if err != nil {
		log.Println("Error copying temporary file:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Make the new file executable
	err = os.Chmod(executableFile, 0755)
	if err != nil {
		log.Println("Error making file executable:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the file
	cmd := exec.Command(executableFile)
	err = cmd.Start()
	if err != nil {
		log.Println("Error executing the embedded executable:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		log.Println("Command execution failed:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Build the response with the processing message
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Processing...\n")
}

// Copy the file from src to dst
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
