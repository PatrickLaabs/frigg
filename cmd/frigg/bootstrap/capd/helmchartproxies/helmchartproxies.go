package helmchartproxies

import (
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/argocdWorkload"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/cni"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/cnibootstrap"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgocdApps"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgocdEvents"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgocdRollouts"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgocdWorkflows"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgohub"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtVault"
)

func InstallMgmtHelmCharts() {
	//argocdHubWorkload.Installation()
	argocdWorkload.Installation()
	cni.Installation()
	mgmtArgocdApps.Installation()
	mgmtArgocdEvents.Installation()
	mgmtArgocdRollouts.Installation()
	mgmtArgocdWorkflows.Installation()
	mgmtArgohub.Installation()
	mgmtVault.Installation()
}

func InstallBootstrapHelmCharts() {
	cnibootstrap.Installation()
}
