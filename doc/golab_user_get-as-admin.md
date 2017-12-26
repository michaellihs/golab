## golab user get-as-admin

Lookup users by username

### Synopsis


Lookup users by username

```
golab user get-as-admin [flags]
```

### Options

```
      --created_after string    (optional) Search users created after, e.g. 2001-01-02
      --created_before string   (optional) Search users created before, e.g. 2001-01-02
      --external                (optional) If set to true only external users will be returned
      --external_uid string     (optional) External UID of the user to look up (only together with provider)
  -h, --help                    help for get-as-admin
      --provider string         (optional) External provider of user to look up
  -u, --username string         (optional) Username of the user to look up
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user](golab_user.md)	 - Manage Gitlab users

