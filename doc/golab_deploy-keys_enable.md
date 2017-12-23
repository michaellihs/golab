## golab deploy-keys enable

Enable a deploy key

### Synopsis


Enables a deploy key for a project so this can be used. Returns the enabled key, with a status code 201 when successful.

```
golab deploy-keys enable [flags]
```

### Options

```
  -h, --help         help for enable
  -i, --id string    (required) The ID or URL-encoded path of the project owned by the authenticated user
  -k, --key_id int   (required) The ID of the deploy key
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab deploy-keys](golab_deploy-keys.md)	 - Deploy Keys API

