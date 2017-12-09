## golab user emails delete

Delete email for current / given user

### Synopsis


If no user_id is given: Deletes email owned by currently authenticated user.
If a user_id is given: Deletes email owned by a specified user. Available only for admin.

```
golab user emails delete [flags]
```

### Options

```
  -i, --email_id int   (required) id of email to be deleted
  -h, --help           help for delete
  -u, --user_id int    (optional) id of user to delete email from
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user emails](golab_user_emails.md)	 - Manage emails for users

