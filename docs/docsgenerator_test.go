package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestReadComments tries to test the readability of the comments inside the code.
// It tries to call the function 'ReadComments' from the docsgenerator.go file.
func TestReadComments(t *testing.T) {
	filePath := "docsgenerator.go"
	comments, err := ReadComments(filePath)
	if err != nil {
		t.Errorf("Error reading comments: %q", err)
	}
	assert.Contains(t, comments, "TestFunc func prints something to stdout. It has no meaning, and is only used for testing the docsgenerator package.")
}
