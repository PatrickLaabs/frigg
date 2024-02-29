package helmchartsproxies

// The package helmchartproxies generates various helmchartproxy yaml files to the .frigg directory
// which will be installed onto the management cluster.
import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"strings"
)

// Cni generates the CNI Helmchartproxy YAML file to the .frigg directory
func Cni() {
	fmt.Println("Generating CNI Calico Helmchartproxy")

	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	filePath := "templates/helmchartproxies/cni.yaml"
	newFile := "cni_gen.yaml"
	newfilePath := friggDir + "/" + newFile
	message := "bla"

	err := AppendNewLineToFile(filePath, newfilePath, message)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully appended new line to", newfilePath)
	}
}

func Vault() {
	fmt.Println("Generating Vault Helmchartproxy")
}

func ArgoWorkflows() {
	fmt.Println("Generating Argo Workflows Helmchartproxy")
}

func ArgoRollouts() {
	fmt.Println("Generating Argo Rollouts Helmchartproxy")
}

func ArgoEvents() {
	fmt.Println("Generating Argo Events Helmchartproxy")
}

func MgmtArgoApps() {
	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	filePath := "templates/helmchartproxies/mgmt-argocd-apps.yaml"
	newFile := "mgmt-argocd-apps.yaml"
	newfilePath := friggDir + "/" + newFile

	username, err := retrieveUsername()
	if err != nil {
		println(color.RedString("Error retrieving token: %v\n", err))
		os.Exit(1)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	modifiedLines := make([]string, 0, len(lines))

	re := regexp.MustCompile(`PLACEHOLDER`)

	// URL should be formated like this: ssh://git@github.com:<USERNAME>/argo-hub.git
	url := "ssh://git@github.com:" + username + "/argo-hub.git"
	// url := "https://github.com/" + username + "/argo-hub.git"
	for _, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
			newUrl := fmt.Sprintf(url)

			modifiedLine := re.ReplaceAllString(line, fmt.Sprintf(newUrl))
			modifiedLines = append(modifiedLines, modifiedLine)
		} else {
			modifiedLines = append(modifiedLines, line)
		}
	}

	err = os.WriteFile(newfilePath, []byte(strings.Join(modifiedLines, "\n")), 0755)
	if err != nil {
		return
	}
}

//func MgmtArgoCD() {
//	homedir, _ := os.UserHomeDir()
//	friggDirName := ".frigg"
//	friggDir := homedir + "/" + friggDirName
//
//	filePath := "templates/helmchartproxies/mgmt-argocd.yaml"
//	newFile := "mgmt-argocd.yaml"
//	newfilePath := friggDir + "/" + newFile
//
//	username, err := retrieveUsername()
//	if err != nil {
//		fmt.Println("Error retrieving github username:", err)
//		os.Exit(1)
//	}
//
//	data, err := os.ReadFile(filePath)
//	if err != nil {
//		return
//	}
//
//	lines := strings.Split(string(data), "\n")
//	modifiedLines := make([]string, 0, len(lines))
//
//	re := regexp.MustCompile(`PLACEHOLDER`)
//
//	// URL should be formated like this: ssh://git@github.com:<USERNAME>/argo-hub.git
//	// url := "ssh://git@github.com:" + username + "/argo-hub.git"
//	url := "https://github.com/" + username + "/argo-hub.git"
//	for _, line := range lines {
//		if match := re.FindStringSubmatch(line); match != nil {
//			newUrl := fmt.Sprintf(url)
//
//			modifiedLine := re.ReplaceAllString(line, fmt.Sprintf(newUrl))
//			modifiedLines = append(modifiedLines, modifiedLine)
//		} else {
//			modifiedLines = append(modifiedLines, line)
//		}
//	}
//
//	err = os.WriteFile(newfilePath, []byte(strings.Join(modifiedLines, "\n")), 0755)
//	if err != nil {
//		return
//	}
//}

