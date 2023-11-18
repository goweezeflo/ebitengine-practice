package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct {
	tick uint64
}

func init() {
}

func (g *Game) Update() error {
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick: %d", g.tick), 10, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebitengine - 11 - Time v2")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
