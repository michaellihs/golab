## golab login

Login to Gitlab

### Synopsis


Login to Gitlab using username and password

```
golab login [flags]
```

### Options

```
  -h, --help              help for login
  -s, --host string       (required) Hostname (http://gitlab.my-domain.com) of the gitlab server
  -p, --password string   (optional) Password for the login
  -u, --user string       (required) Username for the login
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab](golab.md)	 - Gitlab CLI written in Go

