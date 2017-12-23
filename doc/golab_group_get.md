## golab group get

Details of a group

### Synopsis


Get all details of a group. This command can be accessed without authentication if the group is publicly accessible.

```
golab group get [flags]
```

### Options

```
  -h, --help        help for get
      --id string   (required) The ID or URL-encoded path of the group owned by the authenticated user
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group](golab_group.md)	 - Manage Gitlab Groups

