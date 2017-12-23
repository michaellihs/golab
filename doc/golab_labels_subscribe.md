## golab labels subscribe

Subscribe to a label

### Synopsis


Subscribes the authenticated user to a label to receive notifications. If the user is already subscribed to the label, the status code 304 is returned.

```
golab labels subscribe [flags]
```

### Options

```
  -h, --help              help for subscribe
  -i, --id string         (required) The ID or URL-encoded path of the project owned by the authenticated user
  -l, --label_id string   (required) The ID or title of a project's label
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab labels](golab_labels.md)	 - Manage labels

