## golab environments delete

Delete an environment

### Synopsis


Deletes an environment with a given ID.

```
golab environments delete [flags]
```

### Options

```
  -e, --environment_id int   (required) The ID of the environment
  -h, --help                 help for delete
  -i, --id string            (required) The ID or URL-encoded path of the project owned by the authenticated user
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab environments](golab_environments.md)	 - Manage environments

