## golab user impersonation-token revoke

Revoke an impersonation token (admin only)

### Synopsis


It revokes an impersonation token. Requires admin permissions.

```
golab user impersonation-token revoke [flags]
```

### Options

```
  -h, --help                         help for revoke
  -t, --impersonation_token_id int   (required) The ID of the impersonation token
  -u, --user_id string               (required) The ID of the user or username of user to revoke token for
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user impersonation-token](golab_user_impersonation-token.md)	 - Impersonation token

