# GoGPAT (Go GitLab Personal Access Token)
[![Build Status](https://travis-ci.org/solidnerd/gogpat.svg?branch=master)](https://travis-ci.org/solidnerd/gogpat)[![Coverage Status](https://coveralls.io/repos/github/solidnerd/gogpat/badge.svg?branch=master)](https://coveralls.io/github/solidnerd/gogpat?branch=master)

GoGPAT helps you to create personal access token via cli.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.


### Installing

```console
 go get -u github.com/solidnerd/gogpat
```

### How to use it

#### Create a token for gitlab.com

```console
$ gogpat create -u example -p example.pw
```

#### Create a token for your own instance

```console
$ gogpat create -u example -p example.pw https://gitlab.example.com
```

#### Options

```console
NAME:
   gogpat create - creates a gitlab api token for the specified gitlab

USAGE:
   gogpat create [command options] [arguments...]

OPTIONS:
   --api, -a                   Access the authenticated user's API: Full access to GitLab as the user, including read/write on all their groups and projects
   --read_user, --ru           Read the authenticated user's personal information: Read-only access to the user's profile information, like username, public email and full name
   --read_registry, --rr       Read the authenticated user's personal information: Read-only access to the user's profile information, like username, public email and full name
   --sudo, -s                  Perform API actions as any user in the system (if the authenticated user is an admin: Access to the Sudo feature, to perform API actions as any user in the system (only available for admins)
   --user value, -u value      Sets the user for the login
   --password value, -p value  Sets the user for the login
   --name value, -n value      Sets the name of the personal token by default it's gogpat
   --expiry value, --ex value  Sets the expiry date of the personal token it's should be in format like this 2017-12-22
```

## Authors

* **Niclas Mietz** - *Initial work* - [solidnerd](https://github.com/solidnerd)

See also the list of [contributors](https://github.com/solidnerd/gogpat/contributors) who participated in this project.

## License

This project is licensed under the Apache License 2.0  - see the [LICENSE](LICENSE) file for details
