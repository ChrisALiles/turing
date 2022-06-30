package turing

var Tape [1000000]uint8
var TapeIndex uint64

func init() {
	Tape[0] = Two // Mark (past) left most cell
	TapeIndex = 1 // Position on the tape

}
