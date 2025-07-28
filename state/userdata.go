package state

import "time"

type Length int

const (
	Short Length = iota
	Medium
	Long
	Thicc
)

// SelectedLength stores the user's choice
var SelectedLength Length

// Total words length in char
var TotalChLength int

// Total user input in char
var TotalUserInputLenght int

// Mistakes
var TotalMistakes int

// User time taken
var TimeTaken time.Duration

// MistakeMap tracks which character indices in the user output were incorrect.
// Key: character index in user input
// Value: true if the character was already counted as a mistake
var MistakeMap map[int]bool

// Resests all values in when playing again
func Reset() {
	TotalChLength = 0
	TotalUserInputLenght = 0
	TotalMistakes = 0
	TimeTaken = 0
	MistakeMap = make(map[int]bool)
}
