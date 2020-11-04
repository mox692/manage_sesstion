package main

import "testing"

func TestRunRedis(t *testing.T) {
	err := runRedis()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckENV(t *testing.T) {
	if !checkENV() {
		t.Errorf("errro!")
	}
}
