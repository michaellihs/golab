## golab project fork

Fork project

### Synopsis


Forks a project into the user namespace of the authenticated user or the one provided.

The forking operation for a project is asynchronous and is completed in a background job. The request will return immediately. To determine whether the fork of the project has completed, query the import_status for the new project.

```
golab project fork [flags]
```

### Options

```
  -h, --help               help for fork
      --id string          (required) The ID or URL-encoded path of the project
      --namespace string   (required) The ID or path of the namespace that the project will be forked to
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project](golab_project.md)	 - Manage projects

