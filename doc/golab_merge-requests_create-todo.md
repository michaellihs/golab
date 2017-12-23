## golab merge-requests create-todo

Create a todo

### Synopsis


Manually creates a todo for the current user on a merge request. If there already exists a todo for the user on that merge request, status code 304 is returned.

```
golab merge-requests create-todo [flags]
```

### Options

```
  -h, --help        help for create-todo
  -i, --id string   (required) The ID or URL encoded path of a project
  -m, --iid int     (required) The internal ID of the merge request
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab merge-requests](golab_merge-requests.md)	 - Manage Merge Requests

