package game

import (
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szkjn/snakeopoly-go/assets"
)

// Define common methods for all data points
type DataPointInterface interface {
	Position() (float32, float32)
	GetImage() *ebiten.Image
	IsColliding(snake Snake) bool
}

// Define a regular data point in the game
type DataPoint struct {
	X, Y  float32
	Image *ebiten.Image
}

// Define a special data point in the game
type SpecialDataPoint struct {
	DataPoint
	Name  string
	Slug  string
	Year  int
	Text  string
	Level string
}

// Return xy coordinates of the DataPoint
func (d DataPoint) Position() (float32, float32) {
	return d.X, d.Y
}

// Get DataPoint or SpecialDataPoint corresponding image
func (d DataPoint) GetImage() *ebiten.Image {
	return d.Image
}

// Check collision with Snake
func (d DataPoint) IsColliding(snake Snake) bool {
	headX, headY := snake.Body[0][0], snake.Body[0][1]
	return headX >= d.X && headX < d.X+1 && headY >= d.Y && headY < d.Y+1
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

func GenerateRandomPosition(snake Snake) [2]float32 {
	// Calculate the available grid positions within the play area
	availablePositions := []struct{ x, y int }{}
	for x := int(PlayAreaX1 / ScreenUnit); x < int(PlayAreaX2/ScreenUnit); x++ {
		for y := int(PlayAreaY1 / ScreenUnit); y < int(PlayAreaY2/ScreenUnit); y++ {
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

	// Randomly select one available position for the data point
	if len(availablePositions) > 0 {
		selected := availablePositions[rand.Intn(len(availablePositions))]
		return [2]float32{float32(selected.x), float32(selected.y)}
	}

	// Return a default position if no available positions
	return [2]float32{PlayAreaX1 / ScreenUnit, PlayAreaY1 / ScreenUnit}
}

// Create a new data point at a valid random position
func NewDataPoint(snake Snake) DataPoint {
	position := GenerateRandomPosition(snake)
	return DataPoint{X: position[0], Y: position[1], Image: assets.DataPoint}
}

// Create a new special data point at a valid random position
func NewSpecialDataPoint(snake Snake, special SpecialDataPoint) SpecialDataPoint {
	// Generate a random position for this special data point
	position := GenerateRandomPosition(snake)
	special.X = position[0]
	special.Y = position[1]
	return special
}

// Place DP image on grid
func PlaceDataPoint(dp DataPointInterface) (float64, float64, float64) {
	img := dp.GetImage()

	// Calculate dimensions and scaling factor
	dpWidth := float32(img.Bounds().Dx())
	dpHeight := float32(img.Bounds().Dy())
	scale := ScreenUnit / (dpWidth + ScreenUnit*0.25)

	// Calculate position
	x, y := dp.Position()
	centeredX := x*ScreenUnit + (ScreenUnit-dpWidth)*0.5
	centeredY := y*ScreenUnit + (ScreenUnit-dpHeight)*0.5

	return float64(scale), float64(centeredX), float64(centeredY)
}
