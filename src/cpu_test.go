package main

import "testing"

func TestHighByte(t *testing.T) {
	var rVal uint16 = 0xfe6c
	result := highByte(rVal)

	if result != 0xfe {
		t.Errorf("Did not correctly return the high byte. Expected 0xfe, but got 0x%X", result)
	}
}

func TestLowByte(t *testing.T) {
	var rVal uint16 = 0xfe6c
	result := lowByte(rVal)

	if result != 0x6c {
		t.Errorf("Did not correctly return the high byte. Expected 0x6c, but got 0x%X", result)
	}
}

func TestSetA(t *testing.T) {
	cpu := Cpu{AF: 0xffcc}
	cpu.setA(0xaa)

	if cpu.AF != 0xaacc {
		t.Errorf("Did not set A-register correctly. Expected 0xAACC, but got 0x%X", cpu.AF)
	}
}

func TestSetB(t *testing.T) {
	cpu := Cpu{BC: 0xffcc}
	cpu.setB(0xaa)

	if cpu.BC != 0xaacc {
		t.Errorf("Did not set B-register correctly. Expected 0xAACC, but got 0x%X", cpu.BC)
	}
}

func TestSetC(t *testing.T) {
	cpu := Cpu{BC: 0xffcc}
	cpu.setC(0xaa)

	if cpu.BC != 0xffaa {
		t.Errorf("Did not set C-register correctly. Expected 0xFFAA, but got 0x%X", cpu.BC)
	}
}

// Test LD A,A (opcode 0x7f)
func TestLoadAToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc}
	mem := Memory{}
	opcodes[0x7f](&cpu, &mem)

	if cpu.AF != 0xffcc {
		t.Errorf("Load A,A did not work correctly. Expected 0xFFCC, but got 0x%X", cpu.AF)
	}
}

// Test LD A,B (opcode 0x78)
func TestLoadBToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, BC: 0xab12}
	mem := Memory{}

	opcodes[0x78](&cpu, &mem)

	if cpu.AF != 0xabcc {
		t.Errorf("Load A,B did not work correctly. Expected 0xABCC, but got 0x%X", cpu.AF)
	}
}

// Test LD A,C (opcode 0x79)
func TestLoadCToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, BC: 0xab12}
	mem := Memory{}
	opcodes[0x79](&cpu, &mem)

	if cpu.AF != 0x12cc {
		t.Errorf("Load A,C did not work correctly. Expected 0x12CC, but got 0x%X", cpu.AF)
	}
}

// Test LD A,D (opcode 0x7a)
func TestLoadDToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x7a](&cpu, &mem)

	if cpu.AF != 0xaacc {
		t.Errorf("Load A,D did not work correctly. Expected 0xAACC, but got 0x%X", cpu.AF)
	}
}

// Test LD A,E (opcode 0x7b)
func TestLoadEToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x7b](&cpu, &mem)

	if cpu.AF != 0xbbcc {
		t.Errorf("Load A,E did not work correctly. Expected 0xBBCC, but got 0x%X", cpu.AF)
	}
}

// Test LD A,H (opcode 0x7c)
func TestLoadHToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, HL: 0xabcd}
	mem := Memory{}
	opcodes[0x7c](&cpu, &mem)

	if cpu.AF != 0xabcc {
		t.Errorf("Load A,H did not work correctly. Expected 0xABCC, but got 0x%X", cpu.AF)
	}
}

// Test LD A,L (opcode 0x7d)
func TestLoadLToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, HL: 0xabcd}
	mem := Memory{}
	opcodes[0x7d](&cpu, &mem)

	if cpu.AF != 0xcdcc {
		t.Errorf("Load A,L did not work correctly. Expected 0xCDCC, but got 0x%X", cpu.AF)
	}
}

