package tetris

import (
	"math/rand"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

type Color string

var pieceColors = []Color{
	ColorCyan,
	ColorYellow,
	ColorGreen,
	ColorRed,
	ColorBlue,
	ColorPurple,
	ColorWhite,
}

var (
	pieces = map[int]Piece{
		0: {
			id: 0,
			// |
			// |
			// |
			// |
			points: []*Point{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}, {x: 3, y: 0}},
			color:  pieceColors[0],
		},
		1: {
			id: 1,
			// ----
			points: []*Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 0, y: 2}, {x: 0, y: 3}},
			color:  pieceColors[1],
		},
		2: {
			id: 2,
			// --
			// --
			points: []*Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 1, y: 0}, {x: 1, y: 1}},
			color:  pieceColors[2],
		},
		3: {
			id: 3,
			//  --
			// --
			points: []*Point{{x: 1, y: 0}, {x: 1, y: 1}, {x: 0, y: 1}, {x: 0, y: 2}},
			color:  pieceColors[3],
		},
		4: {
			id: 4,
			//  -
			// ---
			points: []*Point{{x: 1, y: 0}, {x: 1, y: 1}, {x: 1, y: 2}, {x: 0, y: 1}},
			color:  pieceColors[4],
		},
		5: {
			id: 5,
			// -
			// ---
			points: []*Point{{x: 1, y: 0}, {x: 1, y: 1}, {x: 1, y: 2}, {x: 0, y: 0}},
			color:  pieceColors[5],
		},
		6: {
			id: 6,
			//   -
			// ---
			points: []*Point{{x: 1, y: 0}, {x: 1, y: 1}, {x: 1, y: 2}, {x: 0, y: 2}},
			color:  pieceColors[6],
		},
		7: {
			id: 7,
			// --
			//  --
			points: []*Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 1, y: 1}, {x: 1, y: 2}},
			color:  pieceColors[5],
		},
		8: {
			id: 8,
			// |
			// | |
			//   |
			points: []*Point{{x: 0, y: 0}, {x: 1, y: 0}, {x: 1, y: 1}, {x: 2, y: 1}},
			color:  pieceColors[1],
		},
		9: {
			id: 9,
			// |
			// | |
			// |
			points: []*Point{{x: 0, y: 0}, {x: 1, y: 0}, {x: 1, y: 1}, {x: 2, y: 0}},
			color:  pieceColors[2],
		},
		10: {
			id: 10,
			// ---
			//  -
			points: []*Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 0, y: 2}, {x: 1, y: 1}},
			color:  pieceColors[3],
		},
		11: {
			id: 11,
			//  |
			// ||
			//  |
			points: []*Point{{x: 0, y: 1}, {x: 1, y: 0}, {x: 1, y: 1}, {x: 2, y: 1}},
			color:  pieceColors[4],
		},
		12: {
			id: 12,
			// ||
			// |
			// |
			points: []*Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 1, y: 0}, {x: 2, y: 0}},
			color:  pieceColors[0],
		},
		13: {
			id: 13,
			// ---
			//   _
			points: []*Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 0, y: 2}, {x: 1, y: 2}},
			color:  pieceColors[6],
		},
		14: {
			id: 14,
			//  |
			//  |
			// ||
			points: []*Point{{x: 0, y: 1}, {x: 1, y: 1}, {x: 2, y: 0}, {x: 2, y: 1}},
			color:  pieceColors[1],
		},
		15: {
			id: 15,
			// |
			// |
			// ||
			points: []*Point{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}, {x: 2, y: 1}},
			color:  pieceColors[2],
		},
		16: {
			id: 16,
			// ---
			// -
			points: []*Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 0, y: 2}, {x: 1, y: 0}},
			color:  pieceColors[3],
		},

		17: {
			id: 17,
			// ||
			//  |
			//  |
			points: []*Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 1, y: 1}, {x: 2, y: 1}},
			color:  pieceColors[4],
		},
		18: {
			id: 18,
			//  |
			// ||
			// |
			points: []*Point{{x: 0, y: 1}, {x: 1, y: 0}, {x: 1, y: 1}, {x: 2, y: 0}},
			color:  pieceColors[0],
		},
	}

	initialPieces = []int{1, 2, 3, 4, 5, 6, 7}

	ratations = map[int]int{
		0:  1,
		1:  0,
		3:  8,
		8:  3,
		4:  9,
		9:  10,
		10: 11,
		11: 4,
		5:  12,
		12: 13,
		13: 14,
		14: 5,
		6:  15,
		15: 16,
		16: 17,
		17: 6,
		7:  18,
		18: 7,
	}
)

