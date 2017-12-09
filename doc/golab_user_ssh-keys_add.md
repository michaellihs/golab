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
  -k, --key string     (mandatory) public ssh key
  -t, --title string   (mandatory) title for ssh public key
  -u, --user int       (optional) id of user to add key for
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
  -i, --id int           (optional) id of user to show ssh-keys for - if none is given, logged in user will be used
```

### SEE ALSO
* [golab user ssh-keys](golab_user_ssh-keys.md)	 - Manage a user's ssh keys