// Test LD A,(HL) (opcode 0x7e)
func TestLoadHLToA(t *testing.T) {
	initOpCodes()
	cpu := Cpu{AF: 0xffcc, HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x7e](&cpu, &mem)

	if cpu.AF != 0xabcc {
		t.Errorf("Load A,(HL) did not work correctly. Expected 0xABCC but got 0x%X", cpu.AF)
	}
}

// Test LD B,B (opcode 0x40)
func TestLoadBToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc}
	mem := Memory{}
	opcodes[0x40](&cpu, &mem)

	if cpu.BC != 0xffcc {
		t.Errorf("Load B,B did not work correctly. Expected 0xFFCC but got 0x%X", cpu.BC)
	}
}

// Test LD B,C (opcode 0x41)
func TestLoadCToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc}
	mem := Memory{}
	opcodes[0x41](&cpu, &mem)

	if cpu.BC != 0xcccc {
		t.Errorf("Load B,C did not work correctly. Expected 0xCCCC but got 0x%X", cpu.BC)
	}
}

// Test LD B,D (opcode 0x42)
func TestLoadDToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, DE: 0xbb88}
	mem := Memory{}
	opcodes[0x42](&cpu, &mem)

	if cpu.BC != 0xbbcc {
		t.Errorf("Load B,D did not work correctly. Expected 0xBBCC but got 0x%X", cpu.BC)
	}
}

// Test LD B,E (opcode 0x43)
func TestLoadEToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, DE: 0xbb88}
	mem := Memory{}
	opcodes[0x43](&cpu, &mem)

	if cpu.BC != 0x88cc {
		t.Errorf("Load B,E did not work correctly. Expected 0x88CC but got 0x%X", cpu.BC)
	}
}

// Test LD B,H (opcode 0x44)
func TestLoadHToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	mem := Memory{}
	opcodes[0x44](&cpu, &mem)

	if cpu.BC != 0xaacc {
		t.Errorf("Load B,H did not work correctly. Expected 0xAACC but got 0x%X", cpu.BC)
	}
}

// Test LD B,L (opcode 0x45)
func TestLoadLToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	mem := Memory{}
	opcodes[0x45](&cpu, &mem)

	if cpu.BC != 0xbbcc {
		t.Errorf("Load B,L did not work correctly. Expected 0xBBCC but got 0x%X", cpu.BC)
	}
}

// Test LD B,(HL) (opcode 0x46)
func TestLoadHLToB(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x46](&cpu, &mem)

	if cpu.BC != 0xabcc {
		t.Errorf("Load B,(HL) did not work correctly. Expected 0xABCC but got 0x%X", cpu.BC)
	}
}

// Test LD C,B (opcode 0x48)
func TestLoadBToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc}
	mem := Memory{}
	opcodes[0x48](&cpu, &mem)

	if cpu.BC != 0xffff {
		t.Errorf("Load C,B did not work correctly. Expected 0xFFFF but got 0x%X", cpu.BC)
	}
}

// Test LD C,C (opcode 0x49)
func TestLoadCToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc}
	mem := Memory{}
	opcodes[0x49](&cpu, &mem)

	if cpu.BC != 0xffcc {
		t.Errorf("Load C,C did not work correctly. Expected 0xFFFF but got 0x%X", cpu.BC)
	}
}

// Test LD C,D (opcode 0x4a)
func TestLoadDToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x4a](&cpu, &mem)

	if cpu.BC != 0xffaa {
		t.Errorf("Load C,D did not work correctly. Expected 0xFFAA but got 0x%X", cpu.BC)
	}
}

// Test LD C,E (opcode 0x4b)
func TestLoadEToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x4b](&cpu, &mem)

	if cpu.BC != 0xffbb {
		t.Errorf("Load C,E did not work correctly. Expected 0xFFBB but got 0x%X", cpu.BC)
	}
}

// Test LD C,H (opcode 0x4c)
func TestLoadHToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	mem := Memory{}
	opcodes[0x4c](&cpu, &mem)

	if cpu.BC != 0xffaa {
		t.Errorf("Load C,H did not work correctly. Expected 0xFFAA but got 0x%X", cpu.BC)
	}
}

