## golab project ls

List all projects

### Synopsis


Get a list of all visible projects across GitLab for the authenticated user.

```
golab project ls [flags]
```

### Options

```
      --archived                      (optional) Limit by archived status
  -h, --help                          help for ls
      --membership                    (optional) Limit by projects that the current user is a member of
      --order_by string               (optional) Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at
      --owned                         (optional) Limit by projects owned by the current user
      --page int                      (optional) Page of results to retrieve
      --per_page int                  (optional) The number of results to include per page (max 100)
      --search string                 (optional) Return list of projects matching the search criteria
      --simple                        (optional) Return only the ID, URL, name, and path of each project
      --sort string                   (optional) Return projects sorted in asc or desc order. Default is desc
      --starred                       (optional) Limit by projects starred by the current user
      --statistics                    (optional) Include project statistics
      --visibility string             (optional) Limit by visibility public, internal, or private
      --with_issues_enabled           (optional) Limit by enabled issues feature
      --with_merge_requests_enabled   (optional) Limit by enabled merge requests feature
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project](golab_project.md)	 - Manage projects

