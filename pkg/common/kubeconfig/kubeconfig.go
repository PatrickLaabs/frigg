package kubeconfig

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 6. retrieve kubeconfig for the argphub-mgmt-cluster from the bootstrap cluster
//	  => clusterctl --kubeconfig bootstrapcluster.kubeconfig get kubeconfig argohubmgmtcluster > argohubmgmtcluster.kubeconfig

// RetrieveMgmtKubeconfig retrieves of the newly provisioned mgmt-cluster via capd
// and stores it to the work directory of frigg
func RetrieveMgmtKubeconfig() {
	fmt.Println("Retrieving kubeconfig of mgmt-cluster")

	homedir, _ := os.UserHomeDir()
	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"
	argohubDir := homedir + "/" + argohubDirName

	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName
	fmt.Println("Path to Kubeconfig File:", kubeconfigFlagPath)

	cmd := exec.Command("clusterctl", "--kubeconfig",
		kubeconfigFlagPath, "get", "kubeconfig", "argohubmgmtcluster",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))

	err = os.WriteFile(argohubDir+"/"+"argohubmgmtcluster.kubeconfig", output, 0755)
	if err != nil {
		return
	}
}

// 8. modify the kubeconfig, to continue working with it
//    => sed -i -e "s/server:.*/server: https:\/\/$(docker port argohubmgmtcluster-lb 6443/tcp | sed "s/0.0.0.0/127.0.0.1/")/g" ./argohubmgmtcluster.kubebeconfig

// ModifyMgmtKubeconfig modifies the stored kubeconfig, so its working on macOS
func ModifyMgmtKubeconfig() {
	fmt.Println("Modifying mgmt-clusters kubeconfig..")

	homedir, _ := os.UserHomeDir()
	argohubDirName := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"
	kubeconfigNameNew := "argohubmgmtcluster.kubeconfig.new"

	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName
	kubeconfigFlagPathNew := homedir + "/" + argohubDirName + "/" + kubeconfigNameNew
	fmt.Println("Path to Kubeconfig File:", kubeconfigFlagPath)

	data, err := os.ReadFile(kubeconfigFlagPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var newLines []string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "server:") {
			//port := "6443" // Use environment variable or another mechanism if needed
			replacedLine := strings.ReplaceAll(line, `server:.*`, fmt.Sprintf("server: https://%s/tcp", strings.ReplaceAll(os.Args[1], "0.0.0.0", "127.0.0.1")))
			newLines = append(newLines, replacedLine)
		} else {
			newLines = append(newLines, line)
		}
	}

	err = os.WriteFile(kubeconfigFlagPathNew, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing new file:", err)
		return
	}

	// Optionally move the new file to the original name
	err = os.Rename(kubeconfigFlagPathNew, kubeconfigFlagPath)
	if err != nil {
		fmt.Println("Error moving new file:", err)
	}
}