// Test LD C,L (opcode 0x4d)
func TestLoadLToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0xaabb}
	mem := Memory{}
	opcodes[0x4d](&cpu, &mem)

	if cpu.BC != 0xffbb {
		t.Errorf("Load C,L did not work correctly. Expected 0xFFBB but got 0x%X", cpu.BC)
	}
}

// Test LD C,(HL) (opcode 0x4e)
func TestLoadHLToC(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xffcc, HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x4e](&cpu, &mem)

	if cpu.BC != 0xffab {
		t.Errorf("Load C,(HL) did not work correctly. Expected 0xFFAB but got 0x%X", cpu.BC)
	}
}

// Test LD D,B (opcode 0x50)
func TestLoadBToD(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xddcc, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x50](&cpu, &mem)

	if cpu.DE != 0xddbb {
		t.Errorf("Load D,B did not work correctly. Expected 0xDDBB but got 0x%X", cpu.DE)
	}
}

// Test LD D,C (opcode 0x51)
func TestLoadCToD(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xddcc, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x51](&cpu, &mem)

	if cpu.DE != 0xccbb {
		t.Errorf("Load D,C did not work correctly. Expected 0xCCBB but got 0x%X", cpu.DE)
	}
}

// Test LD D,D (opcode 0x52)
func TestLoadDToD(t *testing.T) {
	initOpCodes()
	cpu := Cpu{DE: 0xaabb}
	mem := Memory{}
	opcodes[0x52](&cpu, &mem)

	if cpu.DE != 0xaabb {
		t.Errorf("Load D,D did not work correctly. Expected 0xAABB but got 0x%X", cpu.DE)
	}
}

// Test LD D,E (opcode 0x53)
func TestLoadEToD(t *testing.T) {
	initOpCodes()
	cpu := Cpu{DE: 0xaabb}
	mem := Memory{}
	opcodes[0x53](&cpu, &mem)

	if cpu.DE != 0xbbbb {
		t.Errorf("Load D,E did not work correctly. Expected 0xBBBB but got 0x%X", cpu.DE)
	}
}

// Test LD D,H (opcode 0x54)
func TestLoadHToD(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xddff, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x54](&cpu, &mem)

	if cpu.DE != 0xddbb {
		t.Errorf("Load D,H did not work correctly. Expected 0xDDBB but got 0x%X", cpu.DE)
	}
}

// Test LD D,L (opcode 0x55)
func TestLoadLToD(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xddff, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x55](&cpu, &mem)

	if cpu.DE != 0xffbb {
		t.Errorf("Load D,L did not work correctly. Expected 0xFFBB but got 0x%X", cpu.DE)
	}
}

// Test LD D,(HL) (opcode 0x56)
func TestLoadHLToD(t *testing.T) {
	initOpCodes()
	cpu := Cpu{DE: 0xffcc, HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x56](&cpu, &mem)

	if cpu.DE != 0xabcc {
		t.Errorf("Load D,(HL) did not work correctly. Expected 0xABCC but got 0x%X", cpu.DE)
	}
}

// Test LD E,B (opcode 0x58)
func TestLoadBToE(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xddcc, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x58](&cpu, &mem)

	if cpu.DE != 0xaadd {
		t.Errorf("Load D,E did not work correctly. Expected 0xAADD but got 0x%X", cpu.DE)
	}
}

// Test LD E,C (opcode 0x59)
func TestLoadCtoE(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xddcc, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x59](&cpu, &mem)

	if cpu.DE != 0xaacc {
		t.Errorf("Load E,C did not work correctly. Expected 0xAACC but got 0x%X", cpu.DE)
	}
}

// Test LD E,D (opcode 0x5a)
func TestLoadDtoE(t *testing.T) {
	initOpCodes()
	cpu := Cpu{DE: 0xaabb}
	mem := Memory{}
	opcodes[0x5a](&cpu, &mem)

	if cpu.DE != 0xaaaa {
		t.Errorf("Load E,D did not work correctly. Expected 0xAAAA but got 0x%X", cpu.DE)
	}
}

