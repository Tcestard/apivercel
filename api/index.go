package handler

import (
	"fmt"
	"net/http"
	"os/exec"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World from Golang serverless</h1>")

	// Execute the apivercel executable
	cmd := exec.Command("./apivercel") // Assuming apivercel is in the same directory as the handler code
	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(w, "Error executing apivercel: %v", err)
		return
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintf(w, "Command execution failed: %v", err)
		return
	}
}
