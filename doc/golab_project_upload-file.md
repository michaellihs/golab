## golab project upload-file

Upload a file

### Synopsis


Uploads a file to the specified project to be used in an issue or merge request description, or a comment.

```
golab project upload-file [flags]
```

### Options

```
  -f, --file string   (required) Path to the file to be uploaded
  -h, --help          help for upload-file
  -i, --id string     (required) The ID or URL-encoded path of the project
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project](golab_project.md)	 - Manage projects

