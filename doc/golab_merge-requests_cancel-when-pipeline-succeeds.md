## golab merge-requests cancel-when-pipeline-succeeds

Cancel Merge When Pipeline Succeeds

### Synopsis


If you don't have permissions to accept this merge request - you'll get a 401

If the merge request is already merged or closed - you get 405 and error message 'Method Not Allowed'

In case the merge request is not set to be merged when the pipeline succeeds, you'll also get a 406 error.

```
golab merge-requests cancel-when-pipeline-succeeds [flags]
```

### Options

```
  -h, --help        help for cancel-when-pipeline-succeeds
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

