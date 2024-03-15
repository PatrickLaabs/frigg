package main

import (
	"github.com/PatrickLaabs/frigg/docs/providers/capd"
	"github.com/PatrickLaabs/frigg/docs/providers/capd_controller"
)

func CreateDocs() {
	capd.MakeReadme("docs/providers/capd/README.md")
	capd_controller.MakeReadme("docs/providers/capd_controller/README.md")
}

func main() {
	CreateDocs()
}
