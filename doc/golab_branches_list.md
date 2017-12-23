## golab branches list

List repository branches

### Synopsis


Get a list of repository branches from a project, sorted by name alphabetically. This endpoint can be accessed without authentication if the repository is publicly accessible.

```
golab branches list [flags]
```

### Options

```
  -h, --help        help for list
  -i, --id string   (required) The ID or URL-encoded path of the project owned by the authenticated user
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab branches](golab_branches.md)	 - Branches

