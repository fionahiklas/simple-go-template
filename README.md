# simple-go-template

## Overview

Simple Go template application

![Main build](../../actions/workflows/build-and-test.yaml/badge.svg?branch=main)


## Quickstart

Install Go and run the following commands

``` 
make install_tools
make generate
make test_clean
```


## Development Notes

Details of how this project was created.  You don't need to follow these 
instructions to get things up and running

### Setup 

* Created initial module file

```
go mod init github.com/fionahiklas/simple-go-template
```

* Adding dependency 

``` 
go get github.com/sirupsen/logrus
```

* Created basic Makefile to be able to run tests and linting 


## References

### Go

* [Converting from arrays of interface](https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces)
* [Trim whitespace](https://yourbasic.org/golang/trim-whitespace-from-string/)
* [Listen on random/free port](https://stackoverflow.com/questions/43424787/how-to-use-next-available-port-in-http-listenandserve)
* [Read whole file](https://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable)
* [Using go ldflags to set version values](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)

### GitHub Actions

* [Setup Go action](https://github.com/actions/setup-go)
* [Go tutorial post](https://medium.com/swlh/setting-up-github-actions-for-go-project-ea84f4ed3a40)
* [Versioning for Go setup](https://github.com/npm/node-semver)
* [Actions build badge](https://docs.github.com/en/actions/monitoring-and-troubleshooting-workflows/adding-a-workflow-status-badge)

### UNIX

* [Getting fields from input](https://stackoverflow.com/questions/16136943/how-to-get-the-second-column-from-command-output)