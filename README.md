# levigo

levigo is a Go wrapper for LevelDB. This fork makes levigo go-gettable by
including the C++ source for LevelDB, and building it with the Go code.

This package currently includes LevelDB 1.19, and supports Linux and OS X.
It should work on Go 1.5 or greater.

The API has been godoc'ed and [is available on the
web](http://godoc.org/github.com/jmhodges/levigo).

Questions answered at `golang-nuts@googlegroups.com`.

## Building

This package supports Linux and OS X.

You'll need gcc and g++. Then:

`go get github.com/elijahandrews/levigo`

This will take a long time as it compiles the LevelDB source code along with the
Go library.

To avoid waiting for compilation every time you want to build your project, you can run:

`go install github.com/elijahandrews/levigo`

## Updating Vendored LevelDB

The embedded LevelDB source code is located in `vendor/leveldb`. In order to get
`go build` to compile LevelDB, we symlink the required LevelDB source files to
the root directory of the project. These symlinks are prefixed with `vendor_`.

To change the embedded version of LevelDB, do the following:

1. `rm vendor_*.cc`
1. Replace `vendor/leveldb` with the source code of the desired version of LevelDB
1. On Linux, run `./vendor/leveldb/build_detect_platform linux.mk build_flags`
1. On OS X, run `./vendor/leveldb/build_detect_platform darwin.mk build_flags`
1. `cat build_flags/*`, take a quick look at the new compiler and linker flags
   and see if there are any flags we're missing in `cgo_flags_*.go`. If so, add
   them.
1. Create symlinks to the new source files:
```
for sf in $(make source_files); do ln -s vendor/leveldb/$sf $(echo vendor_$sf | sed s,/,_,g); done
```
1. On OS X and Linux, run `go build` and verify that everything compiles.

## Caveats

Comparators and WriteBatch iterators must be written in C in your own
library. This seems like a pain in the ass, but remember that you'll have the
LevelDB C API available to your in your client package when you import levigo.

An example of writing your own Comparator can be found in
<https://github.com/jmhodges/levigo/blob/master/examples>.
