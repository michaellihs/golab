## golab branches protect

Protect repository branch

### Synopsis


Protects a single project repository branch. This is an idempotent function, protecting an already protected repository branch still returns a 200 OK status code.

```
golab branches protect [flags]
```

### Options

```
  -b, --branch string          (required) The name of the branch
  -m, --developers_can_merge   (optional) Flag if developers can merge to the branch
  -p, --developers_can_push    (optional) Flag if developers can push to the branch
  -h, --help                   help for protect
  -i, --id string              (required) The ID or URL-encoded path of the project owned by the authenticated user
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab branches](golab_branches.md)	 - Branches

