package manifestmodifier

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestModifyMgmt(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "frigg-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(tempDir)

	// Create the expected directory structure and files
	argohubDir := filepath.Join(tempDir, ".frigg")
	err = os.MkdirAll(argohubDir, 0750)
	if err != nil {
		return
	}

	mgmtClusterManifestPath := filepath.Join(argohubDir, "argohubmgmtclusterManifest.yaml")
	initialYamlContent := []byte(`---
apiVersion: cluster.x-k8s.io/v1beta1
kind: Cluster
metadata:
  name: argohubcluster
  namespace: default
`)
	err = os.WriteFile(mgmtClusterManifestPath, initialYamlContent, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Call the function to modify the file
	ModifyMgmt()

	// Read the modified file content
	modifiedYaml, err := os.ReadFile(mgmtClusterManifestPath)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the modified data
	var modifiedData map[interface{}]interface{}
	err = yaml.Unmarshal(modifiedYaml, &modifiedData)
	if err != nil {
		t.Fatal(err)
	}

	// Expected labels
	expectedLabels := map[string]interface{}{
		"argocd-hub-Enabled":     "default",
		"argocd-hub-appsEnabled": "default",
		"argo-events-Enabled":    "default",
		"argo-rollouts-Enabled":  "default",
		"argo-workflows-Enabled": "default",
		"vaultmgmtEnabled":       "default",
	}

	// Get the actual labels from the modified data
	metadata := modifiedData["metadata"].(map[interface{}]interface{})
	labels := metadata["labels"].(map[string]interface{})

	// Compare the expected and actual labels
	if !reflect.DeepEqual(labels, expectedLabels) {
		t.Errorf("Expected labels: %v, Got: %v", expectedLabels, labels)
	}
}
