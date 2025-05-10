package main

import (
	"fmt"
	"os"
	"time"

	"com.github.anicolaspp/tetris/tetris"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	w = 10
	h = 20

	blockChar = "0"

	menu = "\np - pause, q - quit, space - drop, r - reset\n"
)

var (
	currentPiece = tetris.PickPiece(w)

	score = 0
	level = 1
	speed = tetris.Speed(1)

	paused = false
)

func main() {
	fmt.Println("Hello Tetris")

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// timeTick is a message sent every 1 second.
type timeTick struct{}

// run generates a timeTick command.
func (t timeTick) run() tea.Cmd {
	return tea.Tick(speed, func(t time.Time) tea.Msg {
		return timeTick{}
	})
}

type model struct {
	board *tetris.Board

	ticker timeTick
}

func initialModel() model {
	m := model{board: tetris.NewBoard(w, h)}

	return m
}

func (m model) Init() tea.Cmd {
	// Execute the first time tick command.
	return m.ticker.run()
}

// View generates a string representing the current state of the board with the
// current piece overlay on top.
func (m model) View() string {

	// menuStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("green"))

	var board string
	board += lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(34)).Render(menu)

	board += lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(32)).Render(fmt.Sprintf("\nScore: %d, level: %d", score, level))

	if paused {
		board += lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(20)).Blink(true).Render("\nGAME PAUSED\n")
	}

	// borderStyle := lipgloss.NewStyle().BorderStyle(lipgloss.BlockBorder()).Foreground(lipgloss.Color("white")).Background(lipgloss.Color("black"))
	top := "\n" // borderStyle.Render("\n")
	for range w + 2 {
		top += "\\"
	}
	board += top + "\n"

	for i := range h {
		var row string
		for j := range w {
			if currentPiece.IsIn(tetris.MakePoint(i, j)) {
				row += currentPiece.Color() + blockChar + tetris.ColorReset
			} else if m.board.At(i, j) == 1 {
				pieceColor := m.board.PieceColors()[i][j]
				row += pieceColor + blockChar + tetris.ColorReset
			} else {
				row += " "
			}
		}
		board += "|" + row + "|" + "\n"
	}

	bottom := ""
	for range w + 2 {
		bottom += "\\"
	}

	board += bottom
	return board
}

// Update updates the model as a response to a IO change.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case " ":
			// Drop the piece.
			if paused {
				return m, nil
			}

			for currentPiece.TryMoveDown(*m.board) {
			}

		case "p":
			// Pause the game.
			paused = !paused
			if !paused {
				return m, m.ticker.run()
			}

		case "r":
			// Reset the game.
			paused = false
			m.board = tetris.NewBoard(w, h)
			currentPiece = tetris.PickPiece(w)
			score = 0
			level = 1
			speed = tetris.Speed(1)

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if paused {
				return m, nil
			}

			currentPiece = currentPiece.Rotate(w)

		case "left":
			if paused {
				return m, nil
			}

			currentPiece.TryMoveLeft(*m.board)

		case "right":
			if paused {
				return m, nil
			}

			currentPiece.TryMoveRight(*m.board)

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if paused {
				return m, nil
			}

			currentPiece.TryMoveDown(*m.board)
		}

	case timeTick:
		if paused {
			return m, nil
		}

		if m.moveDown() != nil {
			return m, tea.Quit
		}

		return m, m.ticker.run()
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) moveDown() tea.Cmd {
	if !currentPiece.TryMoveDown(*m.board) {
		if currentPiece.Moves() == 0 {
			return tea.Quit
		}
	}

	if cnt, ok := m.board.Emprint(*currentPiece); ok {
		currentPiece = tetris.PickPiece(w)

		deletedScore := (10 * cnt) * (cnt + 1) / 2
		score += deletedScore
		level = (score / 100) + 1
	}

	speed = tetris.Speed(level)

	return nil
}
