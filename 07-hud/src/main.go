package main

import (
	"fmt"
	"github.com/goweezeflo/ebitengine-practice/07-hud/pkg/hud"
	"github.com/goweezeflo/ebitengine-practice/07-hud/pkg/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	gameHUD     hud.HUD
	playerInput input.PlayerInput
)

type Game struct {
	tick uint64
}

func init() {
	playerInput.Create()
	gameHUD.Create()
}

func (g *Game) Update() error {
	g.tick++
	playerInput.DetectPressedKeys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gameHUD.DisplayGameTick(screen, g.tick)
	gameHUD.DisplayPressedKeys(screen, playerInput)
	gameHUD.DisplayPressedKeys(screen, playerInput)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()), 10, 50)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebitengine - 07 - HUD")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
