package capd

import (
	"fmt"
	"io"
	"os"
	"time"
)

func MakeReadme(filename string) {
	date := time.Now().Format("2 Jan 2006")

	newLine := "\n"
	doubleNewLine := "\n\n"

	header := "# CAPD (ClusterAPI Provider for Docker) Documentation"
	body := "## Get started" + doubleNewLine +
		"You will need to pass two Environment variables:" + newLine +
		"- GITHUB_TOKEN" + newLine +
		"- GITHUB_USERNAME" + newLine +
		"- GITHUB_MAIL" + doubleNewLine +
		"```" + newLine +
		"export GITHUB_TOKEN=" + newLine +
		"export GITHUB_USERNAME=" + newLine +
		"export GITHUB_MAIL" + newLine +
		"```"
	footer := "Updated on: " + date
	data := fmt.Sprintf("%s\n\n%s\n\n%s", header, body, footer)

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
