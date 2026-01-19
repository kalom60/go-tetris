package main

import "image/color"

type Point struct {
	X, Y int
}

var Tetrominoes = [][]Point{
	{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},   // I
	{{X: -1, Y: -1}, {X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}}, // J
	{{X: 1, Y: -1}, {X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}},  // L
	{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},    // O
	{{X: 0, Y: -1}, {X: 1, Y: -1}, {X: -1, Y: 0}, {X: 0, Y: 0}}, // S
	{{X: 0, Y: -1}, {X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}},  // T
	{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 0, Y: 0}, {X: 1, Y: 0}}, // Z
}

type Piece struct {
	Type  int
	Pos   Point
	Shape []Point
}

func getShapeColor(t int) color.RGBA {
	colors := []color.RGBA{
		{0, 255, 255, 255}, // I - Cyan
		{0, 0, 255, 255},   // J - Blue
		{255, 165, 0, 255}, // L - Orange
		{255, 255, 0, 255}, // O - Yellow
		{0, 255, 0, 255},   // S - Green
		{128, 0, 128, 255}, // T - Purple
		{255, 0, 0, 255},   // Z - Red
	}
	return colors[t]
}
