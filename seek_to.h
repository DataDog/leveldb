
#include "leveldb/c.h"

// SeekResult is a struct that contains information about the results of a
// leveldb seek.
typedef struct {
    unsigned char valid;
    unsigned char equal;
} SeekResult;

// Seek the given iterator to the key and return a result.
SeekResult leveldb_iter_seek_to(leveldb_iterator_t* iter, const char* k, size_t klen);
