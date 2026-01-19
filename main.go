package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	game := &Game{
		random: rand.New(source),
	}
	game.spawnPiece()

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Go Tetris")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
