package main

import "testing"

func TestSwap(t *testing.T) {
	err := swapFiles("", "")
	if err != nil {
		t.Errorf("swap != nil")
	}
}
