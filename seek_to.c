
#include "seek_to.h"

// Seek to the given key and return results about the seek.
SeekResult leveldb_iter_seek_to(leveldb_iterator_t* iter, const char* k, size_t klen) {
    leveldb_iter_seek(iter, k, klen);

    SeekResult sr;
    sr.valid = leveldb_iter_valid(iter);

    if (sr.valid != 0) {
        size_t out_key_len;
        const char* out_key = leveldb_iter_key(iter, &out_key_len);

        sr.equal = ((klen == out_key_len) && (memcmp(k, out_key, klen) == 0));
    }

    return sr;
}
