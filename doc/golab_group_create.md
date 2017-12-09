## golab group create

New group

### Synopsis


Creates a new project group. Available only for users who can create groups.

```
golab group create [flags]
```

### Options

```
      --description string       (optional) The group's description
  -h, --help                     help for create
      --lfs_enabled              (optional) Enable/disable Large File Storage (LFS) for the projects in this group
  -n, --name string              (required) The name of the group
      --parent_id int            (optional) The parent group id for creating nested group.
  -p, --path string              (required) The path of the group
      --request_access_enabled   (optional) - Allow users to request member access.
      --visibility string        (optional) The group's visibility. Can be private, internal, or public.
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group](golab_group.md)	 - Manage Gitlab Groups

