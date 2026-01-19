package main

import (
	"image/color"
	"math/rand"

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
		g.spawnPiece()
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.rotatePiece()
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		newPos := Point{X: g.activePiece.Pos.X - 1, Y: g.activePiece.Pos.Y}
		if g.isValidMove(newPos, g.activePiece.Shape) {
			g.activePiece.Pos.X--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		newPos := Point{X: g.activePiece.Pos.X + 1, Y: g.activePiece.Pos.Y}
		if g.isValidMove(newPos, g.activePiece.Shape) {
			g.activePiece.Pos.X++
		}
	}

	g.timer++
	if g.timer >= 30 {
		newPos := Point{X: g.activePiece.Pos.X, Y: g.activePiece.Pos.Y + 1}
		if g.isValidMove(newPos, g.activePiece.Shape) {
			g.activePiece.Pos.Y++
		} else {
			g.lockPiece()
		}
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

			if g.grid[row][col] != 0 {
				c := getShapeColor(g.grid[row][col] - 1)
				vector.FillRect(screen, x, y, BlockSize, BlockSize, c, false)
			}

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

func (g *Game) spawnPiece() {
	typeIdx := rand.Intn(len(Tetrominoes))

	g.activePiece = &Piece{
		Type:  typeIdx,
		Pos:   Point{X: BoardWidth / 2, Y: 0},
		Shape: Tetrominoes[typeIdx],
	}
}

func (g *Game) lockPiece() {
	for _, p := range g.activePiece.Shape {
		x := g.activePiece.Pos.X + p.X
		y := g.activePiece.Pos.Y + p.Y

		if y < 0 {
			g.grid = [BoardHeight][BoardWidth]int{}
			return
		}

		if y >= 0 {
			g.grid[y][x] = g.activePiece.Type + 1
		}
	}

	g.checkLines()
	g.spawnPiece()
}

func (g *Game) isValidMove(pos Point, shape []Point) bool {
	for _, p := range shape {
		newX := pos.X + p.X
		newY := pos.Y + p.Y

		if newX < 0 || newX >= BoardWidth {
			return false
		}

		if newY >= BoardHeight {
			return false
		}

		if newY >= 0 && g.grid[newY][newX] != 0 {
			return false
		}
	}

	return true
}

func (g *Game) rotatePiece() {
	newShape := make([]Point, len(g.activePiece.Shape))

	for i, p := range g.activePiece.Shape {
		newShape[i] = Point{X: -p.Y, Y: p.X}
	}

	if g.isValidMove(g.activePiece.Pos, newShape) {
		g.activePiece.Shape = newShape
	}
}

func (g *Game) checkLines() {
	for y := BoardHeight - 1; y >= 0; y-- {
		full := true
		for x := range BoardWidth {
			if g.grid[y][x] == 0 {
				full = false
				break
			}
		}

		if full {
			for row := y; row > 0; row-- {
				g.grid[row] = g.grid[row-1]
			}
			g.grid[0] = [BoardWidth]int{}

			y++
		}
	}
}
