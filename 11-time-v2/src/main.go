package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/exp/shiny/materialdesign/colornames"
	"image"
	"log"
	"math"
	"time"
)

type Game struct {
	tick uint64
	sin  float64
}

func init() {
}

func (g *Game) Update() error {
	g.tick++
	g.sin = (math.Sin(float64(time.Now().UnixMilli())/1000) + 1) / 2
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(image.NewUniform(colornames.BlueGrey500)) // Fill the screen with blue/grey color
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick: %d", g.tick), 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sin: %.2f", g.sin), 10, 30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Breathing: %d", int(g.sin*10)), 10, int(50+(g.sin*100)))
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
