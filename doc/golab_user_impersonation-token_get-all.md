## golab user impersonation-token get-all

Get all impersonation tokens of a user

### Synopsis


It retrieves every impersonation token of the user.

```
golab user impersonation-token get-all [flags]
```

### Options

```
  -h, --help             help for get-all
  -s, --state string     (optional) filter tokens based on state (all, active, inactive)
  -u, --user_id string   (required) The ID of the user or the name of the user to get tokens for
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user impersonation-token](golab_user_impersonation-token.md)	 - Impersonation token

