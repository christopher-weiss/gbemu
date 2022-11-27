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
	cpu.AF = (cpu.AF & 0x00ff) | (uint16(val) << 8)
}

// Set the B-register to the given value.
func (cpu *Cpu) setB(val uint8) {
	cpu.BC = (cpu.BC & 0x00ff) | (uint16(val) << 8)
}

// Set the C-register to the given value.
func (cpu *Cpu) setC(val uint8) {
	cpu.BC = (cpu.BC & 0xff00) | uint16(val)
}

// Set the D-register to the given value.
func (cpu *Cpu) setD(val uint8) {
	cpu.DE = (cpu.DE & 0x00ff) | (uint16(val) << 8)
}

// Set the E-register to the given value.
func (cpu *Cpu) setE(val uint8) {
	cpu.DE = (cpu.DE & 0xff00) | uint16(val)
}

// Set the H-register to the given value.
func (cpu *Cpu) setH(val uint8) {
	cpu.HL = (cpu.HL & 0x00ff) | (uint16(val) << 8)
}

// Set the L-register to the given value.
func (cpu *Cpu) setL(val uint8) {
	cpu.HL = (cpu.HL & 0xff00) | uint16(val)
}

// Set the (HL)-register to the given value.
func (cpu *Cpu) setHL(val uint8) {
	cpu.HL = (cpu.HL & 0x00ff) | (uint16(val) << 8)
}

var opcodes map[uint8]func(*Cpu, *Memory)

