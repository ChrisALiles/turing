// Turing is an exploration of Turing machnines.
// The main reference is the Stanford Encyclopedia of Philosophy (SEP) article "Turing Machines".
package turing

const (
	// Values for tape content.
	// 0 is used as blank because it's the zero value.
	// Numbers are represented on the tape in unary.
	// One "Turing one" = 0
	// Two "Turing ones" = 1
	Blank = 0
	Zero  = 1
	One   = 2
	Two   = 3
	Three = 4
	Four  = 5
	Five  = 6
	// Values for direction of tape movement.
	Left  = 'L'
	Right = 'R'
	Stay  = 'S'
	Halt  = 'H'
)

type Transition []transEntry

type transEntry struct {
	CurState    uint8
	NextState   uint8
	CurContent  uint8
	NextContent uint8
	Direction   byte
}
