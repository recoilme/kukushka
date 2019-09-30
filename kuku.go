package main

import (
	"bufio"
	"bytes"
	"fmt"
	"sync"

	"github.com/recoilme/mcproto"
	cuckoo "github.com/seiflotfy/cuckoofilter"
)

type kuku struct {
	sync.RWMutex
	cf *cuckoo.Filter
}

// newKuku - init filter
func newKuku() mcproto.McEngine {
	eng := &kuku{}
	eng.Lock()
	defer eng.Unlock()
	eng.cf = cuckoo.NewFilter(10_000_000)
	return eng
}

func (en *kuku) Get(key []byte, rw *bufio.ReadWriter) (value []byte, noreply bool, err error) {
	en.RLock()
	defer en.RUnlock()

	if bytes.HasPrefix(key, []byte("count")) {
		return []byte(fmt.Sprintf("%d", en.cf.Count())), false, nil
	}
	if bytes.HasPrefix(key, []byte("len")) {
		return []byte(fmt.Sprintf("%d", len(en.cf.Encode()))), false, nil
	}
	val := "0"
	if en.cf.Lookup(key) {
		val = "1"
	}
	return []byte(val), false, nil
}
func (en *kuku) Gets(keys [][]byte, rw *bufio.ReadWriter) (err error) {
	return
}
func (en *kuku) Set(key, value []byte, flags uint32, exp int32, size int, noreply bool, rw *bufio.ReadWriter) (noreplyresp bool, err error) {
	en.Lock()
	defer en.Unlock()
	en.cf.InsertUnique(key)
	return
}
func (en *kuku) Incr(key []byte, value uint64, rw *bufio.ReadWriter) (result uint64, isFound bool, noreply bool, err error) {
	return
}
func (en *kuku) Decr(key []byte, value uint64, rw *bufio.ReadWriter) (result uint64, isFound bool, noreply bool, err error) {
	return
}
func (en *kuku) Delete(key []byte, rw *bufio.ReadWriter) (isFound bool, noreply bool, err error) {
	return
}
func (en *kuku) Close() (err error) {
	return
}
