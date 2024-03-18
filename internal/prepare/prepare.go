package prepare

import (
	"github.com/PatrickLaabs/frigg/internal/consts"
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/fatih/color"
	"github.com/hashicorp/go-getter"
	"os"
	"path/filepath"
	"runtime"
)

var (
	homedir, _    = os.UserHomeDir()
	friggDir      = filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir = filepath.Join(friggDir, vars.FriggTools)
)

func githubCli() {
	var operatingSystem string
	if runtime.GOOS == "darwin" {
		operatingSystem = "macOS"
	} else {
		operatingSystem = runtime.GOOS
	}

	if _, err := os.Stat(filepath.Join(friggToolsDir, "gh_"+consts.GithubCliVersion)); err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("GH CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/cli/cli/releases/download/v" + consts.GithubCliVersion + "/gh_" + consts.GithubCliVersion + "_" + operatingSystem + "_" + runtime.GOARCH + ".zip"
			dst := filepath.Join(friggToolsDir, "/gh_"+consts.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+".zip")

			if err = getter.GetAny(dst, src); err != nil {
				println(color.RedString("error on downloading github cli: %v\n", err))
			}

			ghSrcPath := filepath.Join(friggToolsDir, "gh_"+consts.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+".zip/"+"gh_"+consts.GithubCliVersion+"_"+operatingSystem+"_"+runtime.GOARCH+"/bin/gh")
			ghDstPath := filepath.Join(friggToolsDir, "gh_"+consts.GithubCliVersion)

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

func kubectl() {
	if _, err := os.Stat(filepath.Join(friggToolsDir, "kubectl_"+consts.KubectlVersion)); err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("Kubectl CLI does not exist inside tools dir. Downloading now.."))

			src := "https://dl.k8s.io/release/v" + consts.KubectlVersion + "/bin/" + runtime.GOOS + "/" + runtime.GOARCH + "/kubectl"
			dst := filepath.Join(friggToolsDir, consts.KubectlVersion+"/bin/"+runtime.GOOS+"/"+runtime.GOARCH)

			if err := getter.GetAny(dst, src); err != nil {
				println(color.RedString("error on downloading kubectl cli: %v\n", err))
			}

			kubectlSrcPath := filepath.Join(friggToolsDir, consts.KubectlVersion+"/bin/"+runtime.GOOS+"/"+runtime.GOARCH+"/kubectl")
			kubectlDstPath := filepath.Join(friggToolsDir, "kubectl_"+consts.KubectlVersion)

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

func clusterctl() {
	if _, err := os.Stat(filepath.Join(friggToolsDir, "clusterctl_"+consts.ClusterctlVersion)); err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("Clusterctl CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/kubernetes-sigs/cluster-api/releases/download/v" + consts.ClusterctlVersion + "/clusterctl-" + runtime.GOOS + "-" + runtime.GOARCH
			dst := filepath.Join(friggToolsDir, "clusterctl-directory")

			if err := getter.GetAny(dst, src); err != nil {
				println(color.RedString("error on downloading clusterctl cli: %v\n", err))
			}

			clusterctlSrcPath := filepath.Join(friggToolsDir, "clusterctl-directory/"+"clusterctl-"+runtime.GOOS+"-"+runtime.GOARCH)
			clusterctlDstPath := filepath.Join(friggToolsDir, "clusterctl_"+consts.ClusterctlVersion)

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

func k9s() {
	if _, err := os.Stat(filepath.Join(friggToolsDir, "k9s_"+consts.K9sVersion)); err != nil {
		if os.IsNotExist(err) {
			println(color.YellowString("K9s CLI does not exist inside tools dir. Downloading now.."))

			src := "https://github.com/derailed/k9s/releases/download/v" + consts.K9sVersion + "/k9s_" + runtime.GOOS + "_" + runtime.GOARCH + ".tar.gz"
			dst := filepath.Join(friggToolsDir, "k9s_"+consts.K9sVersion)

			if err := getter.GetAny(dst, src); err != nil {
				println(color.RedString("error on downloading k9s cli: %v\n", err))
			}

			k9sSrcPath := filepath.Join(friggToolsDir, "k9s_"+consts.K9sVersion)
			k9sDstPath := filepath.Join(friggToolsDir, "k9s_"+consts.K9sVersion)

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

// Tools downloads the cli tools needed for frigg to operate correctly.
func Tools() {
	githubCli()
	kubectl()
	clusterctl()
	k9s()
}
