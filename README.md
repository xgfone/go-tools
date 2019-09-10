# go-tools [![Build Status](https://travis-ci.org/xgfone/go-tools.svg?branch=master)](https://travis-ci.org/xgfone/go-tools) [![GoDoc](https://godoc.org/github.com/xgfone/go-tools/v6?status.svg)](http://godoc.org/github.com/xgfone/go-tools) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-tools/master/LICENSE)

A utility tool library of Golang.

The current version is **`v6`** to support Semantic Import Versioning. See [Doc](https://godoc.org/github.com/xgfone/go-tools).

## Announcement
These packages only depend on the standard libraries, not any third-part packages. `Go1.11+` will be supported.

## Installation
```shell
$ go get -u github.com/xgfone/go-tools/v6
```

### Test
```shell
go test github.com/xgfone/go-tools/v6/...
```

## Subpackages

**Notice:**

subpackage   |   notice
-------------|-----------
cache        | Supply some caches, such as `LRUCache`. Notice: LRUCache is copied from `github.com/youtube/vitess/go/cache`.
execution    | execution executes a command line program in a new process and returns an output.
file         | Some convenient functions about the file operation.
function     | Collect some convenient funtions, for example, calling a function or method dynamically, comparing two values, getting a integer range, determining whether a value is in a map or slice, etc.
io2          | The supplement of the standard library of `io`.
json2        | The supplement of the standard library of `json`.
lifecycle    | The manager of the lifecycle of some apps in a program.
net2         | The supplement of the standard library `net`, such as some helpers about net.
option       | Supply a type to represent the optional value referring to Option in Rust.
pools        | Some simple convenient pools, such as `BytesPool`, `BufferPool`, `ResourcePool`, etc.
signal2      | The supplement of the standard library of `signal`, such as `HandleSignal`.
sort2        | The supplement of the standard library of `sort`.
strings2     | The supplement of the standard library of `strings`.
sync2        | The supplement of the standard library `sync`, such as some atomic types.
tag          | Find and get the tags in a struct.
types        | Some assistant functions about type, such as the type validation and conversion, etc.
wait         | Poll or listen for changes to a condition. It's copied from `k8s.io/apimachinery/pkg/util/wait`.

## Example
See the `test` file of each subpackage.
