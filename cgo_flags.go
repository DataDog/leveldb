package levigo

// TODO Elijah: SRCDIR?

// #cgo CFLAGS: -Ivendor/leveldb/include -fno-builtin-memcmp -DOS_LINUX -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT
// #cgo CXXFLAGS: -Ivendor/leveldb/include -Ivendor/leveldb -std=c++11 -fno-builtin-memcmp -lpthread -DOS_LINUX -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT
// #cgo LDFLAGS: -lpthread
// #include "leveldb/c.h"
import "C"
