## golab project share

Share project with group

### Synopsis


Allow to share project with group.

```
golab project share [flags]
```

### Options

```
  -e, --expires_at string     (optional) Share expiration date in ISO 8601 format: 2016-09-26
  -a, --group_access string   (required) The permissions level to grant the group
  -g, --group_id int          (required) The ID of the group to share with
  -h, --help                  help for share
  -i, --id string             (required) The ID or URL-encoded path of the project
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project](golab_project.md)	 - Manage projects

