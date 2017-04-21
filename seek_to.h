
#include "leveldb/c.h"
#include <stdbool.h>

typedef struct {
    unsigned char valid;
    unsigned char equal;
} SeekResult;

SeekResult leveldb_iter_seek_to(leveldb_iterator_t* iter, const char* k, size_t klen) {
    // do the seek
    leveldb_iter_seek(iter, k, klen);

    SeekResult sr;
    sr.valid = leveldb_iter_valid(iter);

    if (sr.valid != 0) {
        size_t olen;
        const char* ok = leveldb_iter_key(iter, &olen);

        sr.equal = ((klen == olen) && (memcmp(k, ok, olen) == 0));
    }

    return sr;
}
