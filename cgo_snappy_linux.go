// +build linux

package levigo

// #cgo CFLAGS: -I${SRCDIR}/vendor/snappy -DSNAPPY
// #cgo CXXFLAGS: -I${SRCDIR}/vendor/snappy -std=c++11 -g -O2 -fPIC -DPIC -DSNAPPY
import "C"