// Initalize opcodes map
func initOpCodes() {
	opcodes = make(map[uint8]func(*Cpu, *Memory))

	//
	// Load (LD) r,n
	// Put immediate 8-bit value n into register r e.g. put value 0xfe into register D
	//
	// Since are word size is 8 bits and we need 5 bits to specify the operation and 3 bits to specify
	// the register we want to store to, the immediate value will be stored in the address following
	// the address of the operation
	//
	//  ------------------------
	// | index: 7 6 5 4 3 2 1 0 |
	// | value: 0 0 <-r-> 1 1 0 |
	//  ------------------------
	// where r refers to the register in the table
	//  --------------------------
	// | register | value (binary)|
	// |        A | 111           |
	// |        B | 000           |
	// |        C | 001           |
	// |        D | 101           |
	// |        E | 011           |
	// |        H | 100           |
	// |        L | 101           |
	//  --------------------------
	//
	// Example: "Load the value into register C" translates to opcode
	//          00<-r->110 = 00001110 = 0x0e

	// LD B,n
	opcodes[0x06] = func(cpu *Cpu, mem *Memory) {
		cpu.PC++
		val := mem.Read(cpu.PC)
		cpu.setB(val)
	}

	// LD C,n
	opcodes[0x0e] = func(cpu *Cpu, mem *Memory) {
		cpu.PC++
		val := mem.Read(cpu.PC)
		cpu.setC(val)
	}

	// LD D,n
	opcodes[0x16] = func(cpu *Cpu, mem *Memory) {
		cpu.PC++
		val := mem.Read(cpu.PC)
		cpu.setD(val)
	}

	// LD E,n
	opcodes[0x1e] = func(cpu *Cpu, mem *Memory) {
		cpu.PC++
		val := mem.Read(cpu.PC)
		cpu.setE(val)
	}

	// LD H,n
	opcodes[0x26] = func(cpu *Cpu, mem *Memory) {
		cpu.PC++
		val := mem.Read(cpu.PC)
		cpu.setH(val)
	}

	// LD L,n
	opcodes[0x2e] = func(cpu *Cpu, mem *Memory) {
		cpu.PC++
		val := mem.Read(cpu.PC)
		cpu.setL(val)
	}

	//
	// Load (LD) r1,r2
	// Take value stored in r2 register and place it in r1
	//

	// LD A,A
	opcodes[0x7f] = func(cpu *Cpu, mem *Memory) {
		cpu.setA(highByte(cpu.AF))
	}

	// LD A,B
	opcodes[0x78] = func(cpu *Cpu, mem *Memory) {
		cpu.setA(highByte(cpu.BC))
	}

	// LD A,C
	opcodes[0x79] = func(cpu *Cpu, mem *Memory) {
		cpu.setA(lowByte(cpu.BC))
	}

	// LD A,D
	opcodes[0x7a] = func(cpu *Cpu, mem *Memory) {
		cpu.setA(highByte(cpu.DE))
	}

	// LD A,E
	opcodes[0x7b] = func(cpu *Cpu, mem *Memory) {
		cpu.setA(lowByte(cpu.DE))
	}

	// LD A,H
	opcodes[0x7c] = func(cpu *Cpu, mem *Memory) {
		cpu.setA(highByte(cpu.HL))
	}

	// LD A,L
	opcodes[0x7d] = func(cpu *Cpu, mem *Memory) {
		cpu.setA(lowByte(cpu.HL))
	}

	// LD A,(HL)
	// (HL) means that the H & L registers are used as single 16-bit memory address, the
	// contents of which are stored to the A register
	opcodes[0x7e] = func(cpu *Cpu, mem *Memory) {
		cpu.setA(mem.Read(cpu.HL))
	}

	// LD B,B
	opcodes[0x40] = func(cpu *Cpu, mem *Memory) {
		cpu.setB(highByte(cpu.BC))
	}

	// LD B,C
	opcodes[0x41] = func(cpu *Cpu, mem *Memory) {
		cpu.setB(lowByte(cpu.BC))
	}

	// LD B,D
	opcodes[0x42] = func(cpu *Cpu, mem *Memory) {
		cpu.setB(highByte(cpu.DE))
	}

	// LD B,E
	opcodes[0x43] = func(cpu *Cpu, mem *Memory) {
		cpu.setB(lowByte(cpu.DE))
	}

	// LD B,H
	opcodes[0x44] = func(cpu *Cpu, mem *Memory) {
		cpu.setB(highByte(cpu.HL))
	}

	// LD B,L
	opcodes[0x45] = func(cpu *Cpu, mem *Memory) {
		cpu.setB(lowByte(cpu.HL))
	}

	// LD B,(HL)
	// (HL) means that the H & L registers are used as single 16-bit memory address, the
	// contents of which are stored to the B register
	opcodes[0x46] = func(cpu *Cpu, mem *Memory) {
		cpu.setB(mem.Read(cpu.HL))
	}

	// LD C,B
	opcodes[0x48] = func(cpu *Cpu, mem *Memory) {
		cpu.setC(highByte(cpu.BC))
	}

	// LD C,C
	opcodes[0x49] = func(cpu *Cpu, mem *Memory) {
		cpu.setC(lowByte(cpu.BC))
	}

	// LD C,D
	opcodes[0x4a] = func(cpu *Cpu, mem *Memory) {
		cpu.setC(highByte(cpu.DE))
	}

	// LD C,E
	opcodes[0x4b] = func(cpu *Cpu, mem *Memory) {
		cpu.setC(lowByte(cpu.DE))
	}

	// LD C,H
	opcodes[0x4c] = func(cpu *Cpu, mem *Memory) {
		cpu.setC(highByte(cpu.HL))
	}

	// LD C,L
	opcodes[0x4d] = func(cpu *Cpu, mem *Memory) {
		cpu.setC(lowByte(cpu.HL))
	}

	// LD C,(HL)
	// (HL) means that the H & L registers are used as single 16-bit memory address, the
	// contents of which are stored to the C register
	opcodes[0x4e] = func(cpu *Cpu, mem *Memory) {
		cpu.setC(mem.Read(cpu.HL))
	}

	// LD D,B
	opcodes[0x50] = func(cpu *Cpu, mem *Memory) {
		cpu.setD(highByte(cpu.BC))
	}

	// LD D,C
	opcodes[0x51] = func(cpu *Cpu, mem *Memory) {
		cpu.setD(lowByte(cpu.BC))
	}

	// LD D,D
	opcodes[0x52] = func(cpu *Cpu, mem *Memory) {
		cpu.setD(highByte(cpu.DE))
	}

	// LD D,E
	opcodes[0x53] = func(cpu *Cpu, mem *Memory) {
		cpu.setD(lowByte(cpu.DE))
	}

	// LD D,H
	opcodes[0x54] = func(cpu *Cpu, mem *Memory) {
		cpu.setD(highByte(cpu.HL))
	}

	// LD D,L
	opcodes[0x55] = func(cpu *Cpu, mem *Memory) {
		cpu.setD(lowByte(cpu.HL))
	}

	// LD D,(HL)
	// (HL) means that the H & L registers are used as single 16-bit memory address, the
	// contents of which are stored to the D register
	opcodes[0x56] = func(cpu *Cpu, mem *Memory) {
		cpu.setD(mem.Read(cpu.HL))
	}

	// LD E,B
	opcodes[0x58] = func(cpu *Cpu, mem *Memory) {
		cpu.setE(highByte(cpu.BC))
	}

	// LD E,C
	opcodes[0x59] = func(cpu *Cpu, mem *Memory) {
		cpu.setE(lowByte(cpu.BC))
	}

	// LD E,D
	opcodes[0x5a] = func(cpu *Cpu, mem *Memory) {
		cpu.setE(highByte(cpu.DE))
	}

	// LD E,E
	opcodes[0x5b] = func(cpu *Cpu, mem *Memory) {
		cpu.setE(lowByte(cpu.DE))
	}

	// LD E,H
	opcodes[0x5c] = func(cpu *Cpu, mem *Memory) {
		cpu.setE(highByte(cpu.HL))
	}

	// LD E,L
	opcodes[0x5d] = func(cpu *Cpu, mem *Memory) {
		cpu.setE(lowByte(cpu.HL))
	}

	// LD E,(HL)
	// (HL) means that the H & L registers are used as single 16-bit memory address, the
	// contents of which are stored to the E register
	opcodes[0x5e] = func(cpu *Cpu, mem *Memory) {
		cpu.setE(mem.Read(cpu.HL))
	}

	// LD H,B
	opcodes[0x60] = func(cpu *Cpu, mem *Memory) {
		cpu.setH(highByte(cpu.BC))
	}

	// LD H,C
	opcodes[0x61] = func(cpu *Cpu, mem *Memory) {
		cpu.setH(lowByte(cpu.BC))
	}

	// LD H,D
	opcodes[0x62] = func(cpu *Cpu, mem *Memory) {
		cpu.setH(highByte(cpu.DE))
	}

	// LD H,E
	opcodes[0x63] = func(cpu *Cpu, mem *Memory) {
		cpu.setH(lowByte(cpu.DE))
	}

	// LD H,H
	opcodes[0x64] = func(cpu *Cpu, mem *Memory) {
		cpu.setH(highByte(cpu.HL))
	}

	// LD H,L
	opcodes[0x65] = func(cpu *Cpu, mem *Memory) {
		cpu.setH(lowByte(cpu.HL))
	}

	// LD H,(HL)
	// (HL) means that the H & L registers are used as single 16-bit memory address, the
	// contents of which are stored to the H register
	opcodes[0x66] = func(cpu *Cpu, mem *Memory) {
		cpu.setH(mem.Read(cpu.HL))
	}

	// LD L,B
	opcodes[0x68] = func(cpu *Cpu, mem *Memory) {
		cpu.setL(highByte(cpu.BC))
	}

	// LD L,C
	opcodes[0x69] = func(cpu *Cpu, mem *Memory) {
		cpu.setL(lowByte(cpu.BC))
	}

	// LD L,D
	opcodes[0x6a] = func(cpu *Cpu, mem *Memory) {
		cpu.setL(highByte(cpu.DE))
	}

	// LD L,E
	opcodes[0x6b] = func(cpu *Cpu, mem *Memory) {
		cpu.setL(lowByte(cpu.DE))
	}

	// LD L,H
	opcodes[0x6c] = func(cpu *Cpu, mem *Memory) {
		cpu.setL(highByte(cpu.HL))
	}

	// LD L,L
	opcodes[0x6d] = func(cpu *Cpu, mem *Memory) {
		cpu.setL(lowByte(cpu.HL))
	}

	// LD L,(HL)
	// (HL) means that the H & L registers are used as single 16-bit memory address, the
	// contents of which are stored to the L register
	opcodes[0x6e] = func(cpu *Cpu, mem *Memory) {
		cpu.setL(mem.Read(cpu.HL))
	}

	// LD (HL),B
	// (HL) means that the H & L registers are used as single 16-bit memory address, which is used as
	// the target address
	opcodes[0x70] = func(cpu *Cpu, mem *Memory) {
		mem.ram[cpu.HL] = highByte(cpu.BC)
	}

	// LD (HL),C
	// (HL) means that the H & L registers are used as single 16-bit memory address, which is used as
	// the target address
	opcodes[0x71] = func(cpu *Cpu, mem *Memory) {
		mem.ram[cpu.HL] = lowByte(cpu.BC)
	}

	// LD (HL),D
	// (HL) means that the H & L registers are used as single 16-bit memory address, which is used as
	// the target address
	opcodes[0x72] = func(cpu *Cpu, mem *Memory) {
		mem.ram[cpu.HL] = highByte(cpu.DE)
	}

	// LD (HL),E
	// (HL) means that the H & L registers are used as single 16-bit memory address, which is used as
	// the target address
	opcodes[0x73] = func(cpu *Cpu, mem *Memory) {
		mem.ram[cpu.HL] = lowByte(cpu.DE)
	}

	// LD (HL),H
	// (HL) means that the H & L registers are used as single 16-bit memory address, which is used as
	// the target address
	opcodes[0x74] = func(cpu *Cpu, mem *Memory) {
		mem.ram[cpu.HL] = highByte(cpu.HL)
	}

	// LD (HL),L
	// (HL) means that the H & L registers are used as single 16-bit memory address, which is used as
	// the target address
	opcodes[0x75] = func(cpu *Cpu, mem *Memory) {
		mem.ram[cpu.HL] = lowByte(cpu.HL)
	}

	// LD (HL),n
	// (HL) means that the H & L registers are used as single 16-bit memory address, which is used as
	// the target address
	opcodes[0x36] = func(cpu *Cpu, mem *Memory) {
		cpu.PC++
		mem.ram[cpu.HL] = mem.Read(cpu.PC)
	}

}
