package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Go Tetris")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
