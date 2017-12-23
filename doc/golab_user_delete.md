## golab user delete

User deletion

### Synopsis


Deletes a user. Available only for administrators. This returns a 204 No Content status code if the operation was successfully or 404 if the resource was not found.

```
golab user delete [flags]
```

### Options

```
  -d, --hard_delete   (optional) If true, contributions that would usually be moved to the ghost user will be deleted instead, as well as groups owned solely by this user.
  -h, --help          help for delete
  -i, --id string     (required) User ID or user name of user to be deleted
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user](golab_user.md)	 - Manage Gitlab users

