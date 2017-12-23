## golab project search

Search for projects by name

### Synopsis


Search for projects by name which are accessible to the authenticated user. This endpoint can be accessed without authentication if the project is publicly accessible.

```
golab project search [flags]
```

### Options

```
  -h, --help              help for search
      --order_by string   (optional) Return requests ordered by id, name, created_at or last_activity_at fields
  -s, --search string     (required) A string contained in the project name
      --sort string       (optional) Return requests sorted in asc or desc order
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project](golab_project.md)	 - Manage projects

