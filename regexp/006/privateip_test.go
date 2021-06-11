package main

import (
	"testing"
)

func TestPrivateIP(t *testing.T) {
	//t.Fatal("not implemented")
	result_private, _ := PrivateIP("192.168.10.1")
	if result_private == false {
		t.Errorf("192.168.10.1 is PrivateIP")
		return
	}
	result_global, _ := PrivateIP("1.1.1.1")
	if result_global == true {
		t.Errorf("1.1.1.1 is GlobalIP")
		return
	}
}
