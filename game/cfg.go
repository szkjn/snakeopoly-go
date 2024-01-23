package game

import (
	"image/color"
	"time"

	"github.com/szkjn/snakeopoly-go/assets"
	"golang.org/x/image/font"
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
	SpecialDataPointsRate int8    = 3
)

// Colors
var (
	LighterGreen color.Color = color.RGBA{160, 210, 160, 255}
	LightGreen   color.Color = color.RGBA{160, 200, 160, 255}
	DarkerGreen  color.Color = color.RGBA{20, 40, 20, 255}
	DarkGreen    color.Color = color.RGBA{20, 70, 20, 255}
)

// Font
var (
	FontXXL font.Face = assets.MustLoadFont(float64(ScreenUnit * 1.9))
	FontXL  font.Face = assets.MustLoadFont(float64(ScreenUnit * 1.6))
	FontL   font.Face = assets.MustLoadFont(float64(ScreenUnit * 1.3))
	FontM   font.Face = assets.MustLoadFont(float64(ScreenUnit * 1))
	FontS   font.Face = assets.MustLoadFont(float64(ScreenUnit * 0.7))
	FontXS  font.Face = assets.MustLoadFont(float64(ScreenUnit * 0.3))
)

// UI effects
const (
	TotalBlinkDuration time.Duration = 1 * time.Second
	BlinkFreq          time.Duration = 200 * time.Millisecond
	TextAnimationSpeed time.Duration = 300 * time.Millisecond
	ShapePixelSize     float64       = float64(ScreenUnit) / 6
	GShapeTime         time.Duration = 2000 * time.Millisecond
	SixShapeTime       time.Duration = 400 * time.Millisecond
)
