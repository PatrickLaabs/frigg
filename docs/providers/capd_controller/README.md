# CAPD (ClusterAPI Provider for Docker) Documentation

## Get started

You will need to pass two Environment variables:
- GITHUB_TOKEN
- GITHUB_USERNAME
- GITHUB_MAIL

```
export GITHUB_TOKEN=
export GITHUB_USERNAME=
export GITHUB_MAIL=
```

## TL;DR:

Create a Management Cluster:

`frigg bootstrap capd-controller cluster`

Create a Workload Cluster:

`frigg bootstrap capd-controller workloadcluster`



Updated on: 15 Mar 2024