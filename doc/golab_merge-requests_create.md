## golab merge-requests create

Create merge request

### Synopsis


Creates a new merge request.

```
golab merge-requests create [flags]
```

### Options

```
  -a, --assignee_id int         (optional) Assignee user ID
  -d, --description string      (optional) Description of MR
  -h, --help                    help for create
  -i, --id string               (required) The ID or URL-encoded path of the project owned by the authenticated user
      --labels string           (optional) Labels for MR as a comma-separated list
      --milestone_id int        (optional) The ID of a milestone
      --remove_source_branch    (optional) Flag indicating if a merge request should remove the source branch when merging
  -s, --source_branch string    (required) The source branch
  -t, --target_branch string    (required) The target branch
      --target_project_id int   (optional) The target project (numeric id)
  -n, --title string            (required) Title of MR
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab merge-requests](golab_merge-requests.md)	 - Manage Merge Requests

