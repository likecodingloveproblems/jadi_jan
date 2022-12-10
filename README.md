## jadi_jan

[![Build Status](https://travis-ci.com/likecodingloveproblems/jadi_jan.svg?branch=master)](https://travis-ci.com/likecodingloveproblems/jadi_jan)
[![GitHub](https://img.shields.io/github/license/likecodingloveproblems/jadi_jan)](https://github.com/likecodingloveproblems/jadi_jan)
[![GoDoc](https://godoc.org/github.com/likecodingloveproblems/jadi_jan?status.svg)](https://godoc.org/github.com/likecodingloveproblems/jadi_jan)  [![Test Coverage](https://codecov.io/gh/likecodingloveproblems/jadi_jan/branch/master/graph/badge.svg)](https://codecov.io/gh/likecodingloveproblems/jadi_jan)
[![Release](https://img.shields.io/github/release/likecodingloveproblems/jadi_jan)](https://github.com/likecodingloveproblems/jadi_jan/releases/latest)

This is a very simple project to be a backend for [JadiJan](https://github.com/milad7212/jadijan).
I try to implement [active record pattern](https://www.martinfowler.com/eaaCatalog/activeRecord.html) in this project.

### Usage
It has four endpioints:
```
GET /counter/  # returns how many persons remind jadi ❤️
POST /counter/  # Anther person remind jadi ❤️

GET /comments/ # Returns all comments on lessons they have learned from jadi ❤️
POST /comments/ # Add a comment on lessons they have learned from jadi ❤️
```

######  Prerequisites

This project is written in Go and uses [echo](https://github.com/labstack/echo) as http framework. default port is :1323, you can customize it in `main.go` file.
For the sake of performance I used [redis](https://github.com/redis/redis) as data store. So you need a redis server on default port (:6379), you can customize redis connection in `models/store.go` file.

Before you get started, make sure you have installed the following tools::

    $ python3 -m pip install -U cookiecutter>=1.4.0
    $ python3 -m pip install pre-commit bump2version invoke ruamel.yaml halo
    $ go get -u golang.org/x/tools/cmd/goimports
    $ go get -u golang.org/x/lint/golint
    $ go get -u github.com/fzipp/gocyclo/cmd/gocyclo
    $ go get -u github.com/mitchellh/gox  # if you want to test building on different architectures

**Remember**: To be able to excecute the tools downloaded with `go get`, 
make sure to include `$GOPATH/bin` in your `$PATH`.
If `echo $GOPATH` does not give you a path make sure to run
(`export GOPATH="$HOME/go"` to set it). In order for your changes to persist, 
do not forget to add these to your shells `.bashrc`.

With the tools in place, it is strongly advised to install the git commit hooks to make sure checks are passing in CI:
```bash
invoke install-hooks
```

You can check if all checks pass at any time:
```bash
invoke pre-commit
```

Note for Maintainers: After merging changes, tag your commits with a new version and push to GitHub to create a release:
```bash
bump2version (major | minor | patch)
git push --follow-tags
```

#### Installation
First install a redis server. Then run the following commands:

```bash
git clone https://github.com/likecodingloveproblems/jadi_jan.git
cd jadi_jan
go build .
./jadi_jan
```

### Note
Redis data must be CONSISTENT.