// Test LD E,E (opcode 0x5b)
func TestLoadEtoE(t *testing.T) {
	initOpCodes()
	cpu := Cpu{DE: 0xaabb}
	mem := Memory{}
	opcodes[0x5b](&cpu, &mem)

	if cpu.DE != 0xaabb {
		t.Errorf("Load E,E did not work correctly. Expected 0xAABB but got 0x%X", cpu.DE)
	}
}

// Test LD E,H (opcode 0x5c)
func TestLoadHtoE(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xddff, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x5c](&cpu, &mem)

	if cpu.DE != 0xaadd {
		t.Errorf("Load E,H did not work correctly. Expected 0xAADD but got 0x%X", cpu.DE)
	}
}

// Test LD E,L (opcode 0x5d)
func TestLoadLtoE(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xddff, DE: 0xaabb}
	mem := Memory{}
	opcodes[0x5d](&cpu, &mem)

	if cpu.DE != 0xaaff {
		t.Errorf("Load E,L did not work correctly. Expected 0xAAFF but got 0x%X", cpu.DE)
	}
}

// Test LD E,(HL) (opcode 0x5e)
func TestLoadHLtoE(t *testing.T) {
	initOpCodes()
	cpu := Cpu{DE: 0xffcc, HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x5e](&cpu, &mem)

	if cpu.DE != 0xffab {
		t.Errorf("Load E,(HL) did not work correctly. Expected 0xFFAB but got 0x%X", cpu.DE)
	}
}

// Test LD H,B (opcode 0x60)
func TestLoadBToH(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xddcc, HL: 0xaabb}
	mem := Memory{}
	opcodes[0x60](&cpu, &mem)

	if cpu.HL != 0xddbb {
		t.Errorf("Load H,B did not work correctly. Expected 0xDDBB but got 0x%X", cpu.HL)
	}
}

// Test LD H,C (opcode 0x61)
func TestLoadCtoH(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xddcc, HL: 0xaabb}
	mem := Memory{}
	opcodes[0x61](&cpu, &mem)

	if cpu.HL != 0xccbb {
		t.Errorf("Load H,C did not work correctly. Expected 0xCCBB but got 0x%X", cpu.HL)
	}
}

// Test LD H,D (opcode 0x62)
func TestLoadDtoH(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xaabb, DE: 0xeeff}
	mem := Memory{}
	opcodes[0x62](&cpu, &mem)

	if cpu.HL != 0xeebb {
		t.Errorf("Load H,D did not work correctly. Expected 0xEEDD but got 0x%X", cpu.HL)
	}
}

// Test LD H,E (opcode 0x63)
func TestLoadEtoH(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xaabb, DE: 0xeeff}
	mem := Memory{}
	opcodes[0x63](&cpu, &mem)

	if cpu.HL != 0xffbb {
		t.Errorf("Load H,E did not work correctly. Expected 0xFFBB but got 0x%X", cpu.HL)
	}
}

// Test LD H,H (opcode 0x64)
func TestLoadHtoH(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xaabb}
	mem := Memory{}
	opcodes[0x64](&cpu, &mem)

	if cpu.HL != 0xaabb {
		t.Errorf("Load H,H did not work correctly. Expected 0xAAbb but got 0x%X", cpu.HL)
	}
}

// Test LD H,L (opcode 0x65)
func TestLoadLtoH(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xaabb}
	mem := Memory{}
	opcodes[0x65](&cpu, &mem)

	if cpu.HL != 0xbbbb {
		t.Errorf("Load H,L did not work correctly. Expected 0xBBBB but got 0x%X", cpu.HL)
	}
}

