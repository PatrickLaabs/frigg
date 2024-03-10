package kubeconfig

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// RetrieveMgmtKubeconfig retrieves of the newly provisioned mgmt-cluster via capd
// and stores it to the work directory of frigg
func RetrieveMgmtKubeconfig() {
	println(color.GreenString("Retrieving kubeconfig of the mgmt-cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := homedir + "/" + vars.FriggDirName

	kubeconfigFlagPath := friggDir + "/" + vars.BootstrapkubeconfigName

	cmd := exec.Command("clusterctl", "--kubeconfig",
		kubeconfigFlagPath, "get", "kubeconfig", vars.FriggMgmtName,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running clusterctl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}

	err = os.WriteFile(friggDir+"/"+vars.ManagementKubeconfigName, output, 0755)
	if err != nil {
		println(color.RedString("Error on writing kubeconfig file for mgmt cluster: %v\n", err))
		return
	}
}

// RetrieveWorkloadKubeconfig retrieves kubeconfig of the workload cluster
func RetrieveWorkloadKubeconfig() {
	println(color.GreenString("Retrieving kubeconfig of workload-cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := homedir + "/" + vars.FriggDirName

	kubeconfigFlagPath := friggDir + "/" + vars.ManagementKubeconfigName

	cmd := exec.Command("clusterctl", "--kubeconfig",
		kubeconfigFlagPath, "get", "kubeconfig", "workloadcluster",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running clusterctl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}

	err = os.WriteFile(friggDir+"/"+vars.WorkloadKubeconfigName, output, 0755)
	if err != nil {
		println(color.RedString("Error on writing kubeconfig file for mgmt cluster: %v\n", err))
		return
	}
}

// ModifyMgmtKubeconfig modifies the stored kubeconfig, so it's working on macOS
func ModifyMgmtKubeconfig() error {
	println(color.GreenString("Modifying the kubeconfig file of the mgmt cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
	}

	kubeconfigNameNew := vars.ManagementKubeconfigName + ".new"
	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.ManagementKubeconfigName

	data, err := os.ReadFile(kubeconfigFlagPath)
	if err != nil {
		return fmt.Errorf("error reading kubeconfig: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	modifiedLines := make([]string, 0, len(lines))

	re := regexp.MustCompile(`^ {4}server: (https?://[^:]+:\d+)`) // Capture the URL

	var port string

	// Fallback to using docker port command only if environment variable is not set
	portBytes, err := exec.Command("docker", "port", "friggmgmtcluster-lb", "6443/tcp").Output()
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

	if err != nil {
		return fmt.Errorf("invalid port: %s", port)
	}

	for _, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
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

	println(color.GreenString("Kubeconfig of mgmt cluster successfully modified"))
	return nil
}

// ModifyWorkloadKubeconfig modifies the stored kubeconfig, so it's working on macOS
func ModifyWorkloadKubeconfig() error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
	}

	kubeconfigNameNew := vars.WorkloadKubeconfigName + ".new"
	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.WorkloadKubeconfigName

	data, err := os.ReadFile(kubeconfigFlagPath)
	if err != nil {
		return fmt.Errorf("error reading kubeconfig: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	modifiedLines := make([]string, 0, len(lines))

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
	if err != nil {
		return fmt.Errorf("invalid port: %s", port)
	}

	for _, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
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

	println(color.GreenString("Kubeconfig of workloadcluster successfully modified"))
	return nil
}
