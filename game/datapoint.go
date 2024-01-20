package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szkjn/snakeopoly-go/assets"
)

// DataPoint represents a single datapoint in the game
type DataPoint struct {
	X, Y float32 // Coordinates of the datapoint
}

// DrawDataPoint draws the data point on the screen.
func DrawDataPoint(screen *ebiten.Image, dp DataPoint) {
	// Get the dimensions of the data point image
	dpWidth := float32(assets.DataPoint.Bounds().Dx())
	dpHeight := float32(assets.DataPoint.Bounds().Dy())

	// Calculate the scaling factor to fit the image within a ScreenUnit
	scale := float64(ScreenUnit) / (float64(dpWidth) + float64(ScreenUnit)*0.25)

	// Calculate the position to center the data point on the cell
	x := dp.X*ScreenUnit + (ScreenUnit-float32(dpWidth))*0.5
	y := dp.Y*ScreenUnit + (ScreenUnit-float32(dpHeight))*0.5

	// Draw the scaled data point
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(assets.DataPoint, op)
}

// Creates a new random data point that is not colliding with the snake.
func NewDataPoint(snake Snake) DataPoint {
	// Calculate the available grid positions within the play area
	availablePositions := []struct{ x, y int }{}
	for x := int(PlayAreaX1 / ScreenUnit); x < int(PlayAreaX2/ScreenUnit); x++ {
		for y := int(PlayAreaY1 / ScreenUnit); y < int(PlayAreaY2/ScreenUnit); y++ {
			// Check if the position is not within the snake's body
			isColliding := false
			for _, segment := range snake.Body {
				if int(segment[0]) == x && int(segment[1]) == y {
					isColliding = true
					break
				}
			}
			if !isColliding {
				availablePositions = append(availablePositions, struct{ x, y int }{x, y})
			}
		}
	}

	// Randomly select one available position for the datapoint
	if len(availablePositions) > 0 {
		selected := availablePositions[rand.Intn(len(availablePositions))]
		x := float32(selected.x)
		y := float32(selected.y)
		return DataPoint{X: x, Y: y}
	}

	// If no available positions, create a datapoint at the center of the play area
	x := (PlayAreaX1 + PlayAreaX2) / 2
	y := (PlayAreaY1 + PlayAreaY2) / 2
	return DataPoint{X: x, Y: y}
}

// Checks if the snake has collided with the datapoint
func (d DataPoint) IsColliding(snake Snake) bool {
	// Check if the snake's head is at the same coordinates as the datapoint
	headX, headY := snake.Body[0][0], snake.Body[0][1]
	return headX >= d.X && headX < d.X+1 && headY >= d.Y && headY < d.Y+1
}
