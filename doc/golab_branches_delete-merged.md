## golab branches delete-merged

Delete merged branches

### Synopsis


Will delete all branches that are merged into the project's default branch.

Protected branches will not be deleted as part of this operation.

```
golab branches delete-merged [flags]
```

### Options

```
  -h, --help        help for delete-merged
  -i, --id string   (required) The ID or URL-encoded path of the project
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab branches](golab_branches.md)	 - Branches

