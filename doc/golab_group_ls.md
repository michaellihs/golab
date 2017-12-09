## golab group ls

List groups

### Synopsis


Get a list of visible groups for the authenticated user.

```
golab group ls [flags]
```

### Options

```
      --all_available             (optional) Show all the groups you have access to (defaults to false for authenticated users)
  -h, --help                      help for ls
      --order_by string           (optional) Order groups by name or path. Default is name
      --owned                     (optional) Limit to groups owned by the current user
      --search string             (optional) Return the list of authorized groups matching the search criteria
      --skip_groups stringArray   (optional) Skip the group IDs passed
      --sort string               (optional) Order groups in asc or desc order. Default is asc
      --statistics                (optional) Include group statistics (admins only)
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group](golab_group.md)	 - Manage Gitlab Groups

