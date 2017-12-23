## golab user ssh-keys add

Add SSH key

### Synopsis


Creates a new key (owned by the currently authenticated user, if no user id was given)

```
golab user ssh-keys add [flags]
```

### Options

```
  -h, --help           help for add
  -k, --key string     (required) Public SSH key
  -t, --title string   (required) New SSH Key's title
  -u, --user string    (required) User ID or user name of user to delete SSH key from
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user ssh-keys](golab_user_ssh-keys.md)	 - Manage a user's SSH keys

