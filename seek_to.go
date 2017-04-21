package levigo

// #include "seek_to.h"
import "C"
import "unsafe"

type SeekResult struct {
	Valid bool
	Equal bool
}

func SeekTo(it *Iterator, key []byte) SeekResult {
	out := C.leveldb_iter_seek_to(it.Iter, (*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)))
	return SeekResult{
		Valid: ucharToBool(out.valid),
		Equal: ucharToBool(out.equal),
	}
}
