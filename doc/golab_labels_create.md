## golab labels create

Create a new label

### Synopsis


Creates a new label for the given repository with the given name and color.

```
golab labels create [flags]
```

### Options

```
  -c, --color string         (required) The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the CSS color names
  -d, --description string   (optional) The description of the label
  -h, --help                 help for create
  -i, --id string            (required) The ID or URL-encoded path of the project owned by the authenticated user
  -n, --name string          (required) The name of the label
  -p, --priority int         (optional) The priority of the label. Must be greater or equal than zero or null to remove the priority.
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab labels](golab_labels.md)	 - Manage labels

