## golab protected-branches protect-branch

Protect repository branches

### Synopsis


Protects a single repository branch or several project repository branches using a wildcard protected branch.

Access Levels:

0  => No access
30 => Developer access
40 => Master access


```
golab protected-branches protect-branch [flags]
```

### Options

```
  -h, --help                        help for protect-branch
      --id string                   (required) The ID or URL-encoded path of the project owned by the authenticated user
      --merge_access_level string   (optional) Access levels allowed to merge (defaults: 40, master access level)
      --name string                 (required) The name of the branch or wildcard
      --push_access_level string    (optional) Access levels allowed to push (defaults: 40, master access level)
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab protected-branches](golab_protected-branches.md)	 - Protected branches

