## tanzu package repository delete

Delete a repository

```
tanzu package repository delete REPOSITORY_NAME [flags]
```

### Examples

```

    # Delete a repository in specified namespace 	
    tanzu package repository delete repo --namespace test-ns
```

### Options

```
  -f, --force   Force deletion of the repository
  -h, --help    help for delete
```

### Options inherited from parent commands

```
      --kubeconfig string   The path to the kubeconfig file, optional
      --log-file string     Log file path
  -n, --namespace string    Namespace for repository command, optional (default "default")
      --verbose int32       Number for the log level verbosity(0-9)
```

### SEE ALSO

* [tanzu package repository](tanzu_package_repository.md)	 - Repository operations

###### Auto generated by spf13/cobra on 15-Jul-2021
