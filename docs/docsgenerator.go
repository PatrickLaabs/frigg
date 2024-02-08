package main

import (
	"fmt"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/docs/capd"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/docs/capv"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/docs/capz"
)

func CreateDocs() {
	fmt.Println("Scraping Go Files and creating docs out of it..")
	capd.Docsgenerator()
	capv.Docsgenerator()
	capz.Docsgenerator()
}

func main() {
	CreateDocs()
}

// TestFunc func prints something to stdout. It has no meaning, and is only used for testing the docsgenerator package.
func TestFunc() {
	fmt.Println("A Testfile with no meaning, beside to feed the docsgenerator")
}
