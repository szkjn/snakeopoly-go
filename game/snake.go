package game

type Snake struct {
	Body [][2]float32 // Array of x, y pairs
}

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

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

func NewSnake() Snake {
	body := make([][2]float32, int(InitialSnakeLength))
	startX := (PlayAreaX1 / SnakeSize) + (PlayAreaX1 * 3 / SnakeSize)
	startY := (PlayAreaY1 / SnakeSize) + (PlayAreaY1 * 4 / SnakeSize)

	// Initialize the snake's body towards the left and at the center
	for i := 0; i < int(InitialSnakeLength); i++ {
		body[i] = [2]float32{startX + float32(i), startY}
	}

	return Snake{Body: body}
}
