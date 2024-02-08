package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestReadComments(t *testing.T) {
	// Replace with your actual Go file path
	filePath := "docsgenerator.go"
	comments, err := ReadComments(filePath)
	if err != nil {
		t.Errorf("Error reading comments: %v", err)
	}
	assert.Contains(t, comments, "TestFunc func prints something to stdout. It has no meaning, and is only used for testing the docsgenerator package.")
}

func ReadComments(filePath string) ([]string, error) {
	// Use the os.ReadFile function to read the file contents
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Convert file contents to string
	fileContent := string(fileBytes)

	// Initialize empty slice to store comments
	comments := []string{}

	// Iterate over each line in the file content
	for _, line := range strings.Split(fileContent, "\n") {
		// Check for single-line comments starting with "//"
		if strings.HasPrefix(line, "//") {
			comments = append(comments, strings.TrimSpace(line[2:]))
		}

		// Check for multi-line comments starting with "/*` and ending with `*/"
		if strings.HasPrefix(line, "/*") && strings.HasSuffix(line, "*/") {
			comments = append(comments, strings.TrimSpace(strings.TrimPrefix(strings.TrimSuffix(line, "*/"), "/*")))
		}
	}

	return comments, nil
}
