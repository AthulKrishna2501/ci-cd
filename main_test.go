// main_test.go
package main

import (
	"bytes"
	"os"
	"testing"
)

// Helper function to capture the output of a function
func captureOutput(f func()) string {
	// Save the current stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	f()

	// Reset stdout and close the writer
	_ = w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestMainOutput(t *testing.T) {
	expected := "Hello, CI/CD!\n" // Expected output
	output := captureOutput(main) // Capture the output of the main function

	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}
