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

func TestSetA(t *testing.T) {
	cpu := Cpu{AF: 0xffcc}
	cpu.setA(0xaa)

	if cpu.AF != 0xaacc {
		t.Errorf("Did not set A-register correctly. Expected 0xAACC, but got 0x%x", cpu.AF)
	}
}

func TestLoadAToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc}
	opcodes[0x7f](&cpu)

	if cpu.AF != 0xffcc {
		t.Errorf("Load A,A did not work correctly. Expected 0xFFCC, but got 0x%x", cpu.AF)
	}
}
