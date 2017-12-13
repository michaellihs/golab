## golab merge-requests subscribe

Subscribe to a merge request

### Synopsis


Subscribes the authenticated user to a merge request to receive notification. If the user is already subscribed to the merge request, the status code 304 is returned.

```
golab merge-requests subscribe [flags]
```

### Options

```
  -h, --help        help for subscribe
  -i, --id string   (required) The ID or URL encoded path of a project
  -m, --iid int     (required) The internal ID of the merge request
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab merge-requests](golab_merge-requests.md)	 - Manage Merge Requests

