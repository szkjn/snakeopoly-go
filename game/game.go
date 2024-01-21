package game

import (
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Snake                    Snake
	LastMoveTime             time.Time // Timestamp of the last movement
	CurrentDir               Direction // Current direction of the snake
	NextDir                  Direction // Next direction to change to
	initialSpecialDataPoints []SpecialDataPoint
	SpecialDataPoints        []SpecialDataPoint
	CurrentDataPoint         DataPointInterface
	CurrentSpecialDataPoint  SpecialDataPoint
	LastSpecialDataPoint     bool
	UI                       *UI
	State                    GameState
	Score                    int8
	Level                    string
	Blinking                 bool
	BlinkTimer               time.Duration
	SnakeVisible             bool
	TextAnimationTimer       time.Duration
	CurrentCharIndex         int
}

type GameState int

const (
	WelcomeState GameState = iota
	PlayState
	GameOverState
	SpecialState
	GoalState
	BlinkState
)

func NewGame() *Game {
	snake := NewSnake()
	dataPoint := NewDataPoint(snake)
	specialDataPoints, err := LoadSpecialDataPoints()
	if err != nil {
		log.Fatalf("Failed to load special data points: %v", err)
	}
	// Deep copy specialDataPoints to initialSpecialDataPoints
	initialSpecialDataPoints := make([]SpecialDataPoint, len(specialDataPoints))
	copy(initialSpecialDataPoints, specialDataPoints)

	initialLevel := specialDataPoints[0].Level

	game := &Game{
		Snake:                    snake,
		LastMoveTime:             time.Now(),
		CurrentDir:               DirRight,
		NextDir:                  DirRight,
		initialSpecialDataPoints: initialSpecialDataPoints,
		SpecialDataPoints:        specialDataPoints,
		CurrentDataPoint:         dataPoint,
		LastSpecialDataPoint:     false,
		UI:                       NewUI(),
		State:                    WelcomeState,
		Score:                    0,
		Level:                    initialLevel,
		Blinking:                 false,
		SnakeVisible:             true,
		CurrentCharIndex:         0,
	}
	return game
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {

	case WelcomeState:
		g.UI.DrawWelcomePage(screen)

	case PlayState:
		g.UI.DrawPlayPage(screen, g)

	case SpecialState:
		g.UI.DrawSpecialPage(screen, g.CurrentSpecialDataPoint, g.CurrentCharIndex)

	case GameOverState:
		g.UI.DrawGameOverPage(screen, g.Score)

	case GoalState:
		g.UI.DrawGoalPage(screen, g.Score)

	case BlinkState:
		g.UI.DrawPlayPage(screen, g)

	}
}

