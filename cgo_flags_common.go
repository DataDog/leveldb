package levigo

// #cgo CFLAGS: -I${SRCDIR}/deps/lz4 -I${SRCDIR}/deps/leveldb/include -I${SRCDIR}/deps/snappy -fno-builtin-memcmp -O2 -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT
// #cgo CXXFLAGS: -I${SRCDIR}/deps/lz4 -I${SRCDIR}/deps/leveldb/include -I${SRCDIR}/deps/leveldb -I${SRCDIR}/deps/snappy -std=c++11 -fno-builtin-memcmp -O2 -lpthread -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT
// #include "leveldb/c.h"
import "C"
