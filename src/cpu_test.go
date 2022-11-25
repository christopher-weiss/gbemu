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

func TestSetB(t *testing.T) {
	cpu := Cpu{BC: 0xffcc}
	cpu.setB(0xaa)

	if cpu.BC != 0xaacc {
		t.Errorf("Did not set B-register correctly. Expected 0xAACC, but got 0x%x", cpu.BC)
	}
}

func TestSetC(t *testing.T) {
	cpu := Cpu{BC: 0xffcc}
	cpu.setC(0xaa)

	if cpu.BC != 0xffaa {
		t.Errorf("Did not set C-register correctly. Expected 0xFFAA, but got 0x%x", cpu.BC)
	}
}

// Test LD A,A (opcode 0x7f)
func TestLoadAToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc}
	opcodes[0x7f](&cpu)

	if cpu.AF != 0xffcc {
		t.Errorf("Load A,A did not work correctly. Expected 0xFFCC, but got 0x%x", cpu.AF)
	}
}

// Test LD A,B (opcode 0x78)
func TestLoadBToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, BC: 0xab12}
	opcodes[0x78](&cpu)

	if cpu.AF != 0xabcc {
		t.Errorf("Load A,B did not work correctly. Expected 0xABCC, but got 0x%x", cpu.AF)
	}
}

// Test LD A,C (opcode 0x79)
func TestLoadCToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, BC: 0xab12}
	opcodes[0x79](&cpu)

	if cpu.AF != 0x12cc {
		t.Errorf("Load A,C did not work correctly. Expected 0x12CC, but got 0x%x", cpu.AF)
	}
}

// Test LD A,D (opcode 0x7a)
func TestLoadDToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, DE: 0xaabb}
	opcodes[0x7a](&cpu)

	if cpu.AF != 0xaacc {
		t.Errorf("Load A,D did not work correctly. Expected 0xAACC, but got 0x%x", cpu.AF)
	}
}

// Test LD A,E (opcode 0x7b)
func TestLoadEToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, DE: 0xaabb}
	opcodes[0x7b](&cpu)

	if cpu.AF != 0xbbcc {
		t.Errorf("Load A,E did not work correctly. Expected 0xBBCC, but got 0x%x", cpu.AF)
	}
}

// Test LD A,H (opcode 0x7c)
func TestLoadHToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, HL: 0xabcd}
	opcodes[0x7c](&cpu)

	if cpu.AF != 0xabcc {
		t.Errorf("Load A,H did not work correctly. Expected 0xABCC, but got 0x%x", cpu.AF)
	}
}

// Test LD A,L (opcode 0x7d)
func TestLoadLToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, HL: 0xabcd}
	opcodes[0x7d](&cpu)

	if cpu.AF != 0xcdcc {
		t.Errorf("Load A,L did not work correctly. Expected 0xCDCC, but got 0x%x", cpu.AF)
	}
}

// Test LD A,(HL) (opcode 0x7e)
func TestLoadHLToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, HL: 0xabcd}
	opcodes[0x7e](&cpu)

	if cpu.AF != 0xabcc {
		t.Errorf("Load A,(HL) did not work correctly. Expected 0xABCC but got 0x%x", cpu.AF)
	}
}

// Test LD B,B (opcode 0x40)
func TestLoadBToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc}
	opcodes[0x40](&cpu)

	if cpu.BC != 0xffcc {
		t.Errorf("Load B,B did not work correctly. Expected 0xFFCC but got 0x%x", cpu.BC)
	}
}

// Test LD B,C (opcode 0x41)
func TestLoadCToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc}
	opcodes[0x41](&cpu)

	if cpu.BC != 0xcccc {
		t.Errorf("Load B,C did not work correctly. Expected 0xCCCC but got 0x%x", cpu.BC)
	}
}

// Test LD B,D (opcode 0x42)
func TestLoadDToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, DE: 0xbb88}
	opcodes[0x42](&cpu)

	if cpu.BC != 0xbbcc {
		t.Errorf("Load B,D did not work correctly. Expected 0xBBCC but got 0x%x", cpu.BC)
	}
}

// Test LD B,E (opcode 0x43)
func TestLoadEToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, DE: 0xbb88}
	opcodes[0x43](&cpu)

	if cpu.BC != 0x88cc {
		t.Errorf("Load B,E did not work correctly. Expected 0x88CC but got 0x%x", cpu.BC)
	}
}

// Test LD B,H (opcode 0x44)
func TestLoadHToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	opcodes[0x44](&cpu)

	if cpu.BC != 0xaacc {
		t.Errorf("Load B,H did not work correctly. Expected 0xAACC but got 0x%x", cpu.BC)
	}
}

// Test LD B,L (opcode 0x45)
func TestLoadLToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	opcodes[0x45](&cpu)

	if cpu.BC != 0xbbcc {
		t.Errorf("Load B,L did not work correctly. Expected 0xBBCC but got 0x%x", cpu.BC)
	}
}

// Test LD B,(HL) (opcode 0x46)
func TestLoadHLToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	opcodes[0x46](&cpu)

	if cpu.BC != 0xaacc {
		t.Errorf("Load B,(HL) did not work correctly. Expected 0xAACC but got 0x%x", cpu.BC)
	}
}

// Test LD C,B (opcode 0x48)
func TestLoadBToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc}
	opcodes[0x48](&cpu)

	if cpu.BC != 0xffff {
		t.Errorf("Load C,B did not work correctly. Expected 0xFFFF but got 0x%x", cpu.BC)
	}
}

// Test LD C,C (opcode 0x49)
func TestLoadCToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc}
	opcodes[0x49](&cpu)

	if cpu.BC != 0xffcc {
		t.Errorf("Load C,C did not work correctly. Expected 0xFFFF but got 0x%x", cpu.BC)
	}
}

// Test LD C,D (opcode 0x4a)
func TestLoadDToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, DE: 0xaabb}
	opcodes[0x4a](&cpu)

	if cpu.BC != 0xffaa {
		t.Errorf("Load C,D did not work correctly. Expected 0xFFAA but got 0x%x", cpu.BC)
	}
}

// Test LD C,E (opcode 0x4b)
func TestLoadEToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, DE: 0xaabb}
	opcodes[0x4b](&cpu)

	if cpu.BC != 0xffbb {
		t.Errorf("Load C,E did not work correctly. Expected 0xFFBB but got 0x%x", cpu.BC)
	}
}

// Test LD C,H (opcode 0x4c)
func TestLoadHToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	opcodes[0x4c](&cpu)

	if cpu.BC != 0xffaa {
		t.Errorf("Load C,H did not work correctly. Expected 0xFFAA but got 0x%x", cpu.BC)
	}
}

// Test LD C,L (opcode 0x4d)
func TestLoadLToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	opcodes[0x4d](&cpu)

	if cpu.BC != 0xffbb {
		t.Errorf("Load C,L did not work correctly. Expected 0xFFBB but got 0x%x", cpu.BC)
	}
}

// Test LD C,(HL) (opcode 0x4e)
func TestLoadHLToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	opcodes[0x4e](&cpu)

	if cpu.BC != 0xffaa {
		t.Errorf("Load C,(HL) did not work correctly. Expected 0xFFAA but got 0x%x", cpu.BC)
	}
}
