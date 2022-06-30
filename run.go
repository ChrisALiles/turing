package turing

import (
	"fmt"
	"os"
)

var currentState uint8

var transLU [120]transEntry
var transLUIndex int

func Run(trans Transition) {
	// set up transition lookup "sparse matrix".
	for _, entry := range trans {
		transLU[entry.CurState*10+entry.CurContent] = entry
	}
runLoop:
	for {
		transLUIndex = searchTrans()
		currentState = transLU[transLUIndex].NextState
		Tape[TapeIndex] = transLU[transLUIndex].NextContent
		switch transLU[transLUIndex].Direction {
		case Halt:
			break runLoop
		case Left:
			TapeIndex -= 1
		case Right:
			TapeIndex += 1
		}
	}

}
func searchTrans() int {
	var idx int = int(currentState)*10 + int(Tape[TapeIndex])
	if transLU[idx].Direction != 0 {
		return idx
	}
	fmt.Println("State and tape content not found in transition table")
	fmt.Println("State:", currentState)
	fmt.Println("Tape content:", Tape[TapeIndex])
	fmt.Println("Tape position;", TapeIndex)
	os.Exit(1)

	return 0
}
