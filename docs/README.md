# Frigg documentation

## Content

- ./frigg
- ./providers
  - ./capd
  - ./capv
  - ./capz
  - ./harvester

## How to get started

You will need to have the following tools installed:

(Currently, we are working on an implementation, 
to download the needed binaries - so that you do not need to care about.)

- [kind](https://formulae.brew.sh/formula/kind#default)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [clusterctl](https://formulae.brew.sh/formula/clusterctl#default)
- [kubectl](https://formulae.brew.sh/formula/kubernetes-cli#default)
- [github cli](https://formulae.brew.sh/formula/gh#default)
- [jq](https://formulae.brew.sh/formula/jq#default)

## Choose on of the providers you'd like to use

We currently support the following ClusterAPI Providers:

- CAPD (ClusterAPI Provider Docker) for local development
- CAPV (ClusterAPI Provider vSphere)
- CAPZ (ClusterAPI Provider Azure)
- CAPHV (ClusterAPI Provider Harvester)
- CAPVC (ClusterAPI Provider vCluster)
- CABPT (ClusterAPI Bootstrap Provider Talos)

We will provide hands-on documentations, on each provider, so that
you can easily get started on using those.

For now: You might want to get started with CAPD, in order to get a feeling how `frigg`
works.

## Hardware recommendations

Depending on your Provider, the hardware recommendations may vary.

Please take a look inside the documentation on each provider.
We provide a minimum and a recommended hardware spec, for each of those.