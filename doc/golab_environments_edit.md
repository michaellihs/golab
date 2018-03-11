## golab environments edit

Edit an existing environment

### Synopsis


Updates an existing environment's name and/or external_url.

It returns 200 if the environment was successfully updated. In case of an error, a status code 400 is returned.

```
golab environments edit [flags]
```

### Options

```
  -e, --environment_id int    (required) The ID of the environment
      --external_url string   (required) The new external_url
  -h, --help                  help for edit
  -i, --id string             (required) The ID or URL-encoded path of the project owned by the authenticated user
      --name string           (optional) The new name of the environment
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab environments](golab_environments.md)	 - Manage environments

