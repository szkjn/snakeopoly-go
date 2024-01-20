package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szkjn/snakeopoly-go/game"
)

func runGame() error {
	ebiten.SetWindowSize(int(game.ScreenWidth), int(game.ScreenHeight))
	ebiten.SetWindowTitle("Snake Game")

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := runGame(); err != nil {
		log.Fatal(err)
	}
}
