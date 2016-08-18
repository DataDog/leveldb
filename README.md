# levigo

levigo is a Go wrapper for LevelDB. This fork makes levigo go-gettable by
including the C++ source for LevelDB, and building it with the Go code.

This package currently includes LevelDB 1.19, and only supports Linux.

The API has been godoc'ed and [is available on the
web](http://godoc.org/github.com/jmhodges/levigo).

Questions answered at `golang-nuts@googlegroups.com`.

## Building

This package (likely) only works on Linux.

You'll need gcc and g++. Then:

`go get github.com/elijahandrews/levigo`

This will take a long time as it compiles the LevelDB source code along with the
Go library.

To avoid waiting for compilation every time you want to build your project, you can run:

`go install github.com/elijahandrews/levigo`

## Caveats

Comparators and WriteBatch iterators must be written in C in your own
library. This seems like a pain in the ass, but remember that you'll have the
LevelDB C API available to your in your client package when you import levigo.

An example of writing your own Comparator can be found in
<https://github.com/jmhodges/levigo/blob/master/examples>.
