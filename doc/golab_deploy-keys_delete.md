## golab deploy-keys delete

Delete deploy key

### Synopsis


Removes a deploy key from the project. If the deploy key is used only for this project, it will be deleted from the system.

```
golab deploy-keys delete [flags]
```

### Options

```
  -h, --help         help for delete
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

