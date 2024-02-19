package kubeconfig

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
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

// RetrieveWorkloadKubeconfig retrieves kubeconfig of the workload cluster
func RetrieveWorkloadKubeconfig() {
	fmt.Println("Retrieving kubeconfig of workload-cluster")

	homedir, _ := os.UserHomeDir()
	argohubDirName := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"
	argohubDir := homedir + "/" + argohubDirName

	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

	cmd := exec.Command("clusterctl", "--kubeconfig",
		kubeconfigFlagPath, "get", "kubeconfig", "workloadcluster",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))

	err = os.WriteFile(argohubDir+"/"+"workloadcluster.kubeconfig", output, 0755)
	if err != nil {
		return
	}
}

// 8. modify the kubeconfig, to continue working with it
//    => sed -i -e "s/server:.*/server: https:\/\/$(docker port argohubmgmtcluster-lb 6443/tcp | sed "s/0.0.0.0/127.0.0.1/")/g" ./argohubmgmtcluster.kubebeconfig

// ModifyMgmtKubeconfig modifies the stored kubeconfig, so it's working on macOS
func ModifyMgmtKubeconfig() error {
	homedir, _ := os.UserHomeDir()
	argohubDirName := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"
	kubeconfigNameNew := "argohubmgmtcluster.kubeconfig.new"

	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName
	fmt.Println(kubeconfigFlagPath)

	data, err := os.ReadFile(kubeconfigFlagPath)
	if err != nil {
		return fmt.Errorf("error reading kubeconfig: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	modifiedLines := make([]string, 0, len(lines))

	// re := regexp.MustCompile(`^server:\s+(?P<url>.*)$`) // Matches "server:" line with URL capture group
	// re := regexp.MustCompile(`^server:\s+(https?://[^:]+:\d+)/tcp`)
	// re := regexp.MustCompile(` {4}server:`)
	re := regexp.MustCompile(`^ {4}server: (https?://[^:]+:\d+)`) // Capture the URL

	var port string

	// Fallback to using docker port command only if environment variable is not set
	portBytes, err := exec.Command("docker", "port", "argohubmgmtcluster-lb", "6443/tcp").Output()
	if err != nil {
		return fmt.Errorf("error getting container port: %w", err)
	}
	port = strings.TrimSpace(string(portBytes))

	// Split the port string to extract the actual port number
	portParts := strings.Split(port, ":")
	if len(portParts) != 2 {
		return fmt.Errorf("invalid port format: %s", port)
	}
	port = portParts[1]

	portInt, err := strconv.Atoi(port)
	fmt.Printf("portInt value:%v\n", portInt)
	if err != nil {
		return fmt.Errorf("invalid port: %s", port)
	}

	for _, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
			fmt.Println("Inside match statement")

			// Use the captured URL in the replacement
			newURL := fmt.Sprintf("https://%s:%v", "127.0.0.1", portInt)

			// Replace only the captured portion of the line
			modifiedLine := re.ReplaceAllString(line, fmt.Sprintf("    server: %s", newURL))
			modifiedLines = append(modifiedLines, modifiedLine)

		} else {
			// Keep lines without matches
			modifiedLines = append(modifiedLines, line)
		}
	}

	err = os.WriteFile(kubeconfigNameNew, []byte(strings.Join(modifiedLines, "\n")), 0755)
	if err != nil {
		return fmt.Errorf("error writing new kubeconfig: %w", err)
	}

	// Optionally move the new file to the original name with error handling
	err = os.Rename(kubeconfigNameNew, kubeconfigFlagPath)
	if err != nil {
		return fmt.Errorf("error renaming new kubeconfig: %w", err)
	}

	fmt.Println("Kubeconfig of mgmt cluster successfully modified")
	return nil
}

// ModifyWorkloadKubeconfig modifies the stored kubeconfig, so it's working on macOS
func ModifyWorkloadKubeconfig() error {
	homedir, _ := os.UserHomeDir()
	argohubDirName := ".frigg"
	kubeconfigName := "workloadcluster.kubeconfig"
	kubeconfigNameNew := "workloadcluster.kubeconfig.new"

	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

	data, err := os.ReadFile(kubeconfigFlagPath)
	if err != nil {
		return fmt.Errorf("error reading kubeconfig: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	modifiedLines := make([]string, 0, len(lines))

	// re := regexp.MustCompile(`^server:\s+(?P<url>.*)$`) // Matches "server:" line with URL capture group
	// re := regexp.MustCompile(`^server:\s+(https?://[^:]+:\d+)/tcp`)
	// re := regexp.MustCompile(` {4}server:`)
	re := regexp.MustCompile(`^ {4}server: (https?://[^:]+:\d+)`) // Capture the URL

	var port string

	// Fallback to using docker port command only if environment variable is not set
	portBytes, err := exec.Command("docker", "port", "workloadcluster-lb", "6443/tcp").Output()
	if err != nil {
		return fmt.Errorf("error getting container port: %w", err)
	}
	port = strings.TrimSpace(string(portBytes))

	// Split the port string to extract the actual port number
	portParts := strings.Split(port, ":")
	if len(portParts) != 2 {
		return fmt.Errorf("invalid port format: %s", port)
	}
	port = portParts[1]

	portInt, err := strconv.Atoi(port)
	fmt.Printf("portInt value:%v\n", portInt)
	if err != nil {
		return fmt.Errorf("invalid port: %s", port)
	}

	for _, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
			fmt.Println("Inside match statement")

			// Use the captured URL in the replacement
			newURL := fmt.Sprintf("https://%s:%v", "127.0.0.1", portInt)

			// Replace only the captured portion of the line
			modifiedLine := re.ReplaceAllString(line, fmt.Sprintf("    server: %s", newURL))
			modifiedLines = append(modifiedLines, modifiedLine)

		} else {
			// Keep lines without matches
			modifiedLines = append(modifiedLines, line)
		}
	}

	err = os.WriteFile(kubeconfigNameNew, []byte(strings.Join(modifiedLines, "\n")), 0755)
	if err != nil {
		return fmt.Errorf("error writing new kubeconfig: %w", err)
	}

	// Optionally move the new file to the original name with error handling
	err = os.Rename(kubeconfigNameNew, kubeconfigFlagPath)
	if err != nil {
		return fmt.Errorf("error renaming new kubeconfig: %w", err)
	}

	fmt.Println("Kubeconfig of workloadcluster successfully modified")
	return nil
}
