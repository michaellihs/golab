## golab personal-access-token

Create a personal access token

### Synopsis


Create a personal access token for a user identified by username and password

```
golab personal-access-token [flags]
```

### Options

```
  -a, --api                 (optional) Access the authenticated user's API (default: false)
      --expires string      (optional) Expiration date of token
  -h, --help                help for personal-access-token
  -s, --host string         (required) Hostname (http://gitlab.my-domain.com) of the gitlab server
  -p, --password string     (optional) Password for the login
      --read_registry       (optional) Grant access to the docker registry (default: false)
      --read_user           (optional) Read the authenticated user's personal information (default: false)
      --sudo                (optional) Perform API actions as any user in the system (default: false)
      --token_name string   (optional) Name of token
  -u, --username string     (required) Username for the login
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab](golab.md)	 - Gitlab CLI written in Go

