# go-tools
A utility tool library of Golang.

## Announcement
These packages only depend on the standard libraries, not any third-part packages.

## Installation
```shell
$ go get -u github.com/xgfone/go-tools
```

## Subpackages

subpackage   |   notice
-------------|-----------
atomics      | Some atomic types, such as `AtomicInt32`, `AtomicInt64`, `AtomicDuration`, `AtomicBool`, `AtomicString`, `Semaphore`, `Count`, etc. Notice: Some types are copied from `github.com/youtube/vitess/go/sync2`.
cache        | Supply some caches, such as `LRUCache`. Notice: LRUCache is copied from `github.com/youtube/vitess/go/cache`.
compare      | Compare whether the first is greater than, less than, or equal to the second.
datetime     | Some convenient functions about datetime.
execution    | execution executes a command line program in a new process and returns an output.
extremum     | Get the maximal or the minimal of both the values.
file         | Some convenient functions about the file operation.
function     | Call a function dynamically.
lifecycle    | The manager of the lifecycle of some apps in a program.
log/handler  | The logger handler, such as `TimedRotatingFile` like `logging.handlers.TimedRotatingFileHandler` in Python.
method       | Call the method of a type dynamically.
nets         | The supplement of the standard library `net`, such as some helpers about net.
nets/https   | The supplement of the standard library `http`, not the protocal `https`.
nets/mac     | Standardize the mac address.
nets/server  | The simple `TCP` / `UDP` server.
os2          | The supplement of the standard library of `os`, such as `Exit`.
pagination   | It is usually used to compute the web pagination.
pools        | Some simple convenient pools, such as `BufPool`, `GoPool`, `ResourcePool`, `AddrTCPConnPool`, etc. Notice: **Rename to `pools` from `v0.40`**.
queue        | Supply the `Queue` interface, and some implementations such as `NewMemoryQueue` based on memory.
signal2      | The supplement of the standard library of `signal`, such as `HandleSignal`.
slice        | Get a value from a slice and check whether a value exists in a slice.
sort2        | The supplement of the standard library of `sort`, such as `SortMap`.
string2      | The supplement of the standard library of `strings`.
sync2        | The supplement of the standard library `sync`, such as `ResourceLock` for locking a certain resource by its id.
tags         | Manage the tags in a struct.
validation   | Some validations, such as email, url, or the type of a value, etc.
values       | Get a value from a `slice` or `map`.
wait         | Poll or listen for changes to a condition. It's copied from `k8s.io/apimachinery/pkg/util/wait`.
worker       | A worker pool with the dispatcher based on channel.

## Example
See the `test` file of each subpackage.