// Test LD H,(HL) (opcode 0x66)
func TestLoadHLtoH(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x66](&cpu, &mem)

	if cpu.HL != 0xab12 {
		t.Errorf("Load H,(HL) did not work correctly. Expected 0xAB12 but got 0x%X", cpu.HL)
	}
}

// Test LD L,B (opcode 0x68)
func TestLoadBToL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xddcc, HL: 0xaabb}
	mem := Memory{}
	opcodes[0x68](&cpu, &mem)

	if cpu.HL != 0xaadd {
		t.Errorf("Load L,B did not work correctly. Expected 0xAADD but got 0x%X", cpu.HL)
	}
}

// Test LD L,C (opcode 0x69)
func TestLoadCtoL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{BC: 0xddcc, HL: 0xaabb}
	mem := Memory{}
	opcodes[0x69](&cpu, &mem)

	if cpu.HL != 0xaacc {
		t.Errorf("Load L,C did not work correctly. Expected 0xAACC but got 0x%X", cpu.HL)
	}
}

// Test LD L,D (opcode 0x6a)
func TestLoadDtoL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xaabb, DE: 0xeeff}
	mem := Memory{}
	opcodes[0x6a](&cpu, &mem)

	if cpu.HL != 0xaaee {
		t.Errorf("Load L,D did not work correctly. Expected 0xAAEE but got 0x%X", cpu.HL)
	}
}

// Test LD L,E (opcode 0x6b)
func TestLoadEtoL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xaabb, DE: 0xeeff}
	mem := Memory{}
	opcodes[0x6b](&cpu, &mem)

	if cpu.HL != 0xaaff {
		t.Errorf("Load L,E did not work correctly. Expected 0xAAFF but got 0x%X", cpu.HL)
	}
}

// Test LD L,H (opcode 0x6c)
func TestLoadHtoL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xaabb}
	mem := Memory{}
	opcodes[0x6c](&cpu, &mem)

	if cpu.HL != 0xaaaa {
		t.Errorf("Load L,H did not work correctly. Expected 0xAAAA but got 0x%X", cpu.HL)
	}
}

// Test LD L,L (opcode 0x6d)
func TestLoadLtoL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0xaabb}
	mem := Memory{}
	opcodes[0x6d](&cpu, &mem)

	if cpu.HL != 0xaabb {
		t.Errorf("Load L,L did not work correctly. Expected 0xAABB but got 0x%X", cpu.HL)
	}
}

// Test LD L,(HL) (opcode 0x6e)
func TestLoadHLtoL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x6e](&cpu, &mem)

	if cpu.HL != 0x00ab {
		t.Errorf("Load L,(HL) did not work correctly. Expected 0x00AB but got 0x%X", cpu.HL)
	}
}

// Test LD (HL),B (opcode 0x70)
func TestLoadBToHL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012, BC: 0xaabb}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x70](&cpu, &mem)

	if mem.ram[0x0012] != 0xaa {
		t.Errorf("Load (HL),B did not work correctly. Expected 0xAA but got 0x%X", mem.ram[0x0012])
	}
}

// Test LD (HL),C (opcode 0x71)
func TestLoadCtoHL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012, BC: 0xaabb}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x71](&cpu, &mem)

	if mem.ram[0x0012] != 0xbb {
		t.Errorf("Load (HL),C did not work correctly. Expected 0xBB but got 0x%X", mem.ram[0x0012])
	}
}

// Test LD (HL),D (opcode 0x72)
func TestLoadDtoHL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012, DE: 0xaabb}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x72](&cpu, &mem)

	if mem.ram[0x0012] != 0xaa {
		t.Errorf("Load (HL),D did not work correctly. Expected 0xAA but got 0x%X", mem.ram[0x0012])
	}
}

// Test LD (HL),E (opcode 0x73)
func TestLoadEtoHL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012, DE: 0xaabb}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x73](&cpu, &mem)

	if mem.ram[0x0012] != 0xbb {
		t.Errorf("Load (HL),E did not work correctly. Expected 0xBB but got 0x%X", mem.ram[0x0012])
	}
}

