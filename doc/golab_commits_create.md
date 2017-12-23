## golab commits create

Create a commit with multiple files and actions

### Synopsis


Create a commit by posting a JSON payload

JSON encoded Actions:

	[
		{
		  "action": "create",
		  "file_path": "foo/bar",
		  "content": "some content"
		},
		{
		  "action": "delete",
		  "file_path": "foo/bar2"
		},
		{
		  "action": "move",
		  "file_path": "foo/bar3",
		  "previous_path": "foo/bar4",
		  "content": "some content"
		},
		{
		  "action": "update",
		  "file_path": "foo/bar5",
		  "content": "new content"
		}
    ]

```
golab commits create [flags]
```

### Options

```
      --actions string          (required) A JSON encoded array of action hashes to commit as a batch.
      --author_email string     (optional) Specify the commit author's email address
      --author_name string      (optional) Specify the commit author's name
      --branch string           (required) Name of the branch to commit into. To create a new branch, also provide start_branch.
      --commit_message string   (required) Commit message
  -h, --help                    help for create
      --id string               (required) The ID or URL-encoded path of the project
      --start_branch string     (optional) Name of the branch to start the new commit from
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab commits](golab_commits.md)	 - Manage Commits

