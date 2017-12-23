## golab project hooks add

Add project hook

### Synopsis


Adds a hook to a specified project.

```
golab project hooks add [flags]
```

### Options

```
      --enable_ssl_verification   (optional) Do SSL verification when triggering the hook
  -h, --help                      help for add
  -i, --id string                 (required) The ID or URL-encoded path of the project
      --issues_events             (optional) Trigger hook on issues events
      --job_events                (optional) Trigger hook on job events
      --merge_requests_events     (optional) Trigger hook on merge requests events
      --note_events               (optional) Trigger hook on note events
      --pipeline_events           (optional) Trigger hook on pipeline events
      --push_events               (optional) Trigger hook on push events
      --tag_push_events           (optional) Trigger hook on tag push events
      --token string              (optional) Secret token to validate received payloads; this will not be returned in the response
  -u, --url string                (required) The hook URL
      --wiki_events               (optional) Trigger hook on wiki events
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project hooks](golab_project_hooks.md)	 - Manage project hooks.

