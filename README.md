Gitlab CLI written in Go
========================

This project is intended to provide a Gitlab Command Line Interface (CLI) written in Go.


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
