## golab branches unprotect

Unprotect repository branch

### Synopsis


Unprotects a single project repository branch. This is an idempotent function, unprotecting an already unprotected repository branch still returns a 200 OK status code.

```
golab branches unprotect [flags]
```

### Options

```
  -b, --branch string   (required) The name of the branch
  -h, --help            help for unprotect
  -i, --id string       (required) The ID or URL-encoded path of the project
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab branches](golab_branches.md)	 - Branches

