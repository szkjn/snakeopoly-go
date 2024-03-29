package game

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	Background:  LighterGreen,
	Grid:        LightGreen,
	DrawElement: DarkerGreen,
}

var NightTheme = ColorTheme{
	Background:  DarkerGreen,
	Grid:        DarkGreen,
	DrawElement: LighterGreen,
}

var ApocalypseTheme = ColorTheme{
	Background:  DarkerRed,
	Grid:        DarkRed,
	DrawElement: LighterRed,
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
func (ui *UI) DrawBaseElements(screen *ebiten.Image, debugMode bool) {
	screen.Fill(ui.Theme.Background)
	if debugMode {
		ui.DrawGrid(screen)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS  %0.0f\nTPS  %0.0f\n", ebiten.ActualFPS(), ebiten.ActualTPS()))
	}
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
func (ui *UI) DrawWelcomePage(screen *ebiten.Image, g *Game) {
	ui.DrawBaseElements(screen, g.DebugMode)

	ui.DrawText(screen, "center", "Welcome to the Google's Snakeopoly!", FontL, 4)
	ui.DrawText(screen, "center", "Slither your way", FontL, 6)
	ui.DrawText(screen, "center", "to Surveillance Sovereignty!", FontL, 7.5)

	// Draw the welcome animation
	ui.DrawWelcomeAnimation(screen, g, ui.Theme)

	if g.BlinkText {
		ui.DrawText(screen, "center", "Press P to play or Q to quit", FontM, 18.5)
	}
}

// Draws the Play Page
func (ui *UI) DrawPlayPage(screen *ebiten.Image, g *Game) {
	ui.DrawBaseElements(screen, g.DebugMode)

	scale, x, y := PlaceDataPoint(g.CurrentDataPoint)
	g.UI.DrawImage(screen, g.CurrentDataPoint.GetImage(), scale, x, y)

	// Draw the snake based on visibility state
	if g.SnakeVisible {
		snakeImage := g.Snake.GetImage() // Get the snake image
		for _, segment := range g.Snake.Body {
			segmentX, segmentY := segment[0]*ScreenUnit, segment[1]*ScreenUnit
			// Draw the snake segment image
			g.UI.DrawImage(screen, snakeImage, 1.0, float64(segmentX), float64(segmentY)) // Assuming scale = 1.0 for snake segments
		}
	}

	scoreDisplay := fmt.Sprintf("Score: %d", g.Score)
	ui.DrawText(screen, "left", scoreDisplay, FontM, 17)
	levelDisplay := fmt.Sprintf("Level: %s", g.Level)
	ui.DrawText(screen, "right", levelDisplay, FontM, 17)
}

// Draws the Special Page
func (ui *UI) DrawSpecialPage(screen *ebiten.Image, g *Game) {
	ui.DrawBaseElements(screen, g.DebugMode)

	name := g.CurrentSpecialDataPoint.Name
	image := g.CurrentSpecialDataPoint.Image
	textStr := g.CurrentSpecialDataPoint.Text
	maxLineWidth := int(ScreenWidth) - 11*int(ScreenUnit)

	ui.DrawText(screen, "center", "Congrats! You've just acquired:", FontL, 3.5)
	ui.DrawText(screen, "center", name, FontL, 5)

	scale, x, y := ui.PlaceImage(image, 6, 3, "center")
	ui.DrawImage(screen, image, scale, x, y)
	ui.DrawMultiLineText(screen, textStr, 7.5, 10.5, FontM, maxLineWidth, g.CurrentCharIndex)

	ui.DrawEvil(screen, float64(ScreenUnit)*2, float64(PlayAreaHeight)-float64(ScreenUnit)*5)
	ui.DrawFire(screen, float64(PlayAreaHeight)-float64(ScreenUnit)*0.7)

	scoreDisplay := fmt.Sprintf("Score: %d", g.Score)
	ui.DrawText(screen, "left", scoreDisplay, FontM, 17)
	levelDisplay := fmt.Sprintf("Level: %s", g.Level)
	ui.DrawText(screen, "right", levelDisplay, FontM, 17)

	totalLength := len(textStr)

	if g.CurrentCharIndex >= totalLength {
		if g.BlinkText {
			ui.DrawText(screen, "center", "Press R to resume or Q to quit", FontM, 18.5)
		}
	}
}

// Draws the Game Over Page
func (ui *UI) DrawGameOverPage(screen *ebiten.Image, g *Game) {
	ui.DrawBaseElements(screen, g.DebugMode)

	scoreDisplay := fmt.Sprintf("Score: %d", g.Score)
	levelDisplay := fmt.Sprintf("Level: %s", g.Level)

	ui.DrawText(screen, "center", "GAME OVER", FontXL, 4)
	ui.DrawText(screen, "center", scoreDisplay, FontM, 6)
	ui.DrawText(screen, "center", levelDisplay, FontM, 7)
	ui.DrawText(screen, "center", "Oops! You've been out-monopolized.", FontM, 10)
	ui.DrawText(screen, "center", "But don't worry, your data", FontM, 11)
	ui.DrawText(screen, "center", "will live on forever with us.", FontM, 12)

	ui.DrawFire(screen, float64(PlayAreaHeight)-float64(ScreenUnit)*0.7)

	if g.BlinkText {
		ui.DrawText(screen, "center", "Press P to play or Q to quit", FontM, 18.5)
	}
}

// Draws the Goal Page
func (ui *UI) DrawGoalPage(screen *ebiten.Image, g *Game) {
	ui.DrawBaseElements(screen, g.DebugMode)

	// scoreDisplay := fmt.Sprintf("Score: %d", score)
	// levelDisplay := fmt.Sprintf("Level: XXX")

	ui.DrawText(screen, "center", "CONGRATULATIONS !", FontXL, 4)
	ui.DrawText(screen, "center", "Master of the Digital Panopticon !", FontL, 7)
	ui.DrawText(screen, "center", "In the world of Surveillance Capitalism,", FontL, 8.5)
	ui.DrawText(screen, "center", "you stand unrivaled !", FontL, 10)
	ui.DrawText(screen, "center", "A true data supremacist !!!", FontL, 13)
	ui.DrawFire(screen, float64(PlayAreaHeight)-float64(ScreenUnit)*0.7)

	if g.BlinkText {
		ui.DrawText(screen, "center", "Press P to replay or Q to quit", FontM, 18.5)
	}
}

// Draws text aligned to the specified side (left or right)
func (ui *UI) DrawText(screen *ebiten.Image, alignment string, textStr string, fontFace font.Face, yUnits float32) {
	// Calculate the text width
	textRect, _ := font.BoundString(fontFace, textStr)
	textWidth := float32((textRect.Max.X - textRect.Min.X).Round())

	y := int(ScreenUnit*yUnits - ScreenUnit*0.1)

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
	// Add double quotes at the beginning and end of the textStr
	textStr = "\"" + textStr + "\""

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

func (ui *UI) PlaceImage(img *ebiten.Image, yUnits float32, scale float32, alignment string) (float64, float64, float64) {
	imgWidth := float32(img.Bounds().Dx())
	y := ScreenUnit*yUnits - ScreenUnit*0.1

	scaledWidth := float32(imgWidth) * scale
	var x float32
	switch alignment {
	case "center":
		x = (ScreenWidth - scaledWidth) / 2
	case "left":
		x = 0
	case "right":
		x = ScreenWidth - scaledWidth
	default:
		panic("Invalid alignment")
	}

	return float64(scale), float64(x), float64(y)
}

// Draw image depending on scale and alignment
func (ui *UI) DrawImage(screen *ebiten.Image, img *ebiten.Image, scale, x, y float64) {
	filteredImage := ApplyMonochromeFilter(img, ui.Theme.DrawElement)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	screen.DrawImage(filteredImage, op)
}

// SetScore sets the current score to be displayed in the UI
func (ui *UI) SetScore(score int8) {
	ui.score = score
}

// SetGameOver sets the game over state and triggers a game over message to be displayed
func (ui *UI) SetGameOver() {
	ui.gameOver = true
}

// Applies a monochrome filter to an image
func ApplyMonochromeFilter(img *ebiten.Image, drawElementColor color.Color) *ebiten.Image {
	// Create a new image with the same size as the original
	filteredImg := ebiten.NewImageFromImage(img)

	// Convert the drawElementColor to RGBA
	r, g, b, _ := drawElementColor.RGBA()

	// Normalize RGB values to 0-255 range
	r, g, b = r>>8, g>>8, b>>8

	// Get the size of the original image
	width, height := img.Size()

	// Iterate over each pixel
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Get the original color of the pixel
			_, _, _, a := img.At(x, y).RGBA()

			// Check if the pixel is not completely transparent
			if a != 0 {
				// Set the pixel to the DrawElement color but keep the original alpha
				filteredColor := color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(a >> 8), // Convert 16-bit alpha to 8-bit
				}
				filteredImg.Set(x, y, filteredColor)
			}
		}
	}

	return filteredImg
}

