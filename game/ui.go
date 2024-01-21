package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

// Define a UI struct to manage UI elements
type UI struct {
	score    int8
	gameOver bool
}

// Initialize and return a new UI instance
func NewUI() *UI {
	return &UI{score: 0, gameOver: false}
}

// Draw base elements common to all displays
func (ui *UI) DrawBaseElements(screen *ebiten.Image) {
	screen.Fill(Black)
	ui.DrawGrid(screen)
	ui.DrawPlayArea(screen)
}

// Draw grid lines on screen
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

// Draw Welcome Page
func (ui *UI) DrawWelcomePage(screen *ebiten.Image) {
	ui.DrawBaseElements(screen)

	ui.DrawText(screen, "center", "Welcome to the Snakeopoly!", FontL, PlayAreaHeight*0.2, White)
	ui.DrawText(screen, "center", "Slither your way", FontL, PlayAreaHeight*0.35, White)
	ui.DrawText(screen, "center", "to Surveillance Sovereignty!", FontL, PlayAreaHeight*0.45, White)
	ui.DrawText(screen, "center", "Press P to play or Q to quit", FontM, ScreenHeight-ScreenUnit*2, White)
}

// Draws the Play Page
func (ui *UI) DrawPlayPage(screen *ebiten.Image, g *Game) {
	ui.DrawBaseElements(screen)

	// Draw the data point image at the data point coordinates
	op := &ebiten.DrawImageOptions{}
	x, y := g.CurrentDataPoint.Position()
	op.GeoM.Translate(float64(x*ScreenUnit), float64(y*ScreenUnit))
	DrawDataPoint(screen, g.CurrentDataPoint)

	// Draw the snake
	for _, segment := range g.Snake.Body {
		segmentX, segmentY := segment[0]*ScreenUnit, segment[1]*ScreenUnit
		vector.DrawFilledRect(screen, segmentX, segmentY, SnakeSize, SnakeSize, White, false)
	}
	scoreDisplay := fmt.Sprintf("Score: %d", g.Score)
	ui.DrawText(screen, "left", scoreDisplay, FontS, PlayAreaHeight+ScreenUnit*2, White)
	debug := fmt.Sprintf("CurrentDataPoint: %v", g.CurrentDataPoint)
	ui.DrawText(screen, "left", debug, FontS, PlayAreaHeight+ScreenUnit*3, White)

}

// Draws the Game Over Page
func (ui *UI) DrawGameOverPage(screen *ebiten.Image, score int8) {
	ui.DrawBaseElements(screen)

	scoreDisplay := fmt.Sprintf("Score: %d", score)
	levelDisplay := fmt.Sprintf("Level: XXX")

	ui.DrawText(screen, "center", "GAME OVER", FontXL, PlayAreaHeight*0.25, White)
	ui.DrawText(screen, "center", scoreDisplay, FontL, PlayAreaHeight*0.4, White)
	ui.DrawText(screen, "center", levelDisplay, FontL, PlayAreaHeight*0.5, White)
	ui.DrawText(screen, "center", "Oops! You've been out-monopolized.", FontL, PlayAreaHeight*0.65, White)
	ui.DrawText(screen, "center", "But don't worry, your data", FontL, PlayAreaHeight*0.75, White)
	ui.DrawText(screen, "center", "will live on forever with us.", FontL, PlayAreaHeight*0.85, White)
	ui.DrawText(screen, "center", "Press P to replay or Q to quit", FontM, ScreenHeight-ScreenUnit*2, White)
}

// Draws the Goal Page
func (ui *UI) DrawGoalPage(screen *ebiten.Image, score int8) {
	ui.DrawBaseElements(screen)

	// scoreDisplay := fmt.Sprintf("Score: %d", score)
	// levelDisplay := fmt.Sprintf("Level: XXX")

	ui.DrawText(screen, "center", "CONGRATULATIONS !", FontXL, PlayAreaHeight*0.25, White)
	ui.DrawText(screen, "center", "Master of the Digital Panopticon !", FontL, PlayAreaHeight*0.4, White)
	ui.DrawText(screen, "center", "In the world of Surveillance Capitalism,", FontL, PlayAreaHeight*0.5, White)
	ui.DrawText(screen, "center", "you stand unrivaled !", FontL, PlayAreaHeight*0.6, White)
	ui.DrawText(screen, "center", "A true data supremacist !!!", FontL, PlayAreaHeight*0.75, White)
	ui.DrawText(screen, "center", "Press P to replay or Q to quit", FontM, ScreenHeight-ScreenUnit*2, White)
}

// Draws text aligned to the specified side (left or right)
func (ui *UI) DrawText(screen *ebiten.Image, alignment string, txt string, fontFace font.Face, y float32, color color.Color) {
	// Calculate the text width
	textRect, _ := font.BoundString(fontFace, txt)
	textWidth := float32((textRect.Max.X - textRect.Min.X).Round())

	// Calculate the x position based on the specified alignment
	var x float32
	if alignment == "left" {
		x = ScreenUnit
	} else if alignment == "right" {
		x = ScreenWidth - textWidth
	} else {
		// Default to center if an invalid side is provided
		x = (ScreenWidth - textWidth) / 2
	}

	// Draw the aligned text
	text.Draw(screen, txt, fontFace, int(x), int(y), color)
}

// SetScore sets the current score to be displayed in the UI
func (ui *UI) SetScore(score int8) {
	ui.score = score
}

// SetGameOver sets the game over state and triggers a game over message to be displayed
func (ui *UI) SetGameOver() {
	ui.gameOver = true
}

// Helper functions
