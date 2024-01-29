# ArgoHub - Batteries included

## Goals

- Enable a user to easily provision the ArgoHub instance on k3d
- Bootstrap the GitOps repository, with the given credentials
- CLI Tool
- Built-in Server to bootstrap everything

## Getting started

- docker
- k3d

```
go run main.go -h

Available Commands:
  bootstrap   A brief description of your command
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     A brief description of your command
```

## Supported Providers

- vCluster

### Under Development
- vSphere
- Azure
- Google