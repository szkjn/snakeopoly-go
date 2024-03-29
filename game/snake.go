package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szkjn/snakeopoly-go/assets"
)

// Define the snake with a slice of x,y coordinates for its body
type Snake struct {
	Body  [][2]float32
	Image *ebiten.Image
}

// Possible directions the snake can move in
type Direction int

// Enumeration of directions
const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

func (s Snake) GetImage() *ebiten.Image {
	return assets.GooglevilImg
}

// Check if the given direction is opposite to the current direction
func (d Direction) IsOpposite(other Direction) bool {
	switch d {
	case DirUp:
		return other == DirDown
	case DirDown:
		return other == DirUp
	case DirLeft:
		return other == DirRight
	case DirRight:
		return other == DirLeft
	}
	return false
}

// Return the unit vector (as x, y increments) for the given direction
func (d Direction) Vector() (int, int) {
	switch d {
	case DirUp:
		return 0, -1
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	case DirRight:
		return 1, 0
	}
	return 0, 0
}

// Initialize and return a new snake with a default length and position
func NewSnake() Snake {
	body := make([][2]float32, int(InitialSnakeLength))

	// Set the initial position of the snake's head
	startX := (PlayAreaX1 / SnakeSize) + (PlayAreaX1 * 3 / SnakeSize)
	startY := (PlayAreaY1 / SnakeSize) + (PlayAreaY1 * 4 / SnakeSize)

	// Align the body towards the left of the head
	for i := 0; i < int(InitialSnakeLength); i++ {
		body[i] = [2]float32{startX - float32(i), startY}
	}

	return Snake{Body: body}
}

func (s *Snake) CollidesWithItself(nextHeadX, nextHeadY float32) bool {
	for _, segment := range s.Body[1:] { // Start from 1 to skip the head
		if segment[0] == nextHeadX && segment[1] == nextHeadY {
			return true
		}
	}
	return false
}
