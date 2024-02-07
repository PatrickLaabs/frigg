package helmchartproxies

import (
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capd/helmchartproxies/cni"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capd/helmchartproxies/cnibootstrap"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capd/helmchartproxies/mgmtVault"
)

func InstallMgmtHelmCharts() {
	//argocdHubWorkload.Installation()
	//argocdWorkload.Installation()
	cni.Installation()
	//mgmtArgocdApps.Installation()
	//mgmtArgocdEvents.Installation()
	//mgmtArgocdRollouts.Installation()
	//mgmtArgocdWorkflows.Installation()
	//mgmtArgohub.Installation()
	mgmtVault.Installation()
}

func InstallBootstrapHelmCharts() {
	cnibootstrap.Installation()
}
