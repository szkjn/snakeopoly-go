package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth        = 500
	screenHeight       = 400
	screenUnit         = screenWidth / 20 // Adjust this divisor as needed
	playAreaWidth      = screenUnit * 16  // Example: 16 units wide
	playAreaHeight     = screenUnit * 12  // Example: 12 units tall
	playAreaStartX     = (screenWidth - playAreaWidth) / 2
	playAreaStartY     = (screenHeight - playAreaHeight) / 2
	snakeSize          = screenUnit // Assuming snake size is one unit
	initialSnakeLength = 3
)

var (
	WHITE = color.RGBA{160, 210, 160, 255}
	BLACK = color.RGBA{20, 40, 20, 255}
)

type Snake struct {
	Body [][2]int // Array of x, y pairs
}

type Game struct {
	snake Snake
}

func NewSnake() Snake {
	body := make([][2]int, initialSnakeLength)
	startX := playAreaStartX/screenUnit + (playAreaWidth/screenUnit-initialSnakeLength)/2
	startY := playAreaStartY/screenUnit + playAreaHeight/(2*screenUnit)

	// Initialize the snake's body towards the left and above the center
	for i := 0; i < initialSnakeLength; i++ {
		body[i] = [2]int{startX - i, startY}
	}

	return Snake{Body: body}
}

func (g *Game) Update() error {
	// Game update logic will go here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with black color
	screen.Fill(BLACK)

	// Draw the border of the play area
	vector.StrokeRect(screen, float32(playAreaStartX), float32(playAreaStartY), float32(playAreaWidth), float32(playAreaHeight), float32(2), WHITE, false)

	// drawGrid(screen)

	// Draw the snake
	for _, segment := range g.snake.Body {
		segmentX, segmentY := segment[0]*screenUnit, segment[1]*screenUnit
		vector.DrawFilledRect(screen, float32(segmentX), float32(segmentY), float32(snakeSize), float32(snakeSize), WHITE, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func drawGrid(screen *ebiten.Image) {
	// Set the color for the grid lines
	gridColor := color.RGBA{255, 255, 255, 255} // White color

	// Vertical lines
	for x := 0; x < screenWidth; x += screenUnit {
		vector.StrokeLine(screen, float32(x), float32(0), float32(x), float32(screenHeight), float32(1), gridColor, false)
	}

	// Horizontal lines
	for y := 0; y < screenHeight; y += screenUnit {
		vector.StrokeLine(screen, 0, float32(y), float32(screenWidth), float32(y), float32(1), gridColor, false)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")

	game := &Game{
		snake: NewSnake(), // Initialize the snake
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