func (g *Game) Update() error {
	if g.State == PlayState {
		// Handle user input for changing direction
		g.updateDirection()

		// Get the current time and calculate the time elapsed since the last movement
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(g.LastMoveTime)
		desiredInterval := time.Second / time.Duration(SnakeSpeed)

		// Check if it's time to move the snake
		if elapsedTime >= desiredInterval {
			// Calculate the new head position
			moveX, moveY := g.CurrentDir.Vector()
			headX, headY := g.Snake.Body[0][0], g.Snake.Body[0][1]
			nextHeadX := headX + float32(moveX)
			nextHeadY := headY + float32(moveY)

			// Check collision with play area border
			if nextHeadX < PlayAreaX1/ScreenUnit || nextHeadX >= PlayAreaX2/ScreenUnit || nextHeadY < PlayAreaY1/ScreenUnit || nextHeadY >= PlayAreaY2/ScreenUnit {
				g.State = GoalState
			} else {
				// Move the snake and handle collision with the current data point
				g.handleSnakeMovementAndCollision(nextHeadX, nextHeadY)

				// Update the last movement time
				g.LastMoveTime = currentTime
			}
		}

	} else if g.State == BlinkState {
		if time.Since(g.LastMoveTime) >= BlinkFreq {
			g.SnakeVisible = !g.SnakeVisible
			g.BlinkTimer += time.Since(g.LastMoveTime)
			g.LastMoveTime = time.Now()
		}
		// Check if blinking duration has elapsed
		if g.BlinkTimer >= TotalBlinkDuration {
			g.State = PlayState
			g.Blinking = false
			g.SnakeVisible = true
		}

	} else if g.State == SpecialState {
		g.handleMacroInput()
		g.TextAnimationTimer += time.Since(g.LastMoveTime)
		if g.TextAnimationTimer >= TextAnimationSpeed {
			g.TextAnimationTimer -= TextAnimationSpeed
			g.CurrentCharIndex++
		}

	} else if g.State == WelcomeState || g.State == GameOverState || g.State == GoalState {
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
	g.CurrentDataPoint = NewDataPoint(g.Snake)

	// Reset specialDataPoints to their initial state
	g.SpecialDataPoints = make([]SpecialDataPoint, len(g.initialSpecialDataPoints))
	copy(g.SpecialDataPoints, g.initialSpecialDataPoints)
}

func (g *Game) ResumeGame() {
	g.State = BlinkState
	g.LastMoveTime = time.Now()
	g.Blinking = true
	g.BlinkTimer = 0
	g.SnakeVisible = false
}

func (g *Game) handleMacroInput() {
	// Handle general macro key interactions
	if g.State == WelcomeState || g.State == GameOverState || g.State == GoalState {

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
	} else if g.State == SpecialState {

		// Get the input characters
		inputChars := ebiten.AppendInputChars(nil)

		// Detect if a key has been pressed
		if len(inputChars) == 1 {
			// If "R" has been pressed
			if inputChars[0] == 114 {
				g.ResumeGame()
				// If "Q" has been pressed
			} else if inputChars[0] == 113 {
				quitGame()
			}
		}
	}
}

func (g *Game) updateDirection() {
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

	// Update the current direction based on the next direction
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

func (g *Game) generateDataPoint() {
	if len(g.SpecialDataPoints) > 0 {
		if g.Score%SpecialDataPointsRate == 0 {
			// Use the first special data point
			g.CurrentDataPoint = NewSpecialDataPoint(g.Snake, g.SpecialDataPoints[0])
			g.SpecialDataPoints = g.SpecialDataPoints[1:]

			// Check if special data points have run out
			if len(g.SpecialDataPoints) == 0 {
				g.LastSpecialDataPoint = true
			}
		} else {
			// Generate a regular data point
			g.CurrentDataPoint = NewDataPoint(g.Snake)
		}
	} else {
		// Generate a regular data point
		g.CurrentDataPoint = NewDataPoint(g.Snake)
	}
}

func (g *Game) handleSnakeMovementAndCollision(nextHeadX, nextHeadY float32) {
	if g.CurrentDataPoint != nil && g.CurrentDataPoint.IsColliding(g.Snake) {
		// Collision detected, increase score
		g.Score++

		// Check if the current data point is special
		if specialDP, isSpecial := g.CurrentDataPoint.(SpecialDataPoint); isSpecial {
			g.State = SpecialState // Trigger special state on collision with special data point
			g.CurrentSpecialDataPoint = specialDP
			g.Level = specialDP.Level

			// Check if this is the last special data point
			if g.LastSpecialDataPoint {
				g.State = GoalState
			}
		}

		// Generate a new data point after handling the current collision
		g.generateDataPoint()

		// Add the new head position and handle snake growth
		g.Snake.Body = append([][2]float32{{nextHeadX, nextHeadY}}, g.Snake.Body...)
	} else {
		// Move the snake without growing
		g.Snake.Body = append([][2]float32{{nextHeadX, nextHeadY}}, g.Snake.Body...)
		if len(g.Snake.Body) > int(InitialSnakeLength) {
			g.Snake.Body = g.Snake.Body[:len(g.Snake.Body)-1]
		}
	}
}
