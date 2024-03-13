package workloadmanifestgen

import (
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

func Gen() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	outputPath := filepath.Join(friggDir, vars.WorkloadManifest)

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
