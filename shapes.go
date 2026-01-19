package main

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
