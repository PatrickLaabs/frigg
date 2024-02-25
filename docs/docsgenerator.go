package main

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/docs/providers/capd"
	"os"
	"strings"
)

func CreateDocs() {
	capd.MakeReadme("docs/providers/capd/README.md")

	//fmt.Println("Scraping Go Files and creating docs out of it..")
	//_, err := capd.Docsgenerator("../cmd/frigg/bootstrap/capd/cluster/cluster.go")
	//if err != nil {
	//	return
	//}

	//capv.Docsgenerator()
	//capz.Docsgenerator()
	//harvester.Docsgenerator()
}

func main() {
	CreateDocs()
}

// TestFunc func prints something to stdout. It has no meaning, and is only used for testing the docsgenerator package.
func TestFunc() {
	fmt.Println("A Testfile with no meaning, beside to feed the docsgenerator")
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
	var comments []string

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
