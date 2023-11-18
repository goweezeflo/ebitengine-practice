package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
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
	speed := 2.0
	tick := float64(g.tick) / 100
	doubleTime := tick * speed
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick: %d", g.tick), 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sin: %.2f", math.Sin(tick)), 10, 30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cos: %.2f", math.Cos(tick)), 10, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tan: %.2f", math.Tan(tick)), 10, 70)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sin: %.2f", math.Sin(doubleTime)), 10, 90)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cos: %.2f", math.Cos(doubleTime)), 10, 110)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tan: %.2f", math.Tan(doubleTime)), 10, 130)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick: %.2f", tick), 10, 150)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebitengine - 08 - Sin/Cos")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
