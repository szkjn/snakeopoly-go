package main

import (
	"image/color"
	"log"
	"time"

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
	snakeSpeed         = 2
)

var (
	LIGHT_WHITE = color.RGBA{50, 100, 50, 50}
	WHITE       = color.RGBA{160, 210, 160, 255}
	BLACK       = color.RGBA{20, 40, 20, 255}
)

type Snake struct {
	Body [][2]int // Array of x, y pairs
}

type Game struct {
	snake        Snake
	lastMoveTime time.Time // Timestamp of the last movement
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
	// Get the current time
	currentTime := time.Now()

	// Calculate the time elapsed since the last movement
	elapsedTime := currentTime.Sub(g.lastMoveTime)

	// Calculate the desired time interval for movement based on snakeSpeed
	desiredInterval := time.Second / time.Duration(snakeSpeed)

	// Check if it's time to move the snake
	if elapsedTime >= desiredInterval {
		// Move the snake to the right
		newHeadX := g.snake.Body[0][0] + 1 // Move one unit to the right
		newHeadY := g.snake.Body[0][1]

		// Add the new head to the snake's body
		g.snake.Body = append([][2]int{{newHeadX, newHeadY}}, g.snake.Body...)

		// Remove the tail of the snake to maintain its length
		if len(g.snake.Body) > initialSnakeLength {
			g.snake.Body = g.snake.Body[:initialSnakeLength]
		}

		// Update the last movement time to the current time
		g.lastMoveTime = currentTime
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with black color
	screen.Fill(BLACK)

	// Draw the border of the play area
	vector.StrokeRect(screen, float32(playAreaStartX), float32(playAreaStartY), float32(playAreaWidth), float32(playAreaHeight), float32(4), WHITE, false)

	drawGrid(screen)

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

	// Vertical lines
	for x := 0; x < screenWidth; x += screenUnit {
		vector.StrokeLine(screen, float32(x), float32(0), float32(x), float32(screenHeight), float32(1), LIGHT_WHITE, false)
	}

	// Horizontal lines
	for y := 0; y < screenHeight; y += screenUnit {
		vector.StrokeLine(screen, 0, float32(y), float32(screenWidth), float32(y), float32(1), LIGHT_WHITE, false)
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
