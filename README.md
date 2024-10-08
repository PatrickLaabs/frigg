# ARCHIVED
Frigg will be refactored and moves into a new home: [eros](https://github.com/PatrickLaabs/eros)

# Frigg

[![Go Reference](https://pkg.go.dev/badge/github.com/PatrickLaabs/frigg.svg)](https://pkg.go.dev/github.com/PatrickLaabs/frigg)
[![Go Report Card](https://goreportcard.com/badge/github.com/PatrickLaabs/frigg)](https://goreportcard.com/badge/github.com/PatrickLaabs/frigg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


## What is Frigg

***Meaning of Frigg**:* Goddess of wisdom and crafts

![Frigg](assets/Frigg_Logo.png)

### **TL;DR**:
With Frigg, you provision **N-Kubernets** Clusters, which are **GitOps**-enabled and have **batteries included**.

Frigg is a cli project, to easily create one to one hundred 
of Kubernetes clusters on different hyperscalers.

**Since we relay heavenly on Cluster-API, we could implement any supported provider to _Frigg_\
[Check the supported Hyperscalers](https://cluster-api.sigs.k8s.io/reference/providers)**

No matter which hyperscaler you choose, your kubernetes clusters will be
attached to one another, and are also GitOps enabled.

At the end, you will have N-amount of clusters, with a Github
repository for each of them, where you are able to add more applications deployments.

## Quick overview of Frigg()
![Friggs Overview](assets/Frigg_Overview.png)

## Usage
### Requirements:

- Docker Desktop / Engine\
  (Podman (is currently not yet fully supported))

#### Get the *frigg* cli:
Get the binary using go:
```
go install github.com/PatrickLaabs/frigg@1.1.2
```

```
curl -L -o frigg.tar.gz https://github.com/PatrickLaabs/frigg/releases/download/1.1.2/frigg_1.1.2_darwin_arm64.tar.gz
tar -xf frigg.tar.gz
chmod +x frigg
./frigg version
```

or download the binary at the releasepage:\
[Frigg - Github Release Page](https://github.com/PatrickLaabs/frigg/releases)

Homebrew is on the way.

#### Start the deployment:

While everything gets bootstrapped and provisioned, the Frigg CLI \
also creates a working directory inside your home directory at
`$HOME/.frigg`

Inside this directory we will store every generated file, such as\
the private and public ssh keypairs, various manifests, etc.
<br></br>

**Set the environment variables:**

```
export GITHUB_TOKEN=
export GITHUB_USERNAME=
export GITHUB_MAIL=
```

**Running Frigg:**

```
frigg provisions kubernetes cluster with capi and gitops in no-time

Usage:
  frigg [command]

Available Commands:
  bootstrap   bootstrap various clusters on different providers
  completion  Generate the autocompletion script for the specified shell
  delete      Deletes one of [cluster]
  help        Help about any command
  version     Prints the frigg CLI version
```

**Provision your first management cluster:**
``` shell
frigg bootstrap capd cluster
```
or - to use the ClusterAPI Controllers - run:
```shell
frigg bootstrap capd-controller cluster
```
This might take a while, since we are doing some heavy lifting.

**Provision your workload cluster ontop:**
``` sh
frigg bootstrap capd workloadcluster

or

frigg bootstrap capd-controller workloadcluster
```

After the provisioning of your management cluster is ready,\
you can port-forward the argocd-server pod and login with:
```
User: admin
Password: frigg
```
## Features

### Supported Providers
- vCluster
- CAPD (Docker)

### Providers under development
- vSphere
- Azure
- Google
- Harvester
- Proxmox

## Documentation

Further documentation is available in the `/docs` directory.

## Relevant Repositories

### Friggs Management-Cluster GitOps Repository Template
This Repository will the beating-heart for your management-cluster.\
[friggs-mgmt-repo-template](https://github.com/PatrickLaabs/friggs-mgmt-repo-template)

### Friggs Workload-Clusters GitOps Repository Template
This Repository is used to provide a good foundation for your workload clusters with various deployments.\
[friggs-workload-repo-template](https://github.com/PatrickLaabs/friggs-workload-repo-template)

### Frigg-Builder
This simple Tool will be used by ArgoCD Workflows to operate various things inside the pipeline.\
[Frigg-Builder](https://github.com/PatrickLaabs/frigg-builder)

### Sample-Application
An example application deployment for your workload clusters.\
[sample-application](https://github.com/PatrickLaabs/sample-application)

## Support

You like the project, and want to support further development?
Glad to hear!

<a href="https://www.buymeacoffee.com/patricklaabs" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>

Thank you very much, for supporting me 🚀
