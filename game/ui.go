package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
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
	text.Draw(screen, "Score: ", FontL, 10, 10, color.White)

	// If the game is over, display a game over message
	if ui.gameOver {
		text.Draw(screen, "GAME OVER", FontL, int(ScreenWidth/2-100), NewUI().score/2, color.White)
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

// Draw Play Area borders
func (ui *UI) DrawPlayArea(screen *ebiten.Image) {
	vector.StrokeRect(screen, PlayAreaX1, PlayAreaY1, PlayAreaWidth, PlayAreaHeight, 2, White, false)
}

// Draws the Welcome Page
func (ui *UI) DrawWelcomePage(screen *ebiten.Image) {
	ui.DrawPlayArea(screen)

	drawCenteredText(screen, "Welcome to the Snakeopoly!", FontL, int(PlayAreaHeight*0.2), White)
	drawCenteredText(screen, "Slither your way", FontL, int(PlayAreaHeight*0.35), White)
	drawCenteredText(screen, "to Surveillance Sovereignty!", FontL, int(PlayAreaHeight*0.45), White)
	drawCenteredText(screen, "Press P to play or Q to quit", FontM, int(ScreenHeight-ScreenUnit*2), White)
}

// Draws the Play Page
func (ui *UI) DrawPlayPage(screen *ebiten.Image, snakeBody [][2]float32) {
	ui.DrawGrid(screen)
	ui.DrawPlayArea(screen)

	// Draw the snake
	for _, segment := range snakeBody {
		segmentX, segmentY := segment[0]*ScreenUnit, segment[1]*ScreenUnit
		vector.DrawFilledRect(screen, segmentX, segmentY, SnakeSize, SnakeSize, White, false)
	}
}

// Draws the Game Over Page
func (ui *UI) DrawGameOverPage(screen *ebiten.Image) {
	text.Draw(screen, "GAME OVER", FontL, int(ScreenWidth/2-100), 50, White)
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

// Helper functions

// drawCenteredText draws centered text on the x-axis
func drawCenteredText(screen *ebiten.Image, txt string, fontFace font.Face, y int, color color.Color) {
	// Calculate the text width
	textRect, _ := font.BoundString(fontFace, txt)
	textWidth := int((textRect.Max.X - textRect.Min.X).Round())
	x := (int(ScreenWidth) - int(textWidth)) / 2

	// Draw the centered text
	text.Draw(screen, txt, fontFace, x, y, color)
}
