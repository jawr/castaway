package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/jawr/castaway/internal/world"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %w\n", err)
		os.Exit(1)
	}
}

func run() error {
	w := world.NewWorld(
		640,
		480,
		// add systems
	)

	if err := ebiten.RunGame(w); err != nil {
		return err
	}

	return nil
}
