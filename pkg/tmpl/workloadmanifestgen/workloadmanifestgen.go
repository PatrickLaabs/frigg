package workloadmanifestgen

import (
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

func GenCapd() {
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

func GenVcluster() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	outputPath := filepath.Join(friggDir, vars.VclusterWorkloadManifest)

	cmd := exec.Command("curl", "-L", "-o", outputPath,
		vars.CurlWorkloadVclusterManifest,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.YellowString(string(output)))
		return
	}
	println(color.GreenString("Successfully written vCluster Workload Cluster Kubernets Manifest"))
}