// PickPiece returns a random piece for the game.
func PickPiece(w int) *Piece {
	rnd := rand.Intn(len(initialPieces))
	picked := pieces[rnd]

	p := picked.cp()
	p.moveTo(0, w/2-1) // center the piece in the board.
	return p
}

// Piece is a game Piece that can be moved across the board until it is
// emprinted which then becomes part of the board.
type Piece struct {
	id     int
	points []*Point
	moves  int
	color  Color
}

func (p *Piece) Moves() int {
	return p.moves
}

// CanMoveDown returns true if the piece can be moved down based on the current
// configuration of the board.
func (p *Piece) CanMoveDown(b Board) bool {
	for _, point := range p.points {
		if !point.canMoveDown(b) {
			return false
		}
	}

	return true
}

func (p *Piece) TryMoveRight(b Board) bool {
	if p.canMoveRight(b) {
		p.moveRight()
		return true
	}

	return false
}

func (p *Piece) TryMoveLeft(b Board) bool {
	if p.canMoveLeft(b) {
		p.moveLeft()
		return true
	}

	return false
}

func (p *Piece) TryMoveDown(b Board) bool {
	if p.CanMoveDown(b) {
		p.moveDown()
		return true
	}

	return false
}

// Rotate rotates the piece based on the current configuration of the board.
func (p *Piece) Rotate(w int) *Piece {
	if r, ok := ratations[p.id]; ok {
		rr := pieces[r]
		rotated := rr.cp()
		rotated.moveTo(p.points[0].x, p.points[0].y)
		rotated.bounds(w)
		return rotated
	}

	return p
}

// canMoveRight returns true if the piece can be moved right based on the
// curretn configuration of the board.
func (p *Piece) canMoveRight(b Board) bool {
	for _, point := range p.points {
		if !point.canMoveRight(b) {
			return false
		}
	}

	return true
}

// canMoveLeft returns true if the piece can be moved left based on the
// current configuration of the board.
func (p *Piece) canMoveLeft(b Board) bool {
	for _, point := range p.points {
		if !point.canMoveLeft(b) {
			return false
		}
	}

	return true
}

func (p *Piece) Color() string {
	return string(p.color)
}

// moveDown moves the piece down.
func (p *Piece) moveDown() {
	p.moves++
	for _, point := range p.points {
		point.x++
	}
}

// moveRight moves the piece right.
func (p *Piece) moveRight() {
	for _, point := range p.points {
		point.y++
	}
}

// moveLeft moves the piece left.
func (p *Piece) moveLeft() {
	for _, point := range p.points {
		point.y--
	}
}

// IsIn return true if the given point is part of the piece.
func (p *Piece) IsIn(point Point) bool {
	for _, pp := range p.points {
		if pp.eq(point) {
			return true
		}
	}
	return false
}

func (p *Piece) cp() *Piece {
	cp := &Piece{
		id:    p.id,
		color: p.color,
	}
	for _, pp := range p.points {
		cp.points = append(cp.points, &Point{x: pp.x, y: pp.y})
	}

	return cp
}

func (p *Piece) moveTo(x, y int) {
	for _, pp := range p.points {
		pp.x += x
		pp.y += y
	}
}

// moves the piece to the left or right if it is out of bounds.
// After rotating the piece, it may be out of bounds, this function fixes that.
func (p *Piece) bounds(w int) {
	oob := true
	for oob {
		moves := false
		for _, pp := range p.points {
			if pp.y < 0 {
				p.moveRight()
				moves = true
				break
			}
			if pp.y >= w {
				p.moveLeft()
				moves = true
				break
			}
		}
		if !moves {
			oob = false
		}
	}
}
