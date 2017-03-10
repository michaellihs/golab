Gitlab CLI written in Go [![Build Status](https://travis-ci.org/michaellihs/golab.svg?branch=master "Travis CI status")](https://travis-ci.org/michaellihs/golab)
======================== 

This project provides a Gitlab Command Line Interface (CLI) written in Go.


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


Further Resources
=================

* [Gitlab API Documentation](https://docs.gitlab.com/ee/api/README.html)
* [Cobra Library (Go CLI Library)](https://github.com/spf13/cobra)
* [Viper Library (Go Flags Library)](https://github.com/spf13/viper)
* [Sling (Go HTTP Library)](https://github.com/dghubble/sling)
* [Ginkgo (Go Testing Library)](https://onsi.github.io/ginkgo/)
* [GoMock (Go Mocking Library)](https://github.com/golang/mock)
* [go-gitlab (Go Gitlab Library)](https://github.com/xanzy/go-gitlab)
