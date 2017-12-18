Gitlab CLI written in Go [![Build Status](https://travis-ci.org/michaellihs/golab.svg?branch=master "Travis CI status")](https://travis-ci.org/michaellihs/golab)
======================== 

This project provides a Command Line Interface (CLI) for Gitlab written in Go. The project uses the [go-gitlab client](https://github.com/xanzy/go-gitlab) for the communication with Gitlab.

It allows you to run Gitlab administration tasks from your command line.

[TOC levels=1-3]: # "## Table of Contents"

## Table of Contents
- [Usage](#usage)
    - [Examples](#examples)
    - [Installation](#installation)
    - [Configuration](#configuration)
        - [Login with Username and Password](#login-with-username-and-password)
        - [Login with Access Token](#login-with-access-token)
    - [ZSH auto-completion](#zsh-auto-completion)
- [Development](#development)
    - [API Debugging](#api-debugging)
    - [Build and test the application](#build-and-test-the-application)
    - [Ginkgo Tests](#ginkgo-tests)
    - [Update vendored dependencies](#update-vendored-dependencies)
    - [Translate API Doc into Flag Structs](#translate-api-doc-into-flag-structs)
    - [Gitlab Docker Image](#gitlab-docker-image)
    - [Troubleshooting](#troubleshooting)
        - [`panic: trying to get string value of flag of type int`](#panic-trying-to-get-string-value-of-flag-of-type-int)
- [TODOs](#todos)
    - [Support multiple Targets](#support-multiple-targets)
    - [Support GPG keys in user command](#support-gpg-keys-in-user-command)
    - [Support for nested groups](#support-for-nested-groups)
    - [Fix password issue on Windows](#fix-password-issue-on-windows)
- [Further Resources](#further-resources)


Usage
=====

Examples
--------

* create a user

   ``` bash
   golab user create --email username@company.com --username username --password 12341234 --name "User McArthur" --skipConfirmation
   ```

* modify a user

   ``` bash
   golab user modify -i 41 --admin true
   ```

* create a new project / repository

   ``` bash
   golab project create -g my-group -n my-project
   ```

* add an ssh key for a user

   ``` bash
   golab user ssh-keys add --key "`cat ~/.ssh/id_rsa.pub`" --title "my dsa key"
   ```

* query your json output with [jq](https://stedolan.github.io/jq/)

   ``` bash
   golab project list | jq ".[] | {name: .name, id: .id}
   ```

For a complete documentation of features, check the [generated documentation](doc/golab.md)


Installation
------------

Install the CLI tool with

    go get -u github.com/michaellihs/golab
    go install github.com/michaellihs/golab

or download a [binary release](https://github.com/michaellihs/golab/releases).


Configuration
-------------

### Login with Username and Password

Run the following command to login with your username and password

    golab login --host <hostname> --user <username> [--password <password>]

If `--password` is omitted, you'll be prompted to enter your password interactively.

According to [this discussion](https://github.com/xanzy/go-gitlab/issues/267) the login with username and password might not work with newer Gitlab versions.


### Login with Access Token

First create a Gitlab [access token for your user](https://docs.gitlab.com/ce/user/profile/personal_access_tokens.html) in Gitlab (most likely an admin user).

Create a file `.golab.yml` in either `~/` or the directory you want to use golab with the following content:

    ---
    url: "http(s)://<gitlab url>"
    token: "<access token>"

Test your configuration - e.g. by running `golab project` to get a list of projects from your Gitlab server.


ZSH auto-completion
-------------------

The auto-completion file for ZSH can be generated with

     golab zsh-completion --path zsh/_golab

TODO: After the `#compdef` header, add a `#autoload` - see http://zsh.sourceforge.net/Doc/Release/Completion-System.html

Check where to add your auto-complete files with `echo $FPATH` and copy the generated file there with

    cp zsh/_golab /usr/local/share/zsh/site-functions/_golab

Don't forget to reload / restart your ZSH shell after changing the auto-complete file (e.g. `source ~/.zshrc`).


Development
===========

API Debugging
-------------

Run `curl` requests against the API:

    curl --header "PRIVATE-TOKEN: FqBiTTJ4oRPdskWDTktr" -H "Content-Type: application/json" -X PUT -d '{"admin": true}' http://localhost:8080/api/v4/users/41


Build and test the application
------------------------------

There is a `makefile` included that can build and test the application and render the automatically generated documentation:

*  `make` - build the application

*  `make test` - run the tests. Before you run the tests, you have to set the environment variables

  * `GITLAB_HOST`
  * `GITLAB_ROOT_USER`
  * `GITLAB_ROOT_PASSWORD`

* `make gendoc` - render the documentation


Ginkgo Tests
------------

Run Ginkgo tests with

    cd cmd
    ginkgo -v


Update vendored dependencies
----------------------------

    govendor fetch github.com/spf13/cobra
    

Translate API Doc into Flag Structs
-----------------------------------

Regular expression for replace in IntelliJ

    (\s+)([^\s]+?)\s+([^\s]+?)\s+([^\s]+?)\s+(.+)
    $1$2 *$3 `flag_name:"$2" type:"$3" required:"$4" description:"$5"`


Gitlab Docker Image
-------------------

For local development, you can use a [Gitlab Docker image](https://docs.gitlab.com/omnibus/docker/README.html). There are some pitfalls, when using Gitlab Docker image on a Mac:

* You cannot properly mount the `/var/opt/gitlab` directory due to issues with NFS mounts on Mac
* The ssh port `22` is already in use on the Mac, if a ssh server is running

Therefore adapt the provided run command to the following:

    sudo docker run --detach \
        --hostname gitlab.example.com \
        --publish 443:443 --publish 80:80 --publish 8022:22 \
        --name gitlab \
        --volume /tmp/gitlab/config:/etc/gitlab \
        --volume /tmp/gitlab/logs:/var/log/gitlab \
        gitlab/gitlab-ce:9.5.10-ce.0

Afterwards you can start the (existing) container with:

    sudo docker start gitlab

**Attention** we are currently testing against Gitlab version `9.5.10-ce.0`. Make sure to pin the version of the Docker image instead of using `latest`.


Troubleshooting
---------------

### `panic: trying to get string value of flag of type int`

If you see `panic: trying to get string value of flag of type int`, most likely you used a flag type other than `*string` for a flag that needs transformation, e.g.

````
GroupAccess *string  `flag_name:"group_access" short:"a" type:"integer" transform:"str2AccessLevel" required:"yes" description:"..."`
````

Remember: you only can use `*string` as field type, when you need a transformation of the value!


TODOs
=====

Support multiple Targets
------------------------

`golab login` should take a parameter `-e` that sets an environment which we can later on select with `-e` in each command or read from `$golab_env`.

Therefore we also need to change the structure of the `.golab.yml` like this:

    ---
    gitlab.com:
      url: "https://gitlab.com"
      token: "gitlab_com_token"

    localhost:
      url: "http://localhost:12345"
      token: "localhost_token"

This allows working with multiple Gitlab servers at the same time.


Support GPG keys in user command
--------------------------------

Currently the [go-gitlab library](https://github.com/xanzy/go-gitlab) provides no support for GPG keys, neither does this to.


Support for nested groups
-------------------------

Currently there is no support for [nested groups](https://docs.gitlab.com/ce/api/groups.html#list-a-groups-39-s-subgroups) since this feature is only available in Gitlab >= 10.3


Fix password issue on Windows
-----------------------------

Currently, when logging in with username and password (on Cygwin) on Windows, the password is not hidden (as it is on MacOS).


Further Resources
=================

* [Gitlab API Documentation](https://docs.gitlab.com/ee/api/README.html)
* [Cobra Library (Go CLI Library)](https://github.com/spf13/cobra)
* [Viper Library (Go Flags Library)](https://github.com/spf13/viper)
* [Sling (Go HTTP Library)](https://github.com/dghubble/sling)
* [Ginkgo (Go Testing Library)](https://onsi.github.io/ginkgo/)
* [GoMock (Go Mocking Library)](https://github.com/golang/mock)
* [go-gitlab (Go Gitlab Library)](https://github.com/xanzy/go-gitlab)
* [govendor on Heroku](https://devcenter.heroku.com/articles/go-dependencies-via-govendor)
