package main

import (
	"encoding/binary"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/cespare/xxhash"
)

/*
func Test_Store(t *testing.T) {

	db := newKuku()
	db.Set([]byte("1"), nil, 0, 0, 1, false, nil)
	val, _, _ := db.Get([]byte("1"), nil)
	if string(val) != "1" {
		t.Errorf("Expected 1, got:%s", string(val))
	}
	valmis, _, _ := db.Get([]byte("mis"), nil)
	if string(valmis) != "0" {
		t.Errorf("Expected 0, got:%s", valmis)
	}

}*/

func Test_Map(t *testing.T) {
	println("map memory usage")
	m := make(map[uint64]uint64)
	for i := 0; i < 1_000_000; i++ {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(i))
		h := xxhash.Sum64(b)
		m[h] = 1
	}

	println(len(m))
	PrintMemUsage()
	for j := 0; j < 1_000_000; j++ {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(j))
		h := xxhash.Sum64(b)
		if _, ok := m[h]; !ok {
			println("false")
		}

	}
	runtime.GC()
	time.Sleep(time.Second)
}

/*
func Test_Size(t *testing.T) {
	db := newKuku()
	for i := 0; i < 1_000_000; i++ {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(i))
		_, err := db.Set(b, nil, 0, 0, 0, false, nil)
		if err != nil {
			panic("bad news")
		}
	}
	println("kuku memory usage")
	PrintMemUsage()
	cnt, _, err := db.Get([]byte("count"), nil)
	if err != nil {
		panic("bad news")
	}
	println(string(cnt))

	len, _, err := db.Get([]byte("len"), nil)
	println(string(len) + " byte")
}
*/
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