// Draw pixelated shape given a 2D array
func (ui *UI) DrawChar(screen *ebiten.Image, char [][]int, x, y, pixelSize float64) {
	for i, row := range char {
		for j, pixel := range row {
			if pixel == 1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(ShapePixelSize/pixelSize, ShapePixelSize/pixelSize)
				op.GeoM.Translate(x+float64(j)*float64(ShapePixelSize), y+float64(i)*float64(ShapePixelSize))

				img := ebiten.NewImage(int(math.Round(ShapePixelSize)), int(math.Round(ShapePixelSize)))
				img.Fill(ui.Theme.DrawElement)
				screen.DrawImage(img, op)
			}
		}
	}
}

func (ui *UI) DrawFlower(screen *ebiten.Image, y float64) {
	ui.DrawChar(screen, FlowerShape, float64(ScreenUnit), y, 8)
	ui.DrawChar(screen, FlowerShape, float64(ScreenWidth)/2, y, 8)
}

func (ui *UI) DrawCity(screen *ebiten.Image, y float64) {
	ui.DrawChar(screen, CityShape, float64(ScreenUnit), y, 8)
	ui.DrawChar(screen, CityShape, float64(ScreenWidth)/2, y, 8)
}

func (ui *UI) DrawFire(screen *ebiten.Image, y float64) {
	ui.DrawChar(screen, FireShape, float64(ScreenUnit), y, 8)
	ui.DrawChar(screen, FireShape, float64(ScreenWidth)/2, y, 8)
}

