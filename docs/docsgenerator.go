package main

import (
	"github.com/PatrickLaabs/frigg/docs/providers/capd"
)

func CreateDocs() {
	capd.MakeReadme("docs/providers/capd/README.md")
}

func main() {
	CreateDocs()
}
