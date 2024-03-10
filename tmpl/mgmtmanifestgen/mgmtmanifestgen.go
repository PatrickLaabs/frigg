package mgmtmanifestgen

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func Gen() {
	println(color.YellowString("Getting Management Clusters Manifest from Github"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	outputPath := homedir + "/" + vars.FriggDirName + "/" + vars.MgmtManifest

	cmd := exec.Command("curl", "-L", "-o", outputPath,
		vars.CurlMgmtManifest,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.YellowString(string(output)))
		return
	}
	println(color.GreenString("Successfully written Workload Cluster Kubernets Manifest"))

	println(color.GreenString("Successfully written Mgmt Kubernets Manifest"))
}
