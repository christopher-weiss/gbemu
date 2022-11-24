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

// HighByte returns the value of the first byte of combined register value e.g. Return the A value of the AF-register.
func highByte(rVal uint16) uint8 {
	return uint8(rVal >> 8)
}

// HighByte returns the value of the first byte of combined register value e.g. Return the F value of the AF-register.
func lowByte(rVal uint16) uint8 {
	return uint8(rVal & 0xff)
}
