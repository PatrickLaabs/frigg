package helmchartsproxies

// The package helmchartproxies generates various helmchartproxy yaml files to the .frigg directory
// which will be installed onto the management cluster.
import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
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
}

func MgmtArgoCD() {
	fmt.Println("Generating Mgmts ArgoCD Helmchartproxy")
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
