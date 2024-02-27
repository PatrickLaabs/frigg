package capd

import (
	"fmt"
	"io"
	"os"
	"time"
)

func MakeReadme(filename string) {
	date := time.Now().Format("2 Jan 2006")

	header := "# CAPD (ClusterAPI Provider for Docker) Documentation"
	body := "## Get started \n\n" +
		"You will need to pass two Environment variables: \n" +
		"- GITHUB_TOKEN\n" +
		"- GITHUB_USERNAME\n" +
		"- GITHUB_USERNAME_EMAIL"
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
