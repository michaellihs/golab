## golab merge-requests accept

Accept merge request

### Synopsis


Merge changes submitted with MR using this API.

If it has some conflicts and can not be merged - you'll get a 405 and the error message 'Branch cannot be merged'

If merge request is already merged or closed - you'll get a 406 and the error message 'Method Not Allowed'

If the sha parameter is passed and does not match the HEAD of the source - you'll get a 409 and the error message 'SHA does not match HEAD of source branch'

If you don't have permissions to accept this merge request - you'll get a 401

```
golab merge-requests accept [flags]
```

### Options

```
  -h, --help                           help for accept
  -i, --id string                      (required) The ID or URL-encoded path of the project owned by the authenticated user
      --merge_commit_message string    (optional) Custom merge commit message
  -m, --merge_request_iid int          (required) Internal ID of MR
      --merge_when_pipeline_succeeds   (optional) if true the MR is merged when the pipeline succeeds
      --sha string                     (optional) if present, then this SHA must match the HEAD of the source branch, otherwise the merge will fail
  -d, --should_remove_source_branch    (optional) if true removes the source branch
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab merge-requests](golab_merge-requests.md)	 - Manage Merge Requests

