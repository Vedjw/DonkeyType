package words

import (
	_ "embed"
	"math/rand"
	"strings"

	"github.com/Vedjw/DonkeyType/state"
)

//go:embed 10k.txt
var wordsRaw string

var dictionary []string

func init() {
	for _, word := range strings.Split(wordsRaw, "\n") {
		word = strings.TrimSpace(word)
		if word != "" {
			dictionary = append(dictionary, word)
		}
	}
}

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
		ri := rand.Intn(10000)
		builder.WriteString(dictionary[ri])
		builder.WriteString(" ")
	}

	displayString := builder.String()

	return displayString
}
