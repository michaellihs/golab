## golab environments create

Create a new environment

### Synopsis


Creates a new environment for the given project with the given name and external URL.

```
golab environments create [flags]
```

### Options

```
      --external_url string   (optional) Place to link to for this environment
  -h, --help                  help for create
  -i, --id string             (required) The ID or URL-encoded path of the project owned by the authenticated user
      --name string           (required) The name of the environment
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab environments](golab_environments.md)	 - Manage environments

