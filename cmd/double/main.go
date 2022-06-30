package main

import (
	"fmt"

	"github.com/ChrisALiles/turing"
)

// This is the transition table for doubling a number.
var trans = turing.Transition{
	{0, 1, turing.One, turing.Three, turing.Right},  // mark left of number, go to next cell
	{1, 2, turing.Blank, turing.Blank, turing.Left}, // next cell blank, must be zero
	{2, 2, turing.Three, turing.One, turing.Halt},   // restore the number and stop
	{1, 3, turing.One, turing.One, turing.Right},    // next cell one, go to next
	{3, 3, turing.One, turing.One, turing.Right},    // next cell one, go to next
	{3, 4, turing.Blank, turing.One, turing.Left},   // past end of number, start result
	{4, 5, turing.One, turing.Four, turing.Left},    // mark right of number, then back to start
	{5, 5, turing.One, turing.One, turing.Left},     // continue back to start
	{5, 6, turing.Three, turing.One, turing.Right},  // left of number, now continue through it
	{6, 7, turing.One, turing.Three, turing.Right},  // mark next cell of number, continue through
	{6, 10, turing.Four, turing.One, turing.Left},   // the end of the number
	{7, 7, turing.One, turing.One, turing.Right},    // continue through the number
	{7, 7, turing.Four, turing.Four, turing.Right},  // continue past end of number
	{7, 8, turing.Blank, turing.One, turing.Left},   // next cell of result
	{8, 8, turing.Four, turing.Four, turing.Left},   // back to current cell of number
	{8, 8, turing.One, turing.One, turing.Left},     // back to current cell of number
	{8, 9, turing.Three, turing.One, turing.Right},  // continue through number
	{9, 7, turing.One, turing.Three, turing.Right},  // continue through number
	{9, 10, turing.Four, turing.One, turing.Left},   // the end of the number
	{10, 10, turing.One, turing.One, turing.Left},   // get to start of number
	{10, 10, turing.Three, turing.One, turing.Left}, // get to start of number
	{10, 11, turing.Two, turing.Two, turing.Right},  // reached (past) the left most cell
	{11, 11, turing.One, turing.One, turing.Halt},   // position at start of result
}

func main() {

	// Populate the number on the tape.
	for i := 1; i < 50002; i++ {
		turing.Tape[i] = turing.One
	}

	// Run the machine.
	turing.Run(trans)

	var result uint64
	// Fetch the result.
	for i := turing.TapeIndex; turing.Tape[i] == turing.One; i++ {
		result += 1
	}
	fmt.Println("Result: ", result-1)

}
