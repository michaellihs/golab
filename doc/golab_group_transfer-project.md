## golab group transfer-project

Transfer project to group

### Synopsis


Transfer a project to the Group namespace. Available only for admin.

```
golab group transfer-project [flags]
```

### Options

```
  -h, --help             help for transfer-project
  -i, --id string        (required) The ID or URL-encoded path of the group owned by the authenticated user
  -p, --project_id int   (required) The ID or URL-encoded path of a project
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group](golab_group.md)	 - Manage Gitlab Groups

