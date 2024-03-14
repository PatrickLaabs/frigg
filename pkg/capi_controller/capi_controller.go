package capi_controller

import (
	"github.com/PatrickLaabs/frigg/pkg/consts"
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type CapiController struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		Version string `yaml:"version"`
		Manager struct {
			FeatureGates struct {
				MachinePool     bool `yaml:"MachinePool"`
				ClusterTopology bool `yaml:"ClusterTopology"`
			} `yaml:"featureGates"`
		} `yaml:"manager"`
	} `yaml:"spec"`
}

func CoreProviderGen() {
	data := &CapiController{
		APIVersion: "operator.cluster.x-k8s.io/v1alpha2",
		Kind:       "CoreProvider",
		Metadata: struct {
			Name      string `yaml:"name"`
			Namespace string `yaml:"namespace"`
		}{
			Name:      "cluster-api",
			Namespace: "capi-system",
		},
		Spec: struct {
			Version string `yaml:"version"`
			Manager struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			} `yaml:"manager"`
		}{
			Version: consts.CoreProvControllerVersion,
			Manager: struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			}(struct {
				FeatureGates struct {
					MachinePool     bool
					ClusterTopology bool
				} `yaml:"featureGates"`
			}{
				FeatureGates: struct {
					MachinePool     bool
					ClusterTopology bool
				}{
					MachinePool:     true,
					ClusterTopology: true,
				},
			}),
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marshaling data to yaml: %v\n", err))
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	controllerDir := filepath.Join(friggDir, vars.ControllerDir)
	newFilePath := filepath.Join(controllerDir, vars.CoreProviderName)

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	// Write to file
	err = os.WriteFile(newFilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing coreprovider yaml: %v\n", err))
	}
}

func ControlPlaneProviderGen() {
	data := &CapiController{
		APIVersion: "operator.cluster.x-k8s.io/v1alpha2",
		Kind:       "ControlPlaneProvider",
		Metadata: struct {
			Name      string `yaml:"name"`
			Namespace string `yaml:"namespace"`
		}{
			Name:      "kubeadm",
			Namespace: "capi-kubeadm-control-plane-system",
		},
		Spec: struct {
			Version string `yaml:"version"`
			Manager struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			} `yaml:"manager"`
		}{
			Version: consts.KubeadmControllerVersion,
			Manager: struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			}(struct {
				FeatureGates struct {
					MachinePool     bool
					ClusterTopology bool
				} `yaml:"featureGates"`
			}{
				FeatureGates: struct {
					MachinePool     bool
					ClusterTopology bool
				}{
					MachinePool:     true,
					ClusterTopology: true,
				},
			}),
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marshaling data to yaml: %v\n", err))
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	controllerDir := filepath.Join(friggDir, vars.ControllerDir)
	newFilePath := filepath.Join(controllerDir, vars.ControlPlaneProvName)

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	// Write to file
	err = os.WriteFile(newFilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing controlplane provider yaml: %v\n", err))
	}
}

func BootstrapProviderGen() {
	data := &CapiController{
		APIVersion: "operator.cluster.x-k8s.io/v1alpha2",
		Kind:       "BootstrapProvider",
		Metadata: struct {
			Name      string `yaml:"name"`
			Namespace string `yaml:"namespace"`
		}{
			Name:      "kubeadm",
			Namespace: "capi-kubeadm-bootstrap-system",
		},
		Spec: struct {
			Version string `yaml:"version"`
			Manager struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			} `yaml:"manager"`
		}{
			Version: consts.KubeadmControllerVersion,
			Manager: struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			}(struct {
				FeatureGates struct {
					MachinePool     bool
					ClusterTopology bool
				} `yaml:"featureGates"`
			}{
				FeatureGates: struct {
					MachinePool     bool
					ClusterTopology bool
				}{
					MachinePool:     true,
					ClusterTopology: true,
				},
			}),
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marshaling data to yaml: %v\n", err))
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	controllerDir := filepath.Join(friggDir, vars.ControllerDir)
	newFilePath := filepath.Join(controllerDir, vars.BootstrapProvName)

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	// Write to file
	err = os.WriteFile(newFilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing bootstrap provider yaml: %v\n", err))
	}
}

func DockerInfraProviderGen() {
	data := &CapiController{
		APIVersion: "operator.cluster.x-k8s.io/v1alpha2",
		Kind:       "InfrastructureProvider",
		Metadata: struct {
			Name      string `yaml:"name"`
			Namespace string `yaml:"namespace"`
		}{
			Name:      "docker",
			Namespace: "capd-system",
		},
		Spec: struct {
			Version string `yaml:"version"`
			Manager struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			} `yaml:"manager"`
		}{
			Version: consts.DockerInfraProvControllerVersion,
			Manager: struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			}(struct {
				FeatureGates struct {
					MachinePool     bool
					ClusterTopology bool
				} `yaml:"featureGates"`
			}{
				FeatureGates: struct {
					MachinePool     bool
					ClusterTopology bool
				}{
					MachinePool:     true,
					ClusterTopology: true,
				},
			}),
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marshaling data to yaml: %v\n", err))
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	controllerDir := filepath.Join(friggDir, vars.ControllerDir)
	newFilePath := filepath.Join(controllerDir, vars.DockerInfraProvName)

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	// Write to file
	err = os.WriteFile(newFilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing Infra provider yaml: %v\n", err))
	}
}

func AddonHelmProviderGen() {
	data := &CapiController{
		APIVersion: "operator.cluster.x-k8s.io/v1alpha2",
		Kind:       "AddonProvider",
		Metadata: struct {
			Name      string `yaml:"name"`
			Namespace string `yaml:"namespace"`
		}{
			Name:      "helm",
			Namespace: "caaph-system",
		},
		Spec: struct {
			Version string `yaml:"version"`
			Manager struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			} `yaml:"manager"`
		}{
			Version: consts.HelmAddonProvControllerVersion,
			Manager: struct {
				FeatureGates struct {
					MachinePool     bool `yaml:"MachinePool"`
					ClusterTopology bool `yaml:"ClusterTopology"`
				} `yaml:"featureGates"`
			}(struct {
				FeatureGates struct {
					MachinePool     bool
					ClusterTopology bool
				} `yaml:"featureGates"`
			}{
				FeatureGates: struct {
					MachinePool     bool
					ClusterTopology bool
				}{
					MachinePool:     true,
					ClusterTopology: true,
				},
			}),
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marshaling data to yaml: %v\n", err))
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	controllerDir := filepath.Join(friggDir, vars.ControllerDir)
	newFilePath := filepath.Join(controllerDir, vars.HelmAddonProvName)

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	// Write to file
	err = os.WriteFile(newFilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing Infra provider yaml: %v\n", err))
	}
}
