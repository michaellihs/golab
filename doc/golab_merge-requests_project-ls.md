## golab merge-requests project-ls

List project merge requests

### Synopsis


Get all merge requests for this project. The state parameter can be used to get only merge requests with a given state (opened, closed, or merged) or all of them (all). The pagination parameters page and per_page can be used to restrict the list of merge requests.

```
golab merge-requests project-ls [flags]
```

### Options

```
      --assignee_id int            (optional) Returns merge requests assigned to the given user id (Introduced in GitLab 9.5)
      --author_id int              (optional) Returns merge requests created by the given user id (Introduced in GitLab 9.5)
      --created_after string       (optional) Return merge requests created after the given time (inclusive)
      --created_before string      (optional) Return merge requests created before the given time (inclusive)
  -h, --help                       help for project-ls
      --id int                     (required) The ID of a project
      --iids stringArray           (optional) Return the request having the given iid
      --labels stringArray         (optional) Return merge requests matching a comma separated list of labels
      --milestone string           (optional) Return merge requests for a specific milestone
      --my_reaction_emoji string   (optional) Return merge requests reacted by the authenticated user by the given emoji (Introduced in GitLab 10.0)
      --order_by string            (optional) Return requests ordered by created_at or updated_at fields. Default is created_at
      --scope string               (optional) Return merge requests for the given scope: created-by-me, assigned-to-me or all (Introduced in GitLab 9.5)
      --sort string                (optional) Return requests sorted in asc or desc order. Default is desc
      --state string               (optional) Return all merge requests or just those that are opened, closed, or merged
      --view string                (optional) If simple, returns the iid, URL, title, description, and basic state of merge request
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab merge-requests](golab_merge-requests.md)	 - Manage Merge Requests

