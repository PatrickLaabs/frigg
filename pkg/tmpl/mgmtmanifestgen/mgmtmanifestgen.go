package mgmtmanifestgen

import (
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

func Gen() {
	println(color.YellowString("Getting Management Clusters Manifest from Github"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	outputPath := filepath.Join(friggDir, vars.MgmtManifest)

	cmd := exec.Command("curl", "-L", "-o", outputPath,
		vars.CurlMgmtManifest,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.YellowString(string(output)))
		return
	}
	println(color.GreenString("Successfully written Mgmt Kubernets Manifest"))
}
