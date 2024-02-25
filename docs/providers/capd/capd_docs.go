package capd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func MakeReadme(filename string) {
	date := time.Now().Format("2 Jan 2006")

	header := "# CAPD (ClusterAPI Provider for Docker) Documentation"
	body := "## Placeholder"
	footer := date
	data := fmt.Sprintf("%s\n\n\n%s\n\n%s", header, body, footer)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error on creating file: %v\n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error on closing file: %v\n", err)
		}
	}(file)

	_, err = io.WriteString(file, data)
	if err != nil {
		fmt.Printf("Error writing README file: %v\n", err)
	}
}

// Docsgenerator for the capd provider usage
func Docsgenerator(filepath string) ([]string, error) {
	//filepath = "cmd/frigg/bootstrap/capd/cluster/cluster.go"
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file from given path: %v", err)
	}

	fileContent := string(fileBytes)

	var comments []string

	//regex := regexp.MustCompile(`// docs:`)
	//
	//for _, line := range strings.Split(fileContent, "\n") {
	//	if match := regex.FindStringSubmatch(line); match != nil {
	//		// Extract and trim the comment text after "// docs:"
	//		comment := strings.TrimSpace(match[1])
	//		comments = append(comments, comment)
	//	}
	//}

	for _, line := range strings.Split(fileContent, "\n") {
		line = strings.TrimLeft(line, " \t")
		if strings.HasPrefix(line, "//") || (strings.HasPrefix(line, "/*") && strings.HasSuffix(line, "*/")) {
			comments = append(comments, strings.TrimSpace(line[2:]))
		}
	}

	//for _, line := range strings.Split(fileContent, "\n") {
	//	if strings.HasPrefix(line, "//") {
	//		comments = append(comments, strings.TrimSpace(line[2:]))
	//	}
	//
	//	if strings.HasPrefix(line, "/*") && strings.HasSuffix(line, "*/") {
	//		comments = append(comments, strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "*/"), "/*")))
	//	}
	//}

	var combinedComments string
	for _, comment := range comments {
		combinedComments += comment + "\n"
	}

	err = os.WriteFile("docs/providers/capd/README.md", []byte(combinedComments), 0755)
	if err != nil {
		return nil, err
	}

	fmt.Println(combinedComments)

	return comments, nil
}
