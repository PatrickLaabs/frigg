# Generate your custom gitops Repositories

## Requirements

- GitHub Token
- Github Username

### Set it like
```shell
export GITHUB_TOKEN=
export GITHUB_USERNAME=
```

## Management Cluster
```shell
frigg gitops create-template mgmt-cluster --desired-repo-name "Your Repo Name"
```

## Workload Cluster
```shell
frigg gitops create-template workload-cluster --desired-repo-name "Your Repo Name"
```

Your generated Repositories will be stored directly at your Github Account.

After everything has been successfully created, your just need to tell GitHub, that the new Repository is a Template Repository.

Under your new repo, go to the settings panel a check the box, called \
`Template repository`