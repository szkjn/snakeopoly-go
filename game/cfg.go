package game

import (
	"image/color"

	"github.com/szkjn/snakeopoly-go/assets"
)

// Constants related to screen dimensions
const (
	ScreenRatio  float32 = 5.0 / 4.0
	ScreenWidth  float32 = 500.0
	ScreenHeight float32 = ScreenWidth / ScreenRatio
	ScreenUnit   float32 = ScreenWidth / 25
)

// Constants related to the play area
const (
	PlayAreaWidth  float32 = ScreenWidth - (ScreenUnit * 2)  // Example: 16 units wide
	PlayAreaHeight float32 = ScreenHeight - (ScreenUnit * 4) // Example: 12 units tall
	PlayAreaX1     float32 = ScreenUnit
	PlayAreaY1     float32 = ScreenUnit
	PlayAreaX2     float32 = PlayAreaX1 + PlayAreaWidth
	PlayAreaY2     float32 = PlayAreaY1 + PlayAreaHeight
)

// Constants related to the snake
const (
	SnakeSize          float32 = ScreenUnit // Assuming snake size is one unit
	InitialSnakeLength float32 = 3.0
	SnakeSpeed         float32 = 5.0
)

// Colors
var (
	LightWhite = color.RGBA{50, 100, 50, 50}
	White      = color.RGBA{160, 210, 160, 255}
	Black      = color.RGBA{20, 40, 20, 255}
)

// Font
var FontMain = assets.MustLoadFont("fonts/VT323/VT323-Regular.ttf")
