## golab merge-requests update

Update merge request

### Synopsis


Updates an existing merge request. You can change the target branch, title, or even close the MR.

```
golab merge-requests update [flags]
```

### Options

```
      --assignee_id int         (optional) Assignee user ID
      --description string      (optional) Description of MR
      --discussion_locked       (optional) Flag indicating if the merge request's discussion is locked. If the discussion is locked only project members can add, edit or resolve comments.
  -h, --help                    help for update
  -i, --id string               (required) The ID or URL-encoded path of the project owned by the authenticated user
      --labels string           (optional) Labels for MR as a comma-separated list
  -m, --merge_request_iid int   (required) The ID of a merge request
      --milestone_id int        (optional) The ID of a milestone
      --remove_source_branch    (optional) Flag indicating if a merge request should remove the source branch when merging
      --state_event string      (optional) New state (close/reopen)
      --target_branch string    (optional) The target branch
      --title string            (optional) Title of MR
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab merge-requests](golab_merge-requests.md)	 - Manage Merge Requests