func MgmtArgoCD() {
	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	username, err := retrieveUsername()
	if err != nil {
		fmt.Println("Error retrieving github username:", err)
		os.Exit(1)
	}

	filePath := "templates/helmchartproxies/mgmt-argocd_ssh.yaml"
	newFile := "mgmt-argocd.yaml"
	newfilePath := friggDir + "/" + newFile

	sshprivatekeyPath := friggDir + "/" + "frigg-sshkeypair_gen"
	sshprivatekey, err := os.ReadFile(sshprivatekeyPath)
	sshprivatekeyTrimmed := strings.TrimSuffix(string(sshprivatekey), "\n")

	fmt.Println(sshprivatekeyTrimmed)
	err = MGmtArgoCDReplacementTest(filePath, newfilePath, username, sshprivatekeyTrimmed)
	if err != nil {
		println(color.RedString("error on string replacement for sshkeypair: %v\n", err))
	}
}

func MGmtArgoCDReplacementTest(filePath string, newFilePath string, username string, sshprivatekey string) error {
	// Read file contents
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	username, err = retrieveUsername()
	if err != nil {
		fmt.Println("Error retrieving github username:", err)
		os.Exit(1)
	}

	reGhUser := regexp.MustCompile(`PLACEHOLDER`)
	reSshKey := regexp.MustCompile(`SSHREPLACEMENT`)

	url := "ssh://git@github.com:" + username + "/argo-hub.git"

	// Split the YAML data into lines
	yamlLines := strings.Split(string(data), "\n")

	// Find the line containing the placeholder and extract its indentation
	var placeholderLine string
	var originalIndent string // Placeholder for extracted indentation
	for _, line := range yamlLines {
		if strings.Contains(line, "SSHREPLACEMENT") {
			placeholderLine = line
			originalIndent = extractIndentation(placeholderLine)
			break
		}
	}

	// If placeholder not found, return an error
	if originalIndent == "" {
		return fmt.Errorf("placeholder 'SSHREPLACEMENT' not found in the YAML file")
	}

	// Indent each line of the SSH private key with the relative indentation
	indentedSshKey := strings.ReplaceAll(sshprivatekey, "\n", "\n"+originalIndent)

	// Replace GITHUB_USER and sshkey
	newdata := replaceInString(data, reGhUser, url)
	newdata = replaceInString(newdata, reSshKey, indentedSshKey)

	// Write modified content back to the file
	err = os.WriteFile(newFilePath, newdata, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

// replaceInString replaces specific pattern with a new string
func replaceInString(data []byte, re *regexp.Regexp, replacement string) []byte {
	return re.ReplaceAll(data, []byte(replacement))
}

// Helper function to extract indentation from a line
func extractIndentation(line string) string {
	for i, char := range line {
		if char != ' ' && char != '\t' {
			return line[:i]
		}
	}
	return ""
}

func retrieveUsername() (string, error) {
	// Get GITHUB_USERNAME environment var
	var username string

	if os.Getenv("GITHUB_USERNAME") == "" {
		fmt.Println("Missing Github Username, please set it. Exiting now.")
		os.Exit(1)
	} else {
		username = os.Getenv("GITHUB_USERNAME")
	}

	return username, nil
}

// AppendNewLineToFile appends a new line with the given message to a YAML file and saves it as a new file.
func AppendNewLineToFile(filePath, newFile, message string) error {
	// Read the original YAML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal the YAML data
	var content map[string]interface{}
	err = yaml.Unmarshal(data, &content)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML: %w", err)
	}

	// Add the new line under the valuesTemplate string
	values, ok := content["metadata"].(map[string]interface{})["name"].(string)
	if !ok {
		return fmt.Errorf("key 'valuesTemplate' not found or not a string")
	}
	fmt.Println(values)

	// Append the message with indentation and a newline
	values = message

	// Update the map with the modified value
	content["metadata"].(map[string]interface{})["name"] = values

	// Marshal the modified data back to YAML
	newData, err := yaml.Marshal(content)
	if err != nil {
		return fmt.Errorf("error marshalling modified data: %w", err)
	}

	// Write the modified data to a new file
	err = os.WriteFile(newFile, newData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
