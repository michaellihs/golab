## golab deploy-keys add

Add deploy key

### Synopsis


Creates a new deploy key for a project.

If the deploy key already exists in another project, it will be joined to current project only if original one is accessible by the same user.

```
golab deploy-keys add [flags]
```

### Options

```
  -p, --can_push       (optional) Can deploy key push to the project's repository
  -h, --help           help for add
  -i, --id string      (required) The ID or URL-encoded path of the project owned by the authenticated user
  -k, --key string     (required) New deploy key
  -t, --title string   (required) New deploy key's title
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab deploy-keys](golab_deploy-keys.md)	 - Deploy Keys API

