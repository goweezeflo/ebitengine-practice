package main

import (
	"fmt"
	"log"

	"github.com/goweezeflo/ebitengine-practice/03-controller-input/pkg/player"
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
	ebiten.SetWindowTitle("Ebitengine - 03 - Controller Input")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func init() {
	p1 = player.CreatePlayer("Player", 1, screenWidth, screenHeight)
	fmt.Println(p1.Name, p1.GamepadId)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	p1.Move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	p1.Draw(screen)
}
