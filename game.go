package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	BoardWidth   = 10
	BoardHeight  = 20
	BlockSize    = 30
	ScreenWidth  = (BoardWidth * BlockSize) + 200
	ScreenHeight = BoardHeight * BlockSize
)

type Game struct {
	grid        [BoardHeight][BoardWidth]int
	activePiece *Piece
	nextPiece   *Piece
	timer       int
	moveTimer   int
	score       int
	random      *rand.Rand
}

func (g *Game) Update() error {
	if g.activePiece == nil {
		g.spawnPiece()
	}

	g.moveTimer++
	if g.moveTimer >= 7 {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			newPos := Point{X: g.activePiece.Pos.X - 1, Y: g.activePiece.Pos.Y}
			if g.isValidMove(newPos, g.activePiece.Shape) {
				g.activePiece.Pos.X--
				g.moveTimer = 0
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			newPos := Point{X: g.activePiece.Pos.X + 1, Y: g.activePiece.Pos.Y}
			if g.isValidMove(newPos, g.activePiece.Shape) {
				g.activePiece.Pos.X++
				g.moveTimer = 0
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			newPos := Point{X: g.activePiece.Pos.X, Y: g.activePiece.Pos.Y + 1}
			if g.isValidMove(newPos, g.activePiece.Shape) {
				g.activePiece.Pos.Y++
				g.moveTimer = 0
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.rotatePiece()
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
		}
	}

	if g.activePiece != nil {
		c := getShapeColor(g.activePiece.Type)
		for _, p := range g.activePiece.Shape {
			drawX := float32(g.activePiece.Pos.X+p.X) * BlockSize
			drawY := float32(g.activePiece.Pos.Y+p.Y) * BlockSize

			if drawY >= 0 {
				vector.FillRect(screen, drawX, drawY, BlockSize, BlockSize, c, false)
			}
		}
	}

	sidebarX := float32(BoardWidth*BlockSize) + 20

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("SCORE: %d", g.score), int(sidebarX), 20)
	ebitenutil.DebugPrintAt(screen, "NEXT PIECE:", int(sidebarX), 60)

	if g.nextPiece != nil {
		nextCol := getShapeColor(g.nextPiece.Type)
		for _, p := range g.nextPiece.Shape {
			// Offset the preview so it sits nicely in the sidebar
			pX := sidebarX + float32(p.X+1)*BlockSize
			pY := 120 + float32(p.Y+1)*BlockSize

			vector.FillRect(screen, pX, pY, BlockSize, BlockSize, nextCol, false)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) spawnPiece() {
	if g.nextPiece == nil {
		typeIdx := g.random.Intn(len(Tetrominoes))
		g.nextPiece = &Piece{
			Type:  typeIdx,
			Pos:   Point{X: BoardWidth / 2, Y: 0},
			Shape: Tetrominoes[typeIdx],
		}
	}

	g.activePiece = g.nextPiece

	nextTypeIdx := g.random.Intn(len(Tetrominoes))
	g.nextPiece = &Piece{
		Type:  nextTypeIdx,
		Pos:   Point{X: BoardWidth / 2, Y: 0},
		Shape: Tetrominoes[nextTypeIdx],
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
	linesCleared := 0

	for y := BoardHeight - 1; y >= 0; y-- {
		full := true
		for x := range BoardWidth {
			if g.grid[y][x] == 0 {
				full = false
				break
			}
		}

		if full {
			linesCleared++
			for row := y; row > 0; row-- {
				g.grid[row] = g.grid[row-1]
			}
			g.grid[0] = [BoardWidth]int{}

			y++
		}
	}

	switch linesCleared {
	case 1:
		g.score += 100
	case 2:
		g.score += 300
	case 3:
		g.score += 500
	case 4:
		g.score += 800
	}
}
