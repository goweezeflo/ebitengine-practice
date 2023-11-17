package main

import (
	"log"

	"github.com/goweezeflo/ebitengine-practice/02-basic-movement/pkg/controller"
	"github.com/goweezeflo/ebitengine-practice/02-basic-movement/pkg/player"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 480
)

var (
	p1   player.Player
	time float64 = 1
)

type Game struct {
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine - 02 - Basic Movement")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func init() {
	p1 = player.CreatePlayer("Player 1", controller.CreateController(1))
}

func (g *Game) Update() error {
	p1.Update(time)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	p1.Draw(screen)
}
