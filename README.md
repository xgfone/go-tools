# go-tools
A utility tool library of Golang.

## Announcement
These packages only depend on the standard libraries, not any third packages.

## Installation
```shell
$ go get -u github.com/xgfone/go-tools
```

## Subpackages

subpackage   |   notice
-------------|-----------
caller       | Get the filename and the line number where to call these functions.
checksum     | Calculate the checksum, such as icmp.
compare      | Compare whether the first is greater than, less than, or equal to the second.
count        | Atomic count.
daemon       | Make the current process to the daemon process.
datetime     | Some convenient functions about datetime.
exception    | Exception handler like "parent.child.sub-child...sub-child".
extremum     | Get the maximal or the minimal of both the values.
file         | Some convenient functions about the file operation.
function     | Call a function dynamically.
method       | Call the method of a type dynamically.
net/endian   | Convert between int/uint and the big/little endian.
net/mac      | Standardize the mac address.
net/server   | The simple TCP/UDP server.
parse        | Convert a string to bool, int, uint, float, etc.
pool         | Some simple convenient pools, such as BufPool, GoPool, etc.
slice        | Get a value from a slice.
tbucket      | The Simple Token Bucket like HTB in Linux TC.
values       | Get a value from the slice or the map.

## Example
See the `test` file of each subpackage.
