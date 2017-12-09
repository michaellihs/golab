## golab group-members sync

Synchronizes members of 2 groups

### Synopsis


Synchronizes the members of 2 groups, by either

* merging them (default) - members that exist in target group but not in source group are kept
* removing them (--remove) - members that exist in target group but not in source group are deleted

```
golab group-members sync [flags]
```

### Options

```
  -h, --help         help for sync
  -r, --remove       (optional) remove members in target group that don't exist in source group
  -s, --source int   (required) id of group to copy members from
  -t, --target int   (required) id of group to copy members to
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group-members](golab_group-members.md)	 - Access group members

