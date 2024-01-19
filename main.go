package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	borderThickness    = 2
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
	// Note: The play area's border is drawn as four lines (top, bottom, left, right)
	ebitenutil.DrawLine(screen, float64(playAreaStartX), float64(playAreaStartY), float64(playAreaStartX+playAreaWidth), float64(playAreaStartY), WHITE)
	ebitenutil.DrawLine(screen, float64(playAreaStartX), float64(playAreaStartY+playAreaHeight), float64(playAreaStartX+playAreaWidth), float64(playAreaStartY+playAreaHeight), WHITE)
	ebitenutil.DrawLine(screen, float64(playAreaStartX), float64(playAreaStartY), float64(playAreaStartX), float64(playAreaStartY+playAreaHeight), WHITE)
	ebitenutil.DrawLine(screen, float64(playAreaStartX+playAreaWidth), float64(playAreaStartY), float64(playAreaStartX+playAreaWidth), float64(playAreaStartY+playAreaHeight), WHITE)

	// Draw the snake
	for _, segment := range g.snake.Body {
		segmentX, segmentY := segment[0]*screenUnit, segment[1]*screenUnit
		ebitenutil.DrawRect(screen, float64(segmentX), float64(segmentY), float64(snakeSize), float64(snakeSize), WHITE)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
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
