package game

import (
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Snake        Snake
	LastMoveTime time.Time // Timestamp of the last movement
	CurrentDir   Direction // Current direction of the snake
	NextDir      Direction // Next direction to change to*
	DataPoint    DataPoint
	UI           *UI
	State        GameState
	Score        int
}

type GameState int

const (
	WelcomeState GameState = iota
	PlayState
	GameOverState
)

func NewGame() *Game {
	snake := NewSnake()
	dataPoint := NewDataPoint(snake)
	game := &Game{
		Snake:        snake,        // Initialize the snake
		LastMoveTime: time.Now(),   // Initialize lastMoveTime
		CurrentDir:   DirRight,     // Initialize the direction (e.g., DirRight for right)
		NextDir:      DirRight,     // Initialize the next direction
		DataPoint:    dataPoint,    // Initialize the first data point
		UI:           NewUI(),      // Initialize the UI
		State:        WelcomeState, // Set initial State to WelcomeState
		Score:        0,
	}
	return game
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(Black)

	switch g.State {

	case WelcomeState:
		g.UI.DrawWelcomePage(screen)

	case PlayState:
		g.UI.DrawPlayPage(screen, g.Snake.Body, g.DataPoint)

	case GameOverState:
		g.UI.DrawGameOverPage(screen)

	}
}

func (g *Game) Update() error {
	if g.State == PlayState {
		// Handle user input for changing direction
		g.handleInput()

		// Get the current time
		currentTime := time.Now()

		// Calculate the time elapsed since the last movement
		elapsedTime := currentTime.Sub(g.LastMoveTime)

		// Calculate the desired time interval for movement based on SnakeSpeed
		desiredInterval := time.Second / time.Duration(SnakeSpeed)

		// Check if it's time to move the snake
		if elapsedTime >= desiredInterval {
			// Move the snake according to the current direction
			moveX, moveY := g.CurrentDir.Vector()

			// Get the head coordinates from the snake's body
			headX, headY := g.Snake.Body[0][0], g.Snake.Body[0][1]

			// Calculate the new head position before moving
			nextHeadX1 := headX + float32(moveX)
			nextHeadY1 := headY + float32(moveY)

			// Check collision with play area border after moving
			if nextHeadX1 < PlayAreaX1/ScreenUnit || nextHeadX1 >= PlayAreaX2/ScreenUnit || nextHeadY1 < PlayAreaY1/ScreenUnit || nextHeadY1 >= PlayAreaY2/ScreenUnit {
				g.State = GameOverState
			} else {

				// Check collision with the data point
				if g.DataPoint.IsColliding(g.Snake) {
					g.Score++

					g.Snake.Body = append([][2]float32{{nextHeadX1, nextHeadY1}}, g.Snake.Body...)

					// Create a new data point
					g.DataPoint = NewDataPoint(g.Snake)
				} else {
					g.Snake.Body = append([][2]float32{{nextHeadX1, nextHeadY1}}, g.Snake.Body...)

					// Remove the last tail segment to maintain its length
					if len(g.Snake.Body) > int(InitialSnakeLength) {
						g.Snake.Body = g.Snake.Body[:len(g.Snake.Body)-1]
					}
				}

				// Update the last movement time to the current time
				g.LastMoveTime = currentTime
			}
		}
	} else if g.State == WelcomeState {
		g.handleMacroInput()
	} else if g.State == GameOverState {
		g.handleMacroInput()
	}

	return nil
}

func (g *Game) ResetGame() {
	g.Snake = NewSnake()
	g.LastMoveTime = time.Now()
	g.CurrentDir = DirRight
	g.NextDir = DirRight
	g.Score = 0
	g.State = PlayState
}

func (g *Game) handleMacroInput() {
	// Handle general macro key interactions
	if g.State == WelcomeState || g.State == GameOverState {

		// Get the input characters
		inputChars := ebiten.AppendInputChars(nil)

		// Detect if a key has been pressed
		if len(inputChars) == 1 {
			// If "P" has been pressed
			if inputChars[0] == 112 {
				g.State = PlayState
				g.ResetGame()
				// If "Q" has been pressed
			} else if inputChars[0] == 113 {
				quitGame()
			}
		}
	}
}

func (g *Game) handleInput() {
	// Handle arrow key input to change direction
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && g.CurrentDir != DirDown {
		g.NextDir = DirUp
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && g.CurrentDir != DirUp {
		g.NextDir = DirDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && g.CurrentDir != DirRight {
		g.NextDir = DirLeft
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && g.CurrentDir != DirLeft {
		g.NextDir = DirRight
	}

	// Apply the next direction if it's not opposite to the current direction
	if g.CurrentDir != g.NextDir && !g.CurrentDir.IsOpposite(g.NextDir) {
		g.CurrentDir = g.NextDir
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(ScreenWidth), int(ScreenHeight)
}

func quitGame() {
	os.Exit(0)
}
