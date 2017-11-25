# go-tools
A utility tool library of Golang.

The current version is **`v2`**, which is not compatible with `v1`. `v2` rearranges the sub-packages. Relative to `v1`, `v2` has some changes as follows:

- Add some new functions.
- Merge some sub-packages into one sub-package, such as merging `extremum`, `compare`, `method`, `values` into `function`, etc.
- Rename the name of some sub-packages, such as `log2`, `net2`, `strings2`, etc.
- Remove some sub-packages, such as `datetime`, etc.
- Remove some types, such as `SMap` and `Slice` in the sub-package `values` in `v1`, etc.
- Remove some deprecated types or functions, such as `Dispatcher` in the sub-package `worker`, etc.
- ......

## Announcement
These packages only depend on the standard libraries, not any third-part packages.

## Installation
```shell
$ go get -u github.com/xgfone/go-tools
```

### Test
```shell
go test github.com/xgfone/go-tools/...
```

## Subpackages

**Notice:**

Some sub-packages need the logging output. Now the sub-package `log2` supplies the unified logging output function, `ErrorF` and `DebugF`, which will output the logging to `os.Stderr`. You can set them to set the logging output of all the sub-packages, such as `net2/http2`, `net2/server`, and `worker`, etc.

subpackage   |   notice
-------------|-----------
cache        | Supply some caches, such as `LRUCache`. Notice: LRUCache is copied from `github.com/youtube/vitess/go/cache`.
execution    | execution executes a command line program in a new process and returns an output.
file         | Some convenient functions about the file operation.
function     | Collect some convenient funtions, for example, calling a function or method dynamically, comparing two values, deciding the maximum or minimum, getting a integer range, determining whether a value is in a map or slice, etc.
io2          | The supplement of the standard library of `io`, such as `Close`.
lifecycle    | The manager of the lifecycle of some apps in a program.
log2         | Supply a global debug and error log function, and a log function adaptor.
log2/handler | The logger handler, such as `TimedRotatingFile` and `SizedRotatingFile`.
net2         | The supplement of the standard library `net`, such as some helpers about net.
net2/http2   | The supplement of the standard library `http`, not the protocal `http2`.
net2/server  | The simple `TCP` / `UDP` server.
os2          | The supplement of the standard library of `os`, such as `Exit`.
pagination   | It is usually used to compute the web pagination.
pools        | Some simple convenient pools, such as `BufPool`, `ResourcePool`, `AddrTCPConnPool`, etc.
queue        | Supply the `Queue` interface, and some implementations such as `NewMemoryQueue` based on channel and `NewListQueue` based on list.
signal2      | The supplement of the standard library of `signal`, such as `HandleSignal`.
sort2        | The supplement of the standard library of `sort`, such as the key-value sort.
strings2     | The supplement of the standard library of `strings`.
sync2        | The supplement of the standard library `sync`, such as `ResourceLock` for locking a certain resource by its id and some atomic types.
tags         | Manage the tags in a struct.
types        | Some assistant functions about type, such as the type validation and conversion, etc.
wait         | Poll or listen for changes to a condition. It's copied from `k8s.io/apimachinery/pkg/util/wait`.
worker       | A worker pool with the dispatcher based on channel.

## Example
See the `test` file of each subpackage.
