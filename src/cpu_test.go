package main

import "testing"

func TestHighByte(t *testing.T) {
	var rVal uint16 = 0xfe6c
	result := highByte(rVal)

	if result != 0xfe {
		t.Errorf("Did not correctly return the high byte. Expected 0xfe, but got 0x%x", result)
	}
}

func TestLowByte(t *testing.T) {
	var rVal uint16 = 0xfe6c
	result := lowByte(rVal)

	if result != 0x6c {
		t.Errorf("Did not correctly return the high byte. Expected 0x6c, but got 0x%x", result)
	}
}
