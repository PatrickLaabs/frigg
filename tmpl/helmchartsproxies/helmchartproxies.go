package helmchartsproxies

// The package helmchartproxies generates various helmchartproxy yaml files to the .frigg directory
// which will be installed onto the management cluster.
import (
	"fmt"
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
	fmt.Println("Generating Mgmts Argo Apps Helmchartproxy")

	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	filePath := "templates/helmchartproxies/mgmt-argocd-apps.yaml"
	newFile := "mgmt-argocd-apps.yaml"
	newfilePath := friggDir + "/" + newFile

	username, err := retrieveUsername()
	if err != nil {
		fmt.Println("Error retrieving token:", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	modifiedLines := make([]string, 0, len(lines))

	// URL should be formated like this: https://github.com/<USERNAME>/argo-hub.git
	re := regexp.MustCompile(`PLACEHOLDER`)

	url := "https://github.com/" + username + "/argo-hub.git"

	for _, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
			fmt.Println("inside match statement")

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

func MgmtArgoCD() {
	fmt.Println("Generating Mgmts ArgoCD Helmchartproxy")

	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	filePath := "templates/helmchartproxies/mgmt-argocd.yaml"
	newFile := "mgmt-argocd.yaml"
	newfilePath := friggDir + "/" + newFile

	username, err := retrieveUsername()
	if err != nil {
		fmt.Println("Error retrieving github username:", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	modifiedLines := make([]string, 0, len(lines))

	// URL should be formated like this: https://github.com/<USERNAME>/argo-hub.git
	re := regexp.MustCompile(`PLACEHOLDER`)

	url := "https://github.com/" + username + "/argo-hub.git"

	for _, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
			fmt.Println("inside match statement")

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
