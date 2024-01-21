package game

import (
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szkjn/snakeopoly-go/assets"
)

// Defines common methods for all data points
type DataPointInterface interface {
	Position() (float32, float32)
	Render(screen *ebiten.Image)
	GetImage() *ebiten.Image
	IsSpecial() bool
	IsColliding(snake Snake) bool
}

// Defines a regular data point in the game
type DataPoint struct {
	X, Y  float32
	Image *ebiten.Image
}

// Defines a special data point in the game
type SpecialDataPoint struct {
	DataPoint
	Name  string
	Slug  string
	Year  int
	Text  string
	Level string
}

// Returns xy coordinates of the DataPoint
func (d DataPoint) Position() (float32, float32) {
	return d.X, d.Y
}

// Render for a DataPoint
func (d DataPoint) Render(screen *ebiten.Image) {
	// Implement rendering logic for a DataPoint
}

// IsSpecial returns false for regular DataPoint
func (d DataPoint) IsSpecial() bool {
	return false
}

// Render for a SpecialDataPoint
func (s SpecialDataPoint) Render(screen *ebiten.Image) {
	// Implement rendering logic for a SpecialDataPoint
}

// IsSpecial returns true for SpecialDataPoint
func (s SpecialDataPoint) IsSpecial() bool {
	return true
}

func (d DataPoint) GetImage() *ebiten.Image {
	return d.Image
}

func (s SpecialDataPoint) GetImage() *ebiten.Image {
	return s.Image // This will refer to DataPoint.Image due to embedding
}

// DrawDataPoint draws the data point on the screen.
func DrawDataPoint(screen *ebiten.Image, dp DataPointInterface) {
	img := dp.GetImage() // Get the image using the interface method

	// Calculate dimensions and scaling factor
	dpWidth := float32(img.Bounds().Dx())
	dpHeight := float32(img.Bounds().Dy())
	scale := float64(ScreenUnit) / (float64(dpWidth) + float64(ScreenUnit)*0.25)

	// Calculate position
	x, y := dp.Position()
	centeredX := x*ScreenUnit + (ScreenUnit-float32(dpWidth))*0.5
	centeredY := y*ScreenUnit + (ScreenUnit-float32(dpHeight))*0.5

	// Draw the scaled data point
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(centeredX), float64(centeredY))
	screen.DrawImage(img, op)
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
		return DataPoint{X: x, Y: y, Image: assets.DataPoint}
	}

	// If no available positions, create a datapoint at a default position within the play area
	return DataPoint{X: PlayAreaX1 / ScreenUnit, Y: PlayAreaY1 / ScreenUnit, Image: assets.DataPoint}
}

func NewSpecialDataPoint(special SpecialDataPoint) SpecialDataPoint {
	// Ensure that the special data point's position is within the play area
	if special.X < PlayAreaX1/ScreenUnit || special.X >= PlayAreaX2/ScreenUnit ||
		special.Y < PlayAreaY1/ScreenUnit || special.Y >= PlayAreaY2/ScreenUnit {
		// Reset position to a default within the play area
		special.X = PlayAreaX1 / ScreenUnit
		special.Y = PlayAreaY1 / ScreenUnit
	}
	return special
}
func (d DataPoint) IsColliding(snake Snake) bool {
	headX, headY := snake.Body[0][0], snake.Body[0][1]
	return headX >= d.X && headX < d.X+1 && headY >= d.Y && headY < d.Y+1
}

func (s SpecialDataPoint) IsColliding(snake Snake) bool {
	return s.DataPoint.IsColliding(snake)
}

func LoadSpecialDataPoints() ([]SpecialDataPoint, error) {
	records, err := assets.LoadSpecialDataPointsCSV()
	if err != nil {
		return nil, err
	}

	var specialDataPoints []SpecialDataPoint
	for _, record := range records[1:] {
		year, _ := strconv.Atoi(record[2])
		image := assets.MustLoadImage("images/30x30/" + record[1] + ".png")
		specialDataPoint := SpecialDataPoint{
			DataPoint: DataPoint{
				X:     0,
				Y:     0,
				Image: image,
			},
			Name:  record[0],
			Slug:  record[1],
			Year:  year,
			Text:  record[3],
			Level: record[4],
		}
		specialDataPoints = append(specialDataPoints, specialDataPoint)
	}

	return specialDataPoints, nil
}
