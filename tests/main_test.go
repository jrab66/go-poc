// main_test.go

package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Redirect standard output to capture printed output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Ensure the original state is restored after the test
	defer func() {
		os.Stdout = oldStdout
	}()

	// Run the main function
	main()

	// Close the write end of the pipe and read the captured output
	w.Close()
	capturedOutput := make([]byte, 100)
	n, _ := r.Read(capturedOutput)

	// Check if the output matches the expected value
	expected := "hello"
	actual := string(capturedOutput[:n])
	if actual != expected {
		t.Errorf("Expected: %s, Got: %s", expected, actual)
	}
}
