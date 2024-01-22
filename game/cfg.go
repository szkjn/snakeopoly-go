package game

import (
	"image/color"
	"time"

	"github.com/szkjn/snakeopoly-go/assets"
)

// Constants related to screen and play area dimensions
const (
	ScreenRatio    float32 = 5.0 / 4.0
	ScreenWidth    float32 = 800.0
	ScreenHeight   float32 = ScreenWidth / ScreenRatio
	ScreenUnit     float32 = ScreenWidth / 25
	PlayAreaWidth  float32 = ScreenWidth - (ScreenUnit * 2)
	PlayAreaHeight float32 = ScreenHeight - (ScreenUnit * 5)
	PlayAreaX1     float32 = ScreenUnit
	PlayAreaY1     float32 = ScreenUnit
	PlayAreaX2     float32 = PlayAreaX1 + PlayAreaWidth
	PlayAreaY2     float32 = PlayAreaY1 + PlayAreaHeight
)

// Constants related to the snake and data points
const (
	SnakeSize             float32 = ScreenUnit
	InitialSnakeLength    float32 = 3
	SnakeSpeed            float32 = 7
	SpecialDataPointsRate int8    = 1
)

// Colors
var (
	White      = color.RGBA{160, 210, 160, 255}
	LightWhite = color.RGBA{20, 70, 20, 255}
	Black      = color.RGBA{20, 40, 20, 255}
	LightBlack = color.RGBA{160, 190, 160, 255}
)

// Font
var (
	FontXXL = assets.MustLoadFont(float64(ScreenUnit * 1.9))
	FontXL  = assets.MustLoadFont(float64(ScreenUnit * 1.6))
	FontL   = assets.MustLoadFont(float64(ScreenUnit * 1.3))
	FontM   = assets.MustLoadFont(float64(ScreenUnit * 1))
	FontS   = assets.MustLoadFont(float64(ScreenUnit * 0.7))
)

// UI effects
const (
	TotalBlinkDuration = 1 * time.Second
	BlinkFreq          = 200 * time.Millisecond
	TextAnimationSpeed = 200 * time.Millisecond
)
