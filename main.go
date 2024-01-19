package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	currentDir   direction // Current direction of the snake
	nextDir      direction // Next direction to change to
}

type direction int

const (
	dirUp direction = iota
	dirDown
	dirLeft
	dirRight
)

func (d direction) isOpposite(other direction) bool {
	switch d {
	case dirUp:
		return other == dirDown
	case dirDown:
		return other == dirUp
	case dirLeft:
		return other == dirRight
	case dirRight:
		return other == dirLeft
	}
	return false
}

func (d direction) vector() (int, int) {
	switch d {
	case dirUp:
		return 0, -1
	case dirDown:
		return 0, 1
	case dirLeft:
		return -1, 0
	case dirRight:
		return 1, 0
	}
	return 0, 0
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
	// Handle user input for changing direction
	g.handleInput()

	// Get the current time
	currentTime := time.Now()

	// Calculate the time elapsed since the last movement
	elapsedTime := currentTime.Sub(g.lastMoveTime)

	// Calculate the desired time interval for movement based on snakeSpeed
	desiredInterval := time.Second / time.Duration(snakeSpeed)

	// Check if it's time to move the snake
	if elapsedTime >= desiredInterval {
		// Move the snake according to the current direction
		moveX, moveY := g.currentDir.vector()
		newHeadX := g.snake.Body[0][0] + moveX
		newHeadY := g.snake.Body[0][1] + moveY

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

func (g *Game) handleInput() {
	// Handle arrow key input to change direction
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && g.currentDir != dirDown {
		g.nextDir = dirUp
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && g.currentDir != dirUp {
		g.nextDir = dirDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && g.currentDir != dirRight {
		g.nextDir = dirLeft
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && g.currentDir != dirLeft {
		g.nextDir = dirRight
	}

	// Apply the next direction if it's not opposite to the current direction
	if g.currentDir != g.nextDir && !g.currentDir.isOpposite(g.nextDir) {
		g.currentDir = g.nextDir
	}
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
