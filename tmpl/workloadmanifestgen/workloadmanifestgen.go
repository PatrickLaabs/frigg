package workloadmanifestgen

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func Gen() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	// curl -L -o workloadcluster.yaml https://raw.githubusercontent.com/PatrickLaabs/frigg/main/templates/workloadcluster.yaml
	outputPath := homedir + "/" + vars.FriggDirName + "/" + vars.WorkloadManifest

	cmd := exec.Command("curl", "-L", "-o", outputPath,
		vars.CurlWorkloadManifest,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.YellowString(string(output)))
		return
	}
	println(color.GreenString("Successfully written Workload Cluster Kubernets Manifest"))
}
