package main

import (
	"image/color"

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

type Game struct {
	grid        [BoardHeight][BoardWidth]int
	activePiece *Piece
	timer       int
}

func (g *Game) Update() error {
	if g.activePiece == nil {
		g.activePiece = &Piece{
			Type:  0,
			Pos:   Point{X: BoardWidth / 2, Y: 0},
			Shape: Tetrominoes[0],
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.activePiece.Pos.X--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.activePiece.Pos.X++
	}

	g.timer++
	if g.timer >= 30 {
		g.activePiece.Pos.Y++
		g.timer = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for row := range BoardHeight {
		for col := range BoardWidth {
			x := float32(col * BlockSize)
			y := float32(row * BlockSize)

			vector.StrokeRect(screen, x, y, BlockSize, BlockSize, 1, color.RGBA{40, 40, 40, 255}, false)

			if g.activePiece != nil {
				for _, p := range g.activePiece.Shape {
					drawX := (g.activePiece.Pos.X + p.X) * BlockSize
					drawY := (g.activePiece.Pos.Y + p.Y) * BlockSize

					vector.FillRect(screen, float32(drawX), float32(drawY),
						BlockSize, BlockSize, color.RGBA{0, 255, 255, 255}, false)
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
