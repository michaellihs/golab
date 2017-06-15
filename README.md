Gitlab CLI written in Go [![Build Status](https://travis-ci.org/michaellihs/golab.svg?branch=master "Travis CI status")](https://travis-ci.org/michaellihs/golab)
======================== 

This project provides a Command Line Interface (CLI) for Gitlab written in Go. The project uses the [go-gitlab client](https://github.com/xanzy/go-gitlab) for the communication with Gitlab.


Usage
-----

Create a file `.golab.yml` in either `~/` or the directory you want to use golab with the following content:

    ---
    url: "https://gitlab.com"
    token: "YOUR PRIVATE TOKEN"

Replace `gitlab.com` with the URL of your Gitlab server.

Test your configuration - e.g. by running `golab project` to get a list of projects from your Gitlab server.


Build and run the application
-----------------------------

    go install github.com/michaellihs/golab
    golab


Run the Tests
-------------

Install Ginkgo

    go get github.com/onsi/ginkgo/ginkgo
    go get github.com/onsi/gomega

Run the tests (e.g. in `model`) with

    cd model
    ginkgo


Install the application
-----------------------

First run

    go install github.com/michaellihs/golab

You can then use the application by simply typing

    golab
    

Generate Markdown Documentation
-------------------------------

Using the command

    golab gendoc -p PATH

let's you render a set of markdown documentation for the application.
    

Gitlab Docker Image
===================

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
        gitlab/gitlab-ce:latest

Afterwards you can start the (existing) container with:

    sudo docker start gitlab


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


Further Resources
=================

* [Gitlab API Documentation](https://docs.gitlab.com/ee/api/README.html)
* [Cobra Library (Go CLI Library)](https://github.com/spf13/cobra)
* [Viper Library (Go Flags Library)](https://github.com/spf13/viper)
* [Sling (Go HTTP Library)](https://github.com/dghubble/sling)
* [Ginkgo (Go Testing Library)](https://onsi.github.io/ginkgo/)
* [GoMock (Go Mocking Library)](https://github.com/golang/mock)
* [go-gitlab (Go Gitlab Library)](https://github.com/xanzy/go-gitlab)
