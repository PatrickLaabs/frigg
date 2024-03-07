package download

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"github.com/hashicorp/go-getter"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	FriggWorkingDir()
}

var friggDir string

func FriggWorkingDir() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	// Setting Tools directory
	friggDir = filepath.Join(homedir, vars.FriggDirName, vars.FriggTools)

	// Check if directory exists and create it if not
	if _, err := os.Stat(friggDir); os.IsNotExist(err) {
		println(color.YellowString("Tools directory does not exists, Creating %s\n", friggDir))
		err = os.MkdirAll(friggDir, 0755) // Create directory with permissions 0755
		if err != nil {
			println(color.RedString("Error creating directory %s: %v\n", friggDir, err))
		}
	} else if err != nil {
		// Handle other potential errors during stat
		println(color.RedString("Error checking directory %s: %v\n", friggDir, err))
		os.Exit(1)
	}
}

func GithubCli() {
	var operatingSystem string
	if runtime.GOOS == "darwin" {
		operatingSystem = "macOS"
	} else {
		operatingSystem = runtime.GOOS
	}

	_, err := os.Stat(filepath.Join(friggDir, "gh"))
	if err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("GH CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/cli/cli/releases/download/v" + vars.GithubCliVersion + "/gh_" + vars.GithubCliVersion + "_" + operatingSystem + "_" + runtime.GOARCH + ".zip"
			dst := filepath.Join(friggDir, "/gh_"+vars.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+".zip")

			err = getter.GetAny(dst, src)
			if err != nil {
				println(color.RedString("error on downloading github cli: %v\n", err))
			}

			ghSrcPath := filepath.Join(friggDir, "gh_"+vars.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+".zip/"+"gh_"+vars.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+"/bin/gh")
			ghDstPath := filepath.Join(friggDir, "gh")
			err = os.Link(ghSrcPath, ghDstPath)
			if err != nil {
				println(color.RedString("error on linking gh cli: %v\n", err))
			}
		} else {
			println(color.RedString("Error checking file existence of github cli: %v\n", err))
		}
	} else {
		println(color.YellowString("GH CLI already exists. Continuing.."))
	}
}

func Kubectl() {
	_, err := os.Stat(filepath.Join(friggDir, "kubectl"))
	if err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("Kubectl CLI does not exist inside tools dir. Downloading now.."))

			src := "https://dl.k8s.io/release/v" + vars.KubectlVersion + "/bin/" + runtime.GOOS + "/" + runtime.GOARCH + "/kubectl"
			dst := filepath.Join(friggDir, vars.KubectlVersion+"/bin/"+runtime.GOOS+"/"+runtime.GOARCH)

			err := getter.GetAny(dst, src)
			if err != nil {
				println(color.RedString("error on downloading kubectl cli: %v\n", err))
			}

			kubectlSrcPath := filepath.Join(friggDir, vars.KubectlVersion+"/bin/"+runtime.GOOS+"/"+runtime.GOARCH+"/kubectl")
			kubectlDstPath := filepath.Join(friggDir, "kubectl")
			err = os.Link(kubectlSrcPath, kubectlDstPath)
			if err != nil {
				println(color.RedString("error on linking kubectl: %v\n", err))
			}

			err = os.Chmod(kubectlDstPath, 0755)
			if err != nil {
				println(color.RedString("error with chmod on kubectl cli: %v\n", err))
			}
		} else {
			println(color.RedString("Error checking file existence of kubectl: %v\n", err))
		}
	} else {
		println(color.YellowString("Kubectl CLI already exists. Continuing.."))
	}
}

func Clusterctl() {
	_, err := os.Stat(filepath.Join(friggDir, "clusterctl"))
	if err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("Clusterctl CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/kubernetes-sigs/cluster-api/releases/download/v" + vars.ClusterctlVersion + "/clusterctl-" + runtime.GOOS + "-" + runtime.GOARCH
			dst := filepath.Join(friggDir, "clusterctl-directory")

			err := getter.GetAny(dst, src)
			if err != nil {
				println(color.RedString("error on downloading clusterctl cli: %v\n", err))
			}

			clusterctlSrcPath := filepath.Join(friggDir, "clusterctl-directory/"+"clusterctl-"+runtime.GOOS+"-"+runtime.GOARCH)
			clusterctlDstPath := filepath.Join(friggDir, "clusterctl")
			err = os.Link(clusterctlSrcPath, clusterctlDstPath)
			if err != nil {
				println(color.RedString("error on linking clusterctl: %v\n", err))
			}

			err = os.Chmod(clusterctlDstPath, 0755)
			if err != nil {
				println(color.RedString("error with chmod on clusterctl cli: %v\n", err))
			}
		} else {
			println(color.RedString("Error checking file existence of clusterctl: %v\n", err))
		}
	} else {
		println(color.YellowString("Clusterctl CLI already exists. Continuing.."))
	}
}

func K9s() {
	_, err := os.Stat(filepath.Join(friggDir, "k9s"))
	if err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("K9s CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/derailed/k9s/releases/download/v" + vars.K9sVersion + "/k9s_" + runtime.GOOS + "_" + runtime.GOARCH + ".tar.gz"
			dst := filepath.Join(friggDir, "k9s-"+vars.K9sVersion)

			err := getter.GetAny(dst, src)
			if err != nil {
				println(color.RedString("error on downloading k9s cli: %v\n", err))
			}

			k9sSrcPath := filepath.Join(friggDir, "k9s-"+vars.K9sVersion+"/k9s")
			k9sDstPath := filepath.Join(friggDir, "k9s")
			err = os.Link(k9sSrcPath, k9sDstPath)
			if err != nil {
				println(color.RedString("error on linking k9s: %v\n", err))
			}

			err = os.Chmod(k9sDstPath, 0755)
			if err != nil {
				println(color.RedString("error with chmod on k9s cli: %v\n", err))
			}
		} else {
			println(color.RedString("Error checking file existence of k9s: %v\n", err))
		}
	} else {
		println(color.YellowString("K9s CLI already exists. Continuing.."))
	}
}
