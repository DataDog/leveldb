SOURCES=util/status.cc util/arena.cc util/bloom.cc util/hash.cc util/filter_policy.cc util/histogram.cc util/crc32c.cc util/cache.cc util/logging.cc util/options.cc util/comparator.cc util/coding.cc util/env_posix.cc util/env.cc table/table.cc table/block.cc table/merger.cc table/iterator.cc table/block_builder.cc table/table_builder.cc table/filter_block.cc table/format.cc table/two_level_iterator.cc db/log_writer.cc db/db_impl.cc db/filename.cc db/table_cache.cc db/log_reader.cc db/write_batch.cc db/c.cc db/builder.cc db/version_edit.cc db/dbformat.cc db/version_set.cc db/dumpfile.cc db/db_iter.cc db/repair.cc db/memtable.cc helpers/memenv/memenv.cc
MEMENV_SOURCES=helpers/memenv/memenv.cc
CC=cc
CXX=g++
PLATFORM=OS_LINUX
PLATFORM_LDFLAGS=-pthread
PLATFORM_LIBS= -lsnappy -ltcmalloc
PLATFORM_CCFLAGS= -fno-builtin-memcmp -pthread -DOS_LINUX -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT -DSNAPPY
PLATFORM_CXXFLAGS=-std=c++0x -fno-builtin-memcmp -pthread -DOS_LINUX -DLEVELDB_PLATFORM_POSIX -DLEVELDB_ATOMIC_PRESENT -DSNAPPY
PLATFORM_SHARED_CFLAGS=-fPIC
PLATFORM_SHARED_EXT=so
PLATFORM_SHARED_LDFLAGS=-shared -Wl,-soname -Wl,
PLATFORM_SHARED_VERSIONED=true
