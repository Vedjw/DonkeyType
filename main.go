package main

import (
	"fmt"
	"os"

	"github.com/Vedjw/DonkeyType/state"
	"github.com/Vedjw/DonkeyType/ui"
)

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
				fmt.Println("Error running program RenderFullscreen:", err)
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
