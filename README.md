# go-tools [![Build Status](https://travis-ci.org/xgfone/go-tools.svg?branch=master)](https://travis-ci.org/xgfone/go-tools) [![GoDoc](https://godoc.org/github.com/xgfone/go-tools?status.svg)](http://godoc.org/github.com/xgfone/go-tools) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-tools/master/LICENSE)

A utility tool library of Golang.


## Installation
```shell
$ go get -u github.com/xgfone/go-tools/v7
```


## Subpackages

**Notice:**

subpackage   |   notice
-------------|-----------
cache        | Supply some caches, such as `LRUCache`. Notice: LRUCache is copied from `github.com/youtube/vitess/go/cache`.
execution    | execution executes a command line program in a new process and returns an output.
file         | Some convenient functions about the file operation.
function     | Collect some convenient funtions, for example, calling a function or method dynamically, comparing two values, etc.
io2          | The supplement of the standard library of `io`.
json2        | The supplement of the standard library of `json`.
lifecycle    | The manager of the lifecycle of some apps in a program.
net2         | The supplement of the standard library `net`, such as some helpers about net.
pools        | Some simple convenient pools, such as `BytesPool`, `BufferPool`, etc.
scanner      | the replacer of the stdlib `bufio.Scanner`, which adds the read offset.
signal2      | The supplement of the standard library of `signal`, such as `HandleSignal`.
slice        | Supply some assistant functions about slice.
sort2        | The supplement of the standard library of `sort`.
strings2     | The supplement of the standard library of `strings`.
tag          | Find and get the tags in a struct.
types        | Some assistant functions about type, such as the type validation.
wait         | Poll or listen for changes to a condition. It's copied from `k8s.io/apimachinery/pkg/util/wait`.

## Example
See the `test` file of each subpackage.
