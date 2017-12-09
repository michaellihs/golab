## golab user ssh-keys get

Single SSH key

### Synopsis


Get a single ssh key

```
golab user ssh-keys get [flags]
```

### Options

```
  -h, --help         help for get
  -k, --key_id int   (mandatory) key id of ssh key to be shown
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

