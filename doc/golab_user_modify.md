## golab user modify

User modification

### Synopsis


Modifies an existing user. Only administrators can change attributes of a user.

```
golab user modify [flags]
```

### Options

```
      --admin                 (optional) User is admin - true or false (default)
      --bio string            (optional) User's biography
      --can_create_group      (optional) User can create groups - true or false
  -e, --email string          (optional) Email
      --extern_uid string     (optional) External UID
      --external              (optional) Flags the user as external - true or false(default)
  -h, --help                  help for modify
  -i, --id string             (required) User ID or user name of user to be deleted
      --linkedin string       (optional) LinkedIn
      --location string       (optional) User's location
  -n, --name string           (optional) Name
      --organization string   (optional) Organization name
  -p, --password string       (optional) Password
      --projects_limit int    (optional) Number of projects user can create
      --provider string       (optional) External provider name
      --skip_confirmation     (optional) Skip confirmation - true or false (default)
      --skype string          (optional) Skype ID
      --twitter string        (optional) Twitter account
  -u, --username string       (optional) Username
      --website_url string    (optional) Website URL
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user](golab_user.md)	 - Manage Gitlab users