// Test LD (HL),H (opcode 0x74)
func TestLoadHtoHL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x74](&cpu, &mem)

	if mem.ram[0x0012] != 0x00 {
		t.Errorf("Load (HL),H did not work correctly. Expected 0x00 but got 0x%X", mem.ram[0x0012])
	}
}

// Test LD (HL),L (opcode 0x75)
func TestLoadLtoHL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012}
	ram := [20]uint8{0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x75](&cpu, &mem)

	if mem.ram[0x0012] != 0x12 {
		t.Errorf("Load (HL),L did not work correctly. Expected 0x12 but got 0x%X", mem.ram[0x0012])
	}
}

// Test LD (HL),n (opcode 0x36)
func TestLoadValtoHL(t *testing.T) {
	initOpCodes()
	cpu := Cpu{HL: 0x0012}
	ram := [20]uint8{0x0: 0x36, 0x0001: 0xee, 0x0012: 0xab}
	mem := Memory{ram[:]}

	opcodes[0x36](&cpu, &mem)

	if mem.ram[0x0012] != 0xee {
		t.Errorf("Load (HL),0xEE did not work correctly. Expected 0xEE but got 0x%X", mem.ram[0x0012])
	}
}

// Test LD B,n
func TestLoadValToB(t *testing.T) {
	cpu := Cpu{BC: 0xaabb, PC: 0x0000}
	mem := Memory{ram: []uint8{0x06, 0x1a}}
	opcodes[0x06](&cpu, &mem)

	if cpu.BC != 0x1abb {
		t.Errorf("Load B,0x1a did not work correctly. Expected 0x1abb but got 0x%X", cpu.BC)
	}
}

// Test LD C,n
func TestLoadValToC(t *testing.T) {
	cpu := Cpu{BC: 0xaabb, PC: 0x0000}
	mem := Memory{ram: []uint8{0x0e, 0x1a}}
	opcodes[0x0e](&cpu, &mem)

	if cpu.BC != 0xaa1a {
		t.Errorf("Load C,0x1a did not work correctly. Expected 0xaa1a but got 0x%X", cpu.BC)
	}
}

// Test LD D,n
func TestLoadValToD(t *testing.T) {
	cpu := Cpu{DE: 0xaabb, PC: 0x0000}
	mem := Memory{ram: []uint8{0x16, 0x1a}}
	opcodes[0x16](&cpu, &mem)

	if cpu.DE != 0x1abb {
		t.Errorf("Load D,0x1a did not work correctly. Expected 0x1abb but got 0x%X", cpu.DE)
	}
}

// Test LD E,n
func TestLoadValToE(t *testing.T) {
	cpu := Cpu{DE: 0xaabb, PC: 0x0000}
	mem := Memory{ram: []uint8{0x1e, 0x1a}}
	opcodes[0x1e](&cpu, &mem)

	if cpu.DE != 0xaa1a {
		t.Errorf("Load E,0x1a did not work correctly. Expected 0xaa1a but got 0x%X", cpu.DE)
	}
}

// Test LD H,n
func TestLoadValToH(t *testing.T) {
	cpu := Cpu{HL: 0xaabb, PC: 0x0000}
	mem := Memory{ram: []uint8{0x26, 0x1a}}
	opcodes[0x26](&cpu, &mem)

	if cpu.HL != 0x1abb {
		t.Errorf("Load H,0x1a did not work correctly. Expected 0x1abb but got 0x%X", cpu.HL)
	}
}

// Test LD L,n
func TestLoadValToL(t *testing.T) {
	cpu := Cpu{HL: 0xaabb, PC: 0x0000}
	mem := Memory{ram: []uint8{0x2e, 0x1a}}
	opcodes[0x2e](&cpu, &mem)

	if cpu.HL != 0xaa1a {
		t.Errorf("Load L,0x1a did not work correctly. Expected 0xaa1a but got 0x%X", cpu.HL)
	}
}
