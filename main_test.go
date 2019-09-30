package main

import (
	"testing"
)

func Test_Store(t *testing.T) {

	db := NewKuku()
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
