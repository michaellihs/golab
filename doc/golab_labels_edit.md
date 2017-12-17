## golab labels edit

Edit an existing label

### Synopsis


Updates an existing label with new name or new color. At least one parameter is required, to update the label.

```
golab labels edit [flags]
```

### Options

```
  -c, --color string         (optional) (required, if new_name is not provided) The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the CSS color names
  -d, --description string   (optional) The new description of the label
  -h, --help                 help for edit
  -i, --id string            (required) The ID or URL-encoded path of the project owned by the authenticated user
  -n, --name string          (required) The name of the existing label
  -u, --new_name string      (optional) (required, if color is not provided) The new name of the label
  -p, --priority int         (optional) The new priority of the label. Must be greater or equal than zero or null to remove the priority.
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab labels](golab_labels.md)	 - Manage labels

