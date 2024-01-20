package assets

import (
	"embed"
	"encoding/csv"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

var DataPoint = MustLoadImage("images/30x30/user.png")

func MustLoadImage(path string) *ebiten.Image {
	f, err := assets.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("img: %v, f: %v\n", img, f)
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = MustLoadImage(match)
	}

	return images
}

func MustLoadFont(size float64) font.Face {
	f, err := assets.ReadFile("fonts/VT323/VT323-Regular.ttf")
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}
	return face
}

func LoadSpecialDataPointsCSV() ([][]string, error) {
	fileData, err := assets.ReadFile("competitors.csv")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(fileData)))
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
