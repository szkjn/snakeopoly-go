package game

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Snake        Snake
	LastMoveTime time.Time // Timestamp of the last movement
	CurrentDir   Direction // Current direction of the snake
	NextDir      Direction // Next direction to change to*
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
	game := &Game{
		Snake:        NewSnake(),   // Initialize the snake
		LastMoveTime: time.Now(),   // Initialize lastMoveTime
		CurrentDir:   DirRight,     // Initialize the direction (e.g., DirRight for right)
		NextDir:      DirRight,     // Initialize the next direction
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
		// Display the welcome message
		text.Draw(screen, "Welcome to the Snakeopoly!", FontMain, 60, 100, White)
		text.Draw(screen, "Slither your way", FontMain, 120, 200, White)
		text.Draw(screen, "to Surveillance Sovereignty!", FontMain, 70, 250, White)
		text.Draw(screen, "Press P to start or Q to quit", FontMain, 100, 350, White)
	case PlayState:
		// Draw the border of the play area
		vector.StrokeRect(screen, float32(PlayAreaX1), float32(PlayAreaY1), float32(PlayAreaWidth), float32(PlayAreaHeight), float32(4), White, false)

		g.UI.DrawGrid(screen)

		// Draw the snake
		for _, segment := range g.Snake.Body {
			segmentX, segmentY := segment[0]*ScreenUnit, segment[1]*ScreenUnit
			vector.DrawFilledRect(screen, float32(segmentX), float32(segmentY), float32(SnakeSize), float32(SnakeSize), White, false)
		}
	case GameOverState:
		// Display "GAME OVER" text
		text.Draw(screen, "GAME OVER", FontMain, int(ScreenWidth/2-100), 50, White)
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

			// DEBUGGING

			// fmt.Printf("\nnextHeadX1: %v, nextHeadY1: %v\n", nextHeadX1, nextHeadY1)
			// fmt.Printf("PlayAreaX1: %v, PlayAreaY1: %v\n", PlayAreaX1/ScreenUnit, PlayAreaY1/ScreenUnit)
			// fmt.Printf("PlayAreaX2: %v, PlayAreaY2: %v\n", PlayAreaX2/ScreenUnit, PlayAreaY2/ScreenUnit)

			// if nextHeadX1 < PlayAreaX1/ScreenUnit {
			// 	fmt.Printf(">>>>>>>> nextHeadX1 < PlayAreaX1/ScreenUnit")
			// } else if nextHeadX1 >= PlayAreaX2/ScreenUnit {
			// 	fmt.Printf(">>>>>>>> nextHeadX1 >= PlayAreaX/ScreenUnit2")
			// } else if nextHeadY1 < PlayAreaY1/ScreenUnit {
			// 	fmt.Printf(">>>>>>>> nextHeadY1 < PlayAreaY1/ScreenUnit")
			// } else if nextHeadY1 >= PlayAreaY2/ScreenUnit {
			// 	fmt.Printf(">>>>>>>> nextHeadY1 >= PlayAreaY2/ScreenUnit")
			// }

			// Check collision with play area border after moving
			if nextHeadX1 < PlayAreaX1/ScreenUnit || nextHeadX1 >= PlayAreaX2/ScreenUnit || nextHeadY1 < PlayAreaY1/ScreenUnit || nextHeadY1 >= PlayAreaY2/ScreenUnit {
				g.State = GameOverState
			} else {
				// Add the new head to the snake's body
				g.Snake.Body = append([][2]float32{{nextHeadX1, nextHeadY1}}, g.Snake.Body...)

				// Remove the tail of the snake to maintain its length
				if len(g.Snake.Body) > int(InitialSnakeLength) {
					g.Snake.Body = g.Snake.Body[:int(InitialSnakeLength)]
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
				fmt.Println("P is pressed")
				g.State = PlayState
				g.ResetGame()
				// If "Q" has been pressed
			} else if inputChars[0] == 113 {
				fmt.Println("Q is pressed")
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
