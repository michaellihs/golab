## golab merge-requests get-diff-version

Get a single merge request diff version

### Synopsis


Get a single merge request diff version.

```
golab merge-requests get-diff-version [flags]
```

### Options

```
  -h, --help             help for get-diff-version
  -i, --id string        (required) The ID or URL encoded path of a project
  -m, --iid int          (required) The internal ID of the merge request
  -v, --version_id int   (required) The ID of the merge request diff version
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab merge-requests](golab_merge-requests.md)	 - Manage Merge Requests

