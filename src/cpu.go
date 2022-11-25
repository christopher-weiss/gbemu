package main

// The CPU in the original Gameboy has eight general purpose 8-bit registers (A,..,F,H,L).
// and two 16 bit registers acting as the program counter (PC) and stack pointer (SP).
//
// == A register ==
// The A (accumulator) register is used in arithmetic operations.
//
// == B,C,D,E,H,L registers ==
// These are general purpose registers, to be used as necessary.
//
// == F register ==
// The F (flag) register indicates the outcome of the previously
// executed operation. The bits in the flag register are as follows:
//
//  ------------------------
// | index: 7 6 5 4 3 2 1 0 |
// | value: Z N H C 0 0 0 0 |
//  ------------------------
// Z - Zero flag
// N - Subtraction flag
// H - Half-carry flag
// C - Carry flag
//
// There is a quirk to the registers though, in that some can act in pairs to function as a single 16 bit register.
// These are: AF, BC, DE and HL. For this reason we will use a 16-bit unsigned integer to represent a pair of flags.

type Cpu struct {
	// registers
	AF uint16
	BC uint16
	DE uint16
	HL uint16

	// program counter and stack pointer
	PC uint16
	SP uint16
}

func init() {
	initOpCodes()
}

// HighByte returns the value of the first byte of combined register value e.g. Return the A value of the AF-register.
func highByte(rVal uint16) uint8 {
	return uint8(rVal >> 8)
}

// HighByte returns the value of the first byte of combined register value e.g. Return the F value of the AF-register.
func lowByte(rVal uint16) uint8 {
	return uint8(rVal & 0xff)
}

// Set the A-register to the given value.
func (cpu *Cpu) setA(val uint8) {
	cpu.AF = cpu.AF & ((uint16(val) << 8) | 0xff)
}

// Set the B-register to the given value.
func (cpu *Cpu) setB(val uint8) {
	cpu.BC = cpu.BC & ((uint16(val) << 8) | 0xff)
}

// Set the C-register to the given value.
func (cpu *Cpu) setC(val uint8) {
	cpu.BC = (cpu.BC & 0xff00) | uint16(val)
}

var opcodes map[uint8]func(*Cpu)

// Initalize opcodes map
func initOpCodes() {
	opcodes = make(map[uint8]func(*Cpu))

	//
	// Load (LD) r1,r2
	// Take value stored in r2 register and place it in r1
	//

	// LD A,A
	opcodes[0x7f] = func(cpu *Cpu) {
		cpu.setA(highByte(cpu.AF))
	}

	// LD A,B
	opcodes[0x78] = func(cpu *Cpu) {
		cpu.setA(highByte(cpu.BC))
	}

	// LD A,C
	opcodes[0x79] = func(cpu *Cpu) {
		cpu.setA(lowByte(cpu.BC))
	}

	// LD A,D
	opcodes[0x7a] = func(cpu *Cpu) {
		cpu.setA(highByte(cpu.DE))
	}

	// LD A,E
	opcodes[0x7b] = func(cpu *Cpu) {
		cpu.setA(lowByte(cpu.DE))
	}

	// LD A,H
	opcodes[0x7c] = func(cpu *Cpu) {
		cpu.setA(highByte(cpu.HL))
	}

	// LD A,L
	opcodes[0x7d] = func(cpu *Cpu) {
		cpu.setA(lowByte(cpu.HL))
	}

	// LD A,(HL)
	// Functionally identical to LD A,H (opcode 0x7c)
	// however it takes 8 clock cycles to execute
	opcodes[0x7e] = func(cpu *Cpu) {
		cpu.setA(highByte(cpu.HL))
	}

	// LD B,B
	opcodes[0x40] = func(cpu *Cpu) {
		cpu.setB(highByte(cpu.BC))
	}

	// LD B,C
	opcodes[0x41] = func(cpu *Cpu) {
		cpu.setB(lowByte(cpu.BC))
	}

	// LD B,D
	opcodes[0x42] = func(cpu *Cpu) {
		cpu.setB(highByte(cpu.DE))
	}

	// LD B,E
	opcodes[0x43] = func(cpu *Cpu) {
		cpu.setB(lowByte(cpu.DE))
	}

	// LD B,H
	opcodes[0x44] = func(cpu *Cpu) {
		cpu.setB(highByte(cpu.HL))
	}

	// LD B,L
	opcodes[0x45] = func(cpu *Cpu) {
		cpu.setB(lowByte(cpu.HL))
	}

	// LD B,(HL)
	// Functionally identical to LD B,H (opcode 0x44)
	// however it takes 8 clock cycles to execute
	opcodes[0x46] = func(cpu *Cpu) {
		cpu.setB(highByte(cpu.HL))
	}

	// LD C,B
	opcodes[0x48] = func(cpu *Cpu) {
		cpu.setC(highByte(cpu.BC))
	}

	// LD C,C
	opcodes[0x49] = func(cpu *Cpu) {
		cpu.setC(lowByte(cpu.BC))
	}

	// LD C,D
	opcodes[0x4a] = func(cpu *Cpu) {
		cpu.setC(highByte(cpu.DE))
	}

	// LD C,E
	opcodes[0x4b] = func(cpu *Cpu) {
		cpu.setC(lowByte(cpu.DE))
	}

	// LD C,H
	opcodes[0x4c] = func(cpu *Cpu) {
		cpu.setC(highByte(cpu.HL))
	}

	// LD C,L
	opcodes[0x4d] = func(cpu *Cpu) {
		cpu.setC(lowByte(cpu.HL))
	}

	// LD C,(HL)
	// Functionally identical to LD C,H (opcode 0x4c)
	// however it takes 8 clock cycles to execute
	opcodes[0x4e] = func(cpu *Cpu) {
		cpu.setC(highByte(cpu.HL))
	}
}
