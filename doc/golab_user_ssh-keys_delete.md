## golab user ssh-keys delete

Delete SSH key

### Synopsis


Deletes key owned by a specified user (available only for admin) or by currently logged in user.

```
golab user ssh-keys delete [flags]
```

### Options

```
  -h, --help          help for delete
  -k, --key_id int    (required) key id of SSH key to be deleted
  -u, --user string   (required) User ID or user name of user to delete SSH key from
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user ssh-keys](golab_user_ssh-keys.md)	 - Manage a user's SSH keys

