package tmpl

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/tmpl/mgmtmanifestgen"
	"github.com/PatrickLaabs/frigg/tmpl/workloadmanifestgen"
)

func TemplateGenerator() {
	fmt.Println("Writing Template files")
	mgmtmanifestgen.Gen()
	workloadmanifestgen.Gen()
}
