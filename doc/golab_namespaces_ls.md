## golab namespaces ls

List namespaces

### Synopsis


Get a list of the namespaces of the authenticated user. If the user is an administrator, a list of all namespaces in the GitLab instance is shown.

```
golab namespaces ls [flags]
```

### Options

```
  -h, --help           help for ls
      --page int       (optional) Page of results to retrieve
      --per_page int   (optional) The number of results to include per page (max 100)
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab namespaces](golab_namespaces.md)	 - Manage namespaces

