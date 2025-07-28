package words

import (
	_ "embed"
	"math/rand"
	"strings"

	"github.com/Vedjw/DonkeyType/state"
)

//go:embed 5k.txt
var WordsRaw string

var Dictionary []string

func WordsSelector(lengthChoice state.Length) string {
	var wordLength int

	switch lengthChoice {
	case 0:
		wordLength = 25
	case 1:
		wordLength = 50
	case 2:
		wordLength = 80
	case 3:
		wordLength = 150
	default:
		return ""
	}

	var builder strings.Builder

	for i := 0; i < wordLength; i++ {
		ri := rand.Intn(5000)
		builder.WriteString(Dictionary[ri])
		builder.WriteString(" ")
	}

	displayString := builder.String()

	return displayString
}
