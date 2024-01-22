package game

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

// Define a UI struct to manage UI elements
type UI struct {
	score    int8
	gameOver bool
	Theme    ColorTheme
}

// Define color themes
type ColorTheme struct {
	Background  color.Color
	Grid        color.Color
	DrawElement color.Color
}

var DayTheme = ColorTheme{
	Background:  White,
	Grid:        LightBlack,
	DrawElement: Black,
}

var NightTheme = ColorTheme{
	Background:  Black,
	Grid:        LightWhite,
	DrawElement: White,
}

// Initialize and return a new UI instance
func NewUI() *UI {
	return &UI{score: 0, gameOver: false, Theme: DayTheme}
}

// Toggles between color themes
func (ui *UI) ToggleTheme(theme ColorTheme) {
	ui.Theme = theme
}

// Draw base elements common to all displays
func (ui *UI) DrawBaseElements(screen *ebiten.Image) {
	screen.Fill(ui.Theme.Background)
	ui.DrawGrid(screen)
	ui.DrawPlayArea(screen)
}

// Draw grid lines on screen
func (ui *UI) DrawGrid(screen *ebiten.Image) {
	// Vertical lines
	for x := 0; x < int(ScreenWidth); x += int(ScreenUnit) {
		vector.StrokeLine(screen, float32(x), float32(0), float32(x), float32(ScreenHeight), float32(1), ui.Theme.Grid, false)
	}
	// Horizontal lines
	for y := 0; y < int(ScreenHeight); y += int(ScreenUnit) {
		vector.StrokeLine(screen, 0, float32(y), float32(ScreenWidth), float32(y), float32(1), ui.Theme.Grid, false)
	}
}

// Draw Play Area borders
func (ui *UI) DrawPlayArea(screen *ebiten.Image) {
	vector.StrokeRect(screen, PlayAreaX1, PlayAreaY1, PlayAreaWidth, PlayAreaHeight, 2, ui.Theme.DrawElement, false)
}

// Draw Welcome Page
func (ui *UI) DrawWelcomePage(screen *ebiten.Image, blinkText bool) {
	ui.DrawBaseElements(screen)

	ui.DrawText(screen, "center", "Welcome to the Snakeopoly!", FontL, 4)
	ui.DrawText(screen, "center", "Slither your way", FontL, 6)
	ui.DrawText(screen, "center", "to Surveillance Sovereignty!", FontL, 7.5)
	if blinkText {
		ui.DrawText(screen, "center", "Press P to play or Q to quit", FontM, 18.5)
	}
}

// Draws the Play Page
func (ui *UI) DrawPlayPage(screen *ebiten.Image, g *Game) {
	ui.DrawBaseElements(screen)

	// Draw the data point image at the data point coordinates
	op := &ebiten.DrawImageOptions{}
	x, y := g.CurrentDataPoint.Position()
	op.GeoM.Translate(float64(x*ScreenUnit), float64(y*ScreenUnit))
	DrawDataPoint(screen, g.CurrentDataPoint)

	// Draw the snake based on visibility state
	if g.SnakeVisible {
		for _, segment := range g.Snake.Body {
			segmentX, segmentY := segment[0]*ScreenUnit, segment[1]*ScreenUnit
			vector.DrawFilledRect(screen, segmentX, segmentY, SnakeSize, SnakeSize, ui.Theme.DrawElement, false)
		}
	}

	scoreDisplay := fmt.Sprintf("Score: %d", g.Score)
	ui.DrawText(screen, "left", scoreDisplay, FontM, 17)
	levelDisplay := fmt.Sprintf("Level: %s", g.Level)
	ui.DrawText(screen, "right", levelDisplay, FontM, 17)
}

// Draws the Special Page
func (ui *UI) DrawSpecialPage(screen *ebiten.Image, specialDP SpecialDataPoint, currentCharIndex int, blinkText bool) {
	ui.DrawBaseElements(screen)

	name := specialDP.Name
	image := specialDP.Image
	textStr := specialDP.Text
	maxLineWidth := int(ScreenWidth) - 10*int(ScreenUnit)

	ui.DrawText(screen, "center", "Congrats! You've just acquired:", FontL, 4)
	ui.DrawText(screen, "center", name, FontL, 5.5)

	ui.DrawImage(screen, image, 6.5, 3, "center")
	ui.DrawMultiLineText(screen, textStr, 7, 11, FontM, maxLineWidth, currentCharIndex)

	totalLength := len(textStr)

	if currentCharIndex >= totalLength {
		if blinkText {
			ui.DrawText(screen, "center", "Press R to resume or Q to quit", FontM, 18.5)
		}
	}
}

