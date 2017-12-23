## golab labels unsubscribe

Unsubscribe from a label

### Synopsis


Unsubscribes the authenticated user from a label to not receive notifications from it. If the user is not subscribed to the label, the status code 304 is returned.

```
golab labels unsubscribe [flags]
```

### Options

```
  -h, --help              help for unsubscribe
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

