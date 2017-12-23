## golab group-members add

Add a member to a group

### Synopsis


Add a member to a group

  Access Levels:

	10 = Guest Permissions
	20 = Reporter Permissions
	30 = Developer Permissions
	40 = Master Permissions
	50 = Owner Permissions

```
golab group-members add [flags]
```

### Options

```
  -a, --access_level int    (required) access level of new group member
  -e, --expires_at string   (optional) expiry date of membership (yyyy-mm-dd)
  -h, --help                help for add
  -i, --id int              (required) id of group to add new member to
  -u, --user_id int         (required) id of user to be added as new group member
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group-members](golab_group-members.md)	 - Access group members

