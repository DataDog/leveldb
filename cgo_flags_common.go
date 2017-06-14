package levigo

// #cgo CFLAGS: -I${SRCDIR}/vendor/lz4 -I${SRCDIR}/vendor/leveldb/include -fno-builtin-memcmp -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT
// #cgo CXXFLAGS: -I${SRCDIR}/vendor/lz4 -I${SRCDIR}/vendor/leveldb/include -I${SRCDIR}/vendor/leveldb -std=c++11 -fno-builtin-memcmp -lpthread -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT
// #include "leveldb/c.h"
import "C"
