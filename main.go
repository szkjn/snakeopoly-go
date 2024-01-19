package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth    = 640
	screenHeight   = 480
	playAreaWidth  = 400
	playAreaHeight = 400
	playAreaStartX = (screenWidth - playAreaWidth) / 2
	playAreaStartY = (screenHeight - playAreaHeight) / 2
	snakeSize      = 20
	moveSpeed = 4
)

var (
	BLACK = color.RGBA{160, 210, 160, 255}
	WHITE = color.RGBA{20, 40, 20, 255}
)

type Snake struct {
    Body     [][2]int
    Direction string
}

type Game struct {
	snake Snake
}

func (g *Game) Update() error {
    // Handle keyboard input for snake direction
    if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
        g.snake.Direction = "left"
    } else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
        g.snake.Direction = "right"
    } else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
        g.snake.Direction = "up"
    } else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
        g.snake.Direction = "down"
    }

    // Update the snake's position
    head := g.snake.Body[0]
    switch g.snake.Direction {
    case "left":
        head[0] -= moveSpeed
    case "right":
        head[0] += moveSpeed
    case "up":
        head[1] -= moveSpeed
    case "down":
        head[1] += moveSpeed
    }

    // Move the body
    for i := len(g.snake.Body) - 1; i > 0; i-- {
        g.snake.Body[i] = g.snake.Body[i-1]
    }
    g.snake.Body[0] = head

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the play area border in BLACK
	for x := playAreaStartX; x < playAreaStartX+playAreaWidth; x += snakeSize {
		for y := playAreaStartY; y < playAreaStartY+playAreaHeight; y += snakeSize {
			if x == playAreaStartX || x == playAreaStartX+playAreaWidth-snakeSize ||
				y == playAreaStartY || y == playAreaStartY+playAreaHeight-snakeSize {
				borderRect := image.Rect(x, y, x+snakeSize, y+snakeSize)
				screen.SubImage(borderRect).(*ebiten.Image).Fill(BLACK)
			}
		}
	}

	// Draw the snake in BLACK
	for _, segment := range g.snake.Body {
		x, y := segment[0], segment[1]
		snakeRect := image.Rect(x, y, x+snakeSize, y+snakeSize)
		screen.SubImage(snakeRect).(*ebiten.Image).Fill(BLACK)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")

	game := &Game{
		snake: Snake{
			Body: [][2]int{
				{playAreaStartX + 200, playAreaStartY + 200}, // Example starting position
			},
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
