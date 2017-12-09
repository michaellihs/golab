## golab project archive

Archive a project

### Synopsis


Archives the project if the user is either admin or the project owner of this project. This action is idempotent, thus archiving an already archived project will not change the project.

```
golab project archive [flags]
```

### Options

```
  -h, --help        help for archive
  -i, --id string   (required) The ID or URL-encoded path of the project
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project](golab_project.md)	 - Manage projects

