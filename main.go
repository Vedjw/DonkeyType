package main

import (
	"fmt"
	"os"

	"github.com/Vedjw/DonkeyType/ui"
)

func main() {
	exit, err := ui.RenderList()
	if err != nil {
		fmt.Println("Error running program RenderList:", err)
		os.Exit(1)
	}
	if !exit {
		if err := ui.RenderTextarea(); err != nil {
			fmt.Println("Error running program RenderFullscreen:", err)
			os.Exit(1)
		}
	}
}
