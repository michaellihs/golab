## golab branches get

Get single repository branch

### Synopsis


Get a single project repository branch. This endpoint can be accessed without authentication if the repository is publicly accessible.

```
golab branches get [flags]
```

### Options

```
  -b, --branch string   (required) The name of the branch
  -h, --help            help for get
  -i, --id string       (required) The ID or URL-encoded path of the project owned by the authenticated user
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab branches](golab_branches.md)	 - Branches

