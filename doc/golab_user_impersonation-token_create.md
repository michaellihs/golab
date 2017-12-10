## golab user impersonation-token create

Create an impersonation token (admin only)

### Synopsis


It creates a new impersonation token. Note that only administrators can do this. You are only able to create impersonation tokens to impersonate the user and perform both API calls and Git reads and writes. The user will not see these tokens in their profile settings page. Requires admin permissions.

```
golab user impersonation-token create [flags]
```

### Options

```
  -e, --expires_at string    (optional) The expiration date of the impersonation token in ISO format (YYYY-MM-DD)
  -h, --help                 help for create
  -n, --name string          (required) The name of the impersonation token
  -s, --scopes stringArray   (required) The array of scopes of the impersonation token (api, read_user)
  -u, --user_id string       (required) The ID of the user
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user impersonation-token](golab_user_impersonation-token.md)	 - Impersonation token

