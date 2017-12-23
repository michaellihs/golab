## golab user emails ls

List emails

### Synopsis


If no user_id is given: get a list of currently authenticated user's emails.
If a user_id is given: Get a list of a specified user's emails. Available only for admin

```
golab user emails ls [flags]
```

### Options

```
  -h, --help             help for ls
  -u, --user_id string   (optional) The ID of the user or username of user to list emails for. If none is given, emails of currently logged in user are shown
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user emails](golab_user_emails.md)	 - User emails

