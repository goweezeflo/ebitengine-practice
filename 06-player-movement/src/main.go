package main

import (
	"github.com/goweezeflo/ebitengine-practice/06-player-movement/pkg/player"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 480
)

var (
	p1 player.Player
)

type Game struct {
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine - 06 - Player Movement")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func init() {
	p1.Create(screenWidth, screenHeight)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	p1.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	p1.Draw(screen)
}
