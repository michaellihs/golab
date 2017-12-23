## golab merge-requests reset-spent-time

Reset spent time for a merge request

### Synopsis


Resets the total spent time for this merge request to 0 seconds.

```
golab merge-requests reset-spent-time [flags]
```

### Options

```
  -h, --help        help for reset-spent-time
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

