## golab user emails add

Create new email (for user)

### Synopsis


If no user_id is given: Creates a new email owned by the currently authenticated user.
If a user_id is given: Create new email owned by specified user. Available only for admin

Will return created email on success.

```
golab user emails add [flags]
```

### Options

```
  -e, --email string     (required) email address
  -h, --help             help for add
  -u, --user_id string   (optional) id or username of user to add email to
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user emails](golab_user_emails.md)	 - User emails

