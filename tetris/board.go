// Package tetris provides a simple implementation of the Tetris game.
package tetris

// Board represent the current state of the game.
//
// The current moving piece is overlay on top of the Board until it is emprinted
// on the Board.
type Board struct {
	m           [][]int
	w, h        int
	pieceColors [][]string
}

// initBoard creates an empty board.
func NewBoard(w, h int) *Board {
	b := &Board{
		m:           make([][]int, 0),
		w:           w,
		h:           h,
		pieceColors: make([][]string, 0),
	}

	for range h {
		var row []int
		var rowColors []string
		for range w {
			row = append(row, 0)
			rowColors = append(rowColors, "")
		}
		b.m = append(b.m, row)
		b.pieceColors = append(b.pieceColors, rowColors)
	}

	return b
}

func (b *Board) At(i, j int) int {
	return b.m[i][j]
}

// Emprint writes 1' in the board as the points indicate.
func (b *Board) Emprint(piece Piece) (int, bool) {
	if piece.CanMoveDown(*b) {
		return 0, false
	}

	for _, p := range piece.points {
		b.m[p.x][p.y] = 1
		b.pieceColors[p.x][p.y] = piece.Color()
	}

	rmCnt := b.removeFillRows()
	return rmCnt, true
}

func (b *Board) PieceColors() [][]string {
	return b.pieceColors
}

func (b *Board) removeFillRows() int {
	count := 0
	for i := 0; i < len(b.m); i++ {
		var sum int
		for _, j := range b.m[i] {
			sum += j
		}
		if sum == len(b.m[i]) { // the entire row is filled.
			b.m = append(b.m[:i], b.m[i+1:]...)
			i--
			count++
		}
	}

	for range count {
		b.m = append([][]int{make([]int, b.w)}, b.m...)
	}

	return count
}
