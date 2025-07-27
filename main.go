package main

import (
	"fmt"
	"os"

	"github.com/Vedjw/DonkeyType/ui"
)

func main() {

	if err := ui.RenderList(); err != nil {
		fmt.Println("Error running program RenderList:", err)
		os.Exit(1)
	}
	if err := ui.RenderTextarea(); err != nil {
		fmt.Println("Error running program RenderFullscreen:", err)
		os.Exit(1)
	}
}