// Draws the Game Over Page
func (ui *UI) DrawGameOverPage(screen *ebiten.Image, score int8, blinkText bool) {
	ui.DrawBaseElements(screen)

	scoreDisplay := fmt.Sprintf("Score: %d", score)
	levelDisplay := fmt.Sprintf("Level: XXX")

	ui.DrawText(screen, "center", "GAME OVER", FontXL, 4)
	ui.DrawText(screen, "center", scoreDisplay, FontM, 6)
	ui.DrawText(screen, "center", levelDisplay, FontM, 7)
	ui.DrawText(screen, "center", "Oops! You've been out-monopolized.", FontM, 9)
	ui.DrawText(screen, "center", "But don't worry, your data", FontM, 10)
	ui.DrawText(screen, "center", "will live on forever with us.", FontM, 11)
	if blinkText {
		ui.DrawText(screen, "center", "Press P to replay or Q to quit", FontM, 18.5)
	}
}

// Draws the Goal Page
func (ui *UI) DrawGoalPage(screen *ebiten.Image, score int8, blinkText bool) {
	ui.DrawBaseElements(screen)

	// scoreDisplay := fmt.Sprintf("Score: %d", score)
	// levelDisplay := fmt.Sprintf("Level: XXX")

	ui.DrawText(screen, "center", "CONGRATULATIONS !", FontXL, 4)
	ui.DrawText(screen, "center", "Master of the Digital Panopticon !", FontL, 6)
	ui.DrawText(screen, "center", "In the world of Surveillance Capitalism,", FontL, 7.5)
	ui.DrawText(screen, "center", "you stand unrivaled !", FontL, 9)
	ui.DrawText(screen, "center", "A true data supremacist !!!", FontL, 11)
	if blinkText {
		ui.DrawText(screen, "center", "Press P to replay or Q to quit", FontM, 18.5)
	}
}

// Draws text aligned to the specified side (left or right)
func (ui *UI) DrawText(screen *ebiten.Image, alignment string, textStr string, fontFace font.Face, yUnits float32) {
	// Calculate the text width
	textRect, _ := font.BoundString(fontFace, textStr)
	textWidth := float32((textRect.Max.X - textRect.Min.X).Round())

	y := int(ScreenUnit*yUnits - ScreenUnit*0.1)
	// fmt.Println(textStr, y)

	// Calculate the x position based on the specified alignment
	var x float32
	if alignment == "left" {
		x = ScreenUnit
	} else if alignment == "right" {
		x = ScreenWidth - textWidth - ScreenUnit
	} else {
		// Default to center if an invalid side is provided
		x = (ScreenWidth - textWidth) / 2
	}

	// Draw the aligned text
	text.Draw(screen, textStr, fontFace, int(x), y, ui.Theme.DrawElement)
}

func (ui *UI) DrawMultiLineText(screen *ebiten.Image, textStr string, xUnits, yUnits float32, fontFace font.Face, maxLineWidth int, currentCharIndex int) {

	// Split the text into words
	words := strings.Fields(textStr)
	var lines []string
	var currentLine string
	x := int(ScreenUnit * xUnits)
	y := int(ScreenUnit*yUnits - ScreenUnit*0.1)

	for _, word := range words {
		// Check line width with the new word added
		testLine := currentLine
		if currentLine != "" {
			testLine += " " // Add a space before the word if it's not the first word in the line
		}
		testLine += word

		bounds, _ := font.BoundString(fontFace, testLine)
		lineWidth := (bounds.Max.X - bounds.Min.X).Ceil()

		if lineWidth <= maxLineWidth {
			// If it fits, add the word to the current line
			currentLine = testLine
		} else {
			// If it doesn't fit, start a new line
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = word
		}
	}

	// Add the last line
	lines = append(lines, currentLine)

	// Draw each line, up to currentCharIndex
	charsDrawn := 0
	for i, line := range lines {
		if charsDrawn+len(line) > currentCharIndex {
			line = line[:currentCharIndex-charsDrawn]
		}
		lineSpacing := i * int(ScreenUnit)
		text.Draw(screen, line, fontFace, x, y+lineSpacing, ui.Theme.DrawElement)
		charsDrawn += len(line)
		if charsDrawn >= currentCharIndex {
			break
		}
	}
}

// Draw image depending on scale and alignment
func (ui *UI) DrawImage(screen *ebiten.Image, image *ebiten.Image, yUnits float32, scale float32, alignment string) {
	// Get image dimensions, we assume every image is a square (width=height)
	imgWidth := image.Bounds().Dx()
	y := int(ScreenUnit*yUnits - ScreenUnit*0.1)

	// Calculate scaled dimensions
	scaledWidth := float32(imgWidth) * scale
	// scaledHeight := float32(imgHeight) * scale

	// Calculate position based on alignment
	var x float32
	switch alignment {
	case "center":
		x = (ScreenWidth - scaledWidth) / 2
	case "left":
		x = 0
	case "right":
		x = ScreenWidth - scaledWidth
	default:
		panic("Invalid alignment") // Handle the error appropriately
	}

	// Draw the image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(scale), float64(scale))
	op.GeoM.Translate(float64(x), float64(y)) // Assuming y position is always 0 for simplicity
	screen.DrawImage(image, op)
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
