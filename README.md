# ArgoHub - Batteries included

## Goals

- Enable a user to easily provision the ArgoHub instance on kind
- Bootstrap the GitOps repository, with the given credentials
- CLI Tool
- Built-in Server to bootstrap everything

## Getting started

- docker
- kind

```
Usage:
  argohub [command]

Available Commands:
  build       Build one of [node-image]
  completion  Output shell completion code for the specified shell (bash, zsh or fish)
  create      Creates one of [cluster]
  delete      Deletes one of [cluster]

```

## Supported Providers

- vCluster

### Under Development
- vSphere
- Azure
- Google