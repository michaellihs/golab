## golab personal-access-token

Create a personal access token

### Synopsis


Create a personal access token for a user identified by username and password

```
golab personal-access-token [flags]
```

### Options

```
  -h, --help              help for personal-access-token
  -s, --host string       (required) Hostname (http://gitlab.my-domain.com) of the gitlab server
  -p, --password string   (optional) Password for the login
  -u, --username string   (required) Username for the login
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab](golab.md)	 - Gitlab CLI written in Go

