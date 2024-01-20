package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Define a UI struct to manage UI elements
type UI struct {
	score    int
	gameOver bool
}

// NewUI initializes and returns a new UI instance
func NewUI() *UI {
	ui := &UI{
		score:    0,
		gameOver: false,
	}
	return ui
}

// DrawUI draws UI elements on the screen
func (ui *UI) DrawUI(screen *ebiten.Image) {
	// Draw the current score on the screen
	text.Draw(screen, "Score: ", FontMain, 10, 10, color.White)

	// If the game is over, display a game over message
	if ui.gameOver {
		text.Draw(screen, "GAME OVER", FontMain, int(ScreenWidth/2-100), NewUI().score/2, color.White)
	}
}

// drawGrid draws the grid lines on the screen
func (ui *UI) DrawGrid(screen *ebiten.Image) {
	// Vertical lines
	for x := 0; x < int(ScreenWidth); x += int(ScreenUnit) {
		vector.StrokeLine(screen, float32(x), float32(0), float32(x), float32(ScreenHeight), float32(1), LightWhite, false)
	}

	// Horizontal lines
	for y := 0; y < int(ScreenHeight); y += int(ScreenUnit) {
		vector.StrokeLine(screen, 0, float32(y), float32(ScreenWidth), float32(y), float32(1), LightWhite, false)
	}
}

// HandleInput handles user input related to UI elements
func (ui *UI) HandleInput() {
	// You can add code here to handle UI input
	// For example, checking for button clicks or keyboard input
}

// SetScore sets the current score to be displayed in the UI
func (ui *UI) SetScore(score int) {
	ui.score = score
}

// SetGameOver sets the game over state and triggers a game over message to be displayed
func (ui *UI) SetGameOver() {
	ui.gameOver = true
}

type Welcome struct {
	UI *UI // UI elements for the welcome page
}

func NewWelcome() *Welcome {
	return &Welcome{
		UI: NewUI(),
	}
}
