## tanzu cluster create

Create a cluster

```
tanzu cluster create CLUSTER_NAME [flags]
```

### Options

```
  -d, --dry-run       Does not create cluster, but show the deployment YAML instead
  -f, --file string   Configuration file from which to create a cluster
  -h, --help          help for create
      --tkr string    TanzuKubernetesRelease(TKR) to be used for creating the workload cluster
```

### Options inherited from parent commands

```
      --log-file string   Log file path
  -v, --verbose int32     Number for the log level verbosity(0-9)
```

### SEE ALSO

* [tanzu cluster](tanzu_cluster.md)	 - Kubernetes cluster operations

###### Auto generated by spf13/cobra on 10-May-2021