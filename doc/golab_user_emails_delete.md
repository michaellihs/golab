## golab user emails delete

Delete email for (current) user

### Synopsis


If no user_id is given: Deletes email owned by currently authenticated user.
If a user_id is given: Deletes email owned by a specified user. Available only for admin.

```
golab user emails delete [flags]
```

### Options

```
  -e, --email_id int     (required) id of email to be deleted
  -h, --help             help for delete
  -u, --user_id string   (optional) id or username of user to delete email from
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user emails](golab_user_emails.md)	 - User emails

