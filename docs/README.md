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

- KinD
- Docker
- clusterctl
- kubectl

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