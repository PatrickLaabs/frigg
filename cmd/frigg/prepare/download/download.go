package download

import (
	"github.com/PatrickLaabs/frigg/pkg/consts"
	"github.com/PatrickLaabs/frigg/pkg/toolsdir"
	"github.com/fatih/color"
	"github.com/hashicorp/go-getter"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	toolsdir.FriggWorkingDir()
}

var friggDir string

//func FriggWorkingDir() {
//	homedir, err := os.UserHomeDir()
//	if err != nil {
//		println(color.RedString("error on accessing home directory: %v\n", err))
//	}
//
//	// Setting Tools directory
//	friggDir = filepath.Join(homedir, vars.FriggDirName, vars.FriggTools)
//
//	// Check if directory exists and create it if not
//	if _, err := os.Stat(friggDir); os.IsNotExist(err) {
//		println(color.YellowString("Tools directory does not exists, Creating %s\n", friggDir))
//		if err = os.MkdirAll(friggDir, 0755); err != nil {
//			println(color.RedString("Error creating directory %s: %v\n", friggDir, err))
//		}
//	} else if err != nil {
//		// Handle other potential errors during stat
//		println(color.RedString("Error checking directory %s: %v\n", friggDir, err))
//		os.Exit(1)
//	}
//}

func GithubCli() {
	var operatingSystem string
	if runtime.GOOS == "darwin" {
		operatingSystem = "macOS"
	} else {
		operatingSystem = runtime.GOOS
	}

	if _, err := os.Stat(filepath.Join(toolsdir.FriggDir, "gh_"+consts.GithubCliVersion)); err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("GH CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/cli/cli/releases/download/v" + consts.GithubCliVersion + "/gh_" + consts.GithubCliVersion + "_" + operatingSystem + "_" + runtime.GOARCH + ".zip"
			dst := filepath.Join(toolsdir.FriggDir, "/gh_"+consts.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+".zip")

			if err = getter.GetAny(dst, src); err != nil {
				println(color.RedString("error on downloading github cli: %v\n", err))
			}

			ghSrcPath := filepath.Join(toolsdir.FriggDir, "gh_"+consts.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+".zip/"+"gh_"+consts.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+"/bin/gh")
			ghDstPath := filepath.Join(toolsdir.FriggDir, "gh_"+consts.GithubCliVersion)

			if err = os.Link(ghSrcPath, ghDstPath); err != nil {
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
	if _, err := os.Stat(filepath.Join(toolsdir.FriggDir, "kubectl_"+consts.KubectlVersion)); err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("Kubectl CLI does not exist inside tools dir. Downloading now.."))

			src := "https://dl.k8s.io/release/v" + consts.KubectlVersion + "/bin/" + runtime.GOOS + "/" + runtime.GOARCH + "/kubectl"
			dst := filepath.Join(toolsdir.FriggDir, consts.KubectlVersion+"/bin/"+runtime.GOOS+"/"+runtime.GOARCH)

			if err := getter.GetAny(dst, src); err != nil {
				println(color.RedString("error on downloading kubectl cli: %v\n", err))
			}

			kubectlSrcPath := filepath.Join(toolsdir.FriggDir, consts.KubectlVersion+"/bin/"+runtime.GOOS+"/"+runtime.GOARCH+"/kubectl")
			kubectlDstPath := filepath.Join(toolsdir.FriggDir, "kubectl_"+consts.KubectlVersion)

			if err = os.Link(kubectlSrcPath, kubectlDstPath); err != nil {
				println(color.RedString("error on linking kubectl: %v\n", err))
			}

			if err = os.Chmod(kubectlDstPath, 0755); err != nil {
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
	if _, err := os.Stat(filepath.Join(toolsdir.FriggDir, "clusterctl_"+consts.ClusterctlVersion)); err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("Clusterctl CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/kubernetes-sigs/cluster-api/releases/download/v" + consts.ClusterctlVersion + "/clusterctl-" + runtime.GOOS + "-" + runtime.GOARCH
			dst := filepath.Join(toolsdir.FriggDir, "clusterctl-directory")

			if err := getter.GetAny(dst, src); err != nil {
				println(color.RedString("error on downloading clusterctl cli: %v\n", err))
			}

			clusterctlSrcPath := filepath.Join(toolsdir.FriggDir, "clusterctl-directory/"+"clusterctl-"+runtime.GOOS+"-"+runtime.GOARCH)
			clusterctlDstPath := filepath.Join(toolsdir.FriggDir, "clusterctl_"+consts.ClusterctlVersion)

			if err = os.Link(clusterctlSrcPath, clusterctlDstPath); err != nil {
				println(color.RedString("error on linking clusterctl: %v\n", err))
			}

			if err = os.Chmod(clusterctlDstPath, 0755); err != nil {
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
	if _, err := os.Stat(filepath.Join(toolsdir.FriggDir, "k9s")); err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("K9s CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/derailed/k9s/releases/download/v" + consts.K9sVersion + "/k9s_" + runtime.GOOS + "_" + runtime.GOARCH + ".tar.gz"
			dst := filepath.Join(toolsdir.FriggDir, "k9s-"+consts.K9sVersion)

			if err := getter.GetAny(dst, src); err != nil {
				println(color.RedString("error on downloading k9s cli: %v\n", err))
			}

			k9sSrcPath := filepath.Join(toolsdir.FriggDir, "k9s-"+consts.K9sVersion+"/k9s")
			k9sDstPath := filepath.Join(toolsdir.FriggDir, "k9s_"+consts.K9sVersion)

			if err := os.Link(k9sSrcPath, k9sDstPath); err != nil {
				println(color.RedString("error on linking k9s: %v\n", err))
			}

			if err := os.Chmod(k9sDstPath, 0755); err != nil {
				println(color.RedString("error with chmod on k9s cli: %v\n", err))
			}
		} else {
			println(color.RedString("Error checking file existence of k9s: %v\n", err))
		}
	} else {
		println(color.YellowString("K9s CLI already exists. Continuing.."))
	}
}
