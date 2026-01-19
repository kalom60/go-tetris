package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 300
	ScreenHeight = 600
	BoardWidth   = 10
	BoardHeight  = 20
	BlockSize    = 30
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for row := range BoardHeight {
		for col := range BoardWidth {
			x := float32(col * BlockSize)
			y := float32(row * BlockSize)

			vector.StrokeRect(screen, x, y, BlockSize, BlockSize, 1, color.RGBA{40, 40, 40, 255}, false)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Go Tetris")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
