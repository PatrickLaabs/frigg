# Running Frigg for your local development

## Get started

You will need to set these Environment variables:

```shell
export GITHUB_TOKEN=
export GITHUB_USERNAME=
export GITHUB_MAIL=
```

### Create your management cluster

```shell
frigg bootstrap capd-controller cluster
```

### Create your workload cluster

This will create a workload cluster ontop of your management cluster.

```shell
frigg bootstrap capd-controller workloadcluster
```

---

## GitOps Repositories

By default, you will use the standard template repositories for your management \
and workload clusters.

### Friggs Management-Cluster GitOps Template Repository
This Repository will the beating-heart for your management-cluster.\
[friggs-mgmt-repo-template](https://github.com/PatrickLaabs/friggs-mgmt-repo-template)

### Friggs Workload-Clusters GitOps Template Repository
This Repository is used to provide a good foundation for your workload clusters with various deployments.\
[friggs-workload-repo-template](https://github.com/PatrickLaabs/friggs-workload-repo-template)

### Use your custom GitOps Repo Template

Frigg provides the ability to let you define your custom GitOps Repository Template URL
while bootstrapping your environment.

You can either set your GitOps repo for your management cluster:
```shell
frigg bootstrap capd-controller cluster --gitops-template-repo <ORG|USERNAME/Repo>
```

or for your Workload Clusters:
```shell
frigg bootstrap capd-controller cluster --gitops-workload-template-repo <ORG|USERNAME/Repo>
```

If you want to set both, you are free to do so:
```shell
frigg bootstrap capd-controller cluster --gitops-template-repo <ORG|USERNAME/Repo> --gitops-workload-template-repo <ORG|USERNAME/Repo>
```

### Creating your custom GitOps Repositories

This is currently under development.\
PR: #128