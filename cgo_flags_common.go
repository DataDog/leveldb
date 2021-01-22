package levigo

// #cgo CFLAGS: -I${SRCDIR}/deps/lz4 -I${SRCDIR}/deps/leveldb/include -fno-builtin-memcmp -g -O2 -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT
// #cgo CXXFLAGS: -I${SRCDIR}/deps/lz4 -I${SRCDIR}/deps/leveldb/include -I${SRCDIR}/deps/leveldb -std=c++11 -fno-builtin-memcmp -g -O2 -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT
// #cgo LDFLAGS: -lpthread
// #include "leveldb/c.h"
import "C"
