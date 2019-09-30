package main

import (
	"encoding/binary"
	"testing"
)

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

}

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
	cnt, _, err := db.Get([]byte("count"), nil)
	if err != nil {
		panic("bad news")
	}
	println(string(cnt))

	len, _, err := db.Get([]byte("len"), nil)
	println(string(len) + " byte")
}