func (ui *UI) DrawEvil(screen *ebiten.Image, x, y float64) {
	ui.DrawChar(screen, EvilShape, x, y, 8)
}

func (ui *UI) DrawGEvil(screen *ebiten.Image, x, y float64) {
	ui.DrawChar(screen, GEvilShape, x, y, 8)
}

func (ui *UI) DrawAsciiArt(screen *ebiten.Image, asciiArt []string, x, y int) {
	for i, line := range asciiArt {
		text.Draw(screen, line, FontXS, x, y+i*int(FontXS.Metrics().Height.Ceil()), ui.Theme.DrawElement)
	}
}

// DrawWelcomeAnimation draws the GShape and SixShape alternately
func (ui *UI) DrawWelcomeAnimation(screen *ebiten.Image, g *Game, initialUserTheme ColorTheme) {

	// Calculate the center of the shape
	shapeWidth := float64(len(GShape)) * ShapePixelSize
	centerX := float64(ScreenWidth)/2 - float64(shapeWidth)/2

	if !g.IsGShape {
		ui.Theme = ApocalypseTheme
		// Draw the shape
		ui.DrawChar(screen, SixShape, centerX, float64(PlayAreaHeight)*0.65, 8)
		ui.DrawChar(screen, SixShape, centerX-shapeWidth-ShapePixelSize, float64(PlayAreaHeight)*0.65, 8)
		ui.DrawChar(screen, SixShape, centerX+shapeWidth+ShapePixelSize, float64(PlayAreaHeight)*0.65, 8)
		ui.DrawFire(screen, float64(PlayAreaHeight)-float64(ScreenUnit)*0.7)

	} else {
		ui.Theme = DayTheme
		ui.DrawCity(screen, float64(PlayAreaHeight)-float64(ScreenUnit)*1.6)
		ui.DrawChar(screen, GShape, centerX, float64(PlayAreaHeight)*0.65, 8)
	}
}
