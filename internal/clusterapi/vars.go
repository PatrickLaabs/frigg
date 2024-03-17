package clusterapi

import (
	"github.com/PatrickLaabs/frigg/internal/consts"
	"github.com/PatrickLaabs/frigg/internal/vars"
	"os"
	"path/filepath"
)

var (
	kubectl            = "kubectl_" + consts.KubectlVersion
	clusterctl         = "clusterctl_" + consts.ClusterctlVersion
	homedir, _         = os.UserHomeDir()
	friggDir           = filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir      = filepath.Join(friggDir, vars.FriggTools)
	kubectlPath        = filepath.Join(friggToolsDir, kubectl)
	friggControllerDir = filepath.Join(friggDir, vars.ControllerDir)
	clusterctlPath     = filepath.Join(friggToolsDir, clusterctl)
	clusterconfigPath  = filepath.Join(friggDir, vars.ClusterctlConfigName)
)
