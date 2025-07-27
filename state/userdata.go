package state

type Length int

const (
	Short Length = iota
	Medium
	Long
	Thicc
)

// SelectedLength stores the user's choice globally
var SelectedLength Length

// Total words length
var TotalChLength int

// Mistakes
var TotalMistakes int
