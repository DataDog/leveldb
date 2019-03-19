
// for memcmp
#include <string.h>

#include "seek_to.h"

// Seek to the given key and return results about the seek.
SeekResult leveldb_iter_seek_to(leveldb_iterator_t* iter, const char* k, size_t klen) {
    leveldb_iter_seek(iter, k, klen);

    SeekResult sr = {0};
    sr.valid = leveldb_iter_valid(iter);

    if (sr.valid != 0) {
        // if we have a valid result, fetch the key
        leveldb_slice_t out_key = leveldb_iter_key(iter);

        sr.equal = ((klen == out_key.len) && (memcmp(k, out_key.data, klen) == 0));

        // if the key is equal also fetch the value.
        if (sr.equal != 0) {
            leveldb_slice_t val_out = leveldb_iter_value(iter);
            sr.val_data = val_out.data;
            sr.val_len = val_out.len;
        }
    }

    return sr;
}

unsigned char leveldb_iter_exists(leveldb_iterator_t* iter, const char* k, size_t klen) {
    leveldb_iter_seek(iter, k, klen);

    unsigned char exists = leveldb_iter_valid(iter);
    if (exists == 0) {
        return 0;
    }

    leveldb_slice_t out_key = leveldb_iter_key(iter);

    return ((klen == out_key.len) && (memcmp(k, out_key.data, klen) == 0));
}
