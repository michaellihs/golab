## golab group-members delete

Remove a member from a group or project

### Synopsis


Removes a user from a group or project.

```
golab group-members delete [flags]
```

### Options

```
  -h, --help          help for delete
  -i, --id int        (required) the id of the group to delete user from
  -u, --user_id int   (required) the id of the user to be removed from group
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group-members](golab_group-members.md)	 - Access group members

