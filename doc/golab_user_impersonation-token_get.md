## golab user impersonation-token get

Get an impersonation token of a user

### Synopsis


It shows a user's impersonation token (admins only).

```
golab user impersonation-token get [flags]
```

### Options

```
  -h, --help                         help for get
  -t, --impersonation_token_id int   (required) The ID of the impersonation token
  -u, --user_id string               (required) The ID of the user or the username for which to get a token
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user impersonation-token](golab_user_impersonation-token.md)	 - Impersonation token

