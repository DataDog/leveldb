package levigo

import (
	"fmt"
)

// MultiKeyError encapsulates multiple errors encountered in a Get/Put-Many() call,
// behind the errors.error interface. Caller may cast the generic error object to
// a MultiKeyError and call Errors() and/or FailedKeyIndexes() to get additional details.
type MultiKeyError struct {
	// invariant: errsByIdx is either nil or has at least ONE entry, it's never an empty map
	errsByIdx map[int]error
}

func (mke *MultiKeyError) Error() string {
	if mke.errsByIdx == nil {
		return ""
	}

	// Show a few errors to informed the user
	var topErrsMsg string
	const maxNumErrs = 3
	cnt := 0
	for i := range mke.errsByIdx {
		topErrsMsg += fmt.Sprintf(" keys[%d]:%s;", i, mke.errsByIdx[i].Error())
		cnt++
		if cnt >= maxNumErrs {
			break
		}
	}
	return fmt.Sprintf("%d keys encountered error (here's a few:%s), cast via err.(*levigo.MultiKeyErrors).Errors() to get all the errors", len(mke.errsByIdx), topErrsMsg)
}

func (mke *MultiKeyError) GoString() string {
	return fmt.Sprintf("*%#v", *mke)
}

// FailedKeyIndexes() gives a list of indexes for which Get/PutMany() failed
// for key at keys[indexes[i]]; e.g. if this returns[2, 3, 7], it means
// Get/PutMany({keys[2], keys[3], keys[7]}) failed.
func (mke *MultiKeyError) FailedKeyIndexes() []int {
	if mke.errsByIdx == nil {
		return nil
	}

	indexes := make([]int, 0, len(mke.errsByIdx))
	for i := range mke.errsByIdx {
		indexes = append(indexes, i)
	}
	return indexes
}

// Errors() gives specific errors failed for each key from FailedKeyIndexes()
func (mke *MultiKeyError) Errors() []error {
	if mke.errsByIdx == nil {
		return nil
	}

	errs := make([]error, 0, len(mke.errsByIdx))
	for i := range mke.errsByIdx {
		errs = append(errs, mke.errsByIdx[i])
	}
	return errs
}

func (mke *MultiKeyError) addKeyErr(i int, err error) {
	if mke.errsByIdx == nil {
		mke.errsByIdx = make(map[int]error)
	}
	mke.errsByIdx[i] = err
}

func (mke *MultiKeyError) errAt(i int) error {
	if mke.errsByIdx == nil {
		return nil
	}
	err, ok := mke.errsByIdx[i]
	if !ok {
		return nil
	}
	return err
}
