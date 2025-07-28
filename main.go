package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Vedjw/DonkeyType/internals/words"
	"github.com/Vedjw/DonkeyType/state"
	"github.com/Vedjw/DonkeyType/ui"
)

func init() {
	// Make mistake map
	state.MistakeMap = make(map[int]bool)

	// Populate dictionary
	for _, word := range strings.Split(words.WordsRaw, "\n") {
		word = strings.TrimSpace(word)
		if word != "" {
			words.Dictionary = append(words.Dictionary, word)
		}
	}
}

func main() {
	playAgain := true
	for playAgain {
		quit, err := ui.RenderList()
		if quit {
			playAgain = !quit
		}
		if err != nil {
			fmt.Println("Error running program RenderList:", err)
			os.Exit(1)
		}
		if !quit {
			quit, err = ui.RenderTextarea()
			if quit {
				playAgain = !quit
			}
			if err != nil {
				fmt.Println("Error running program RenderTextarea:", err)
				os.Exit(1)
			}
		}
		if !quit {
			playAgain, err = ui.RenderResults()
			if err != nil {
				fmt.Println("Error running program RenderResults:", err)
				os.Exit(1)
			}
		}
		state.Reset()
	}
}
