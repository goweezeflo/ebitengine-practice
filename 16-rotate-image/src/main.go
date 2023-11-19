package main

import (
	embeddedAssets "github.com/goweezeflo/ebitengine-practice/13-embedding-images/assets/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"log"
	"math"
)

const (
	screenWidth, screenHeight = 640, 480
	canvasWidth, canvasHeight = 320, 240
	posX                      = 160
	posY                      = 120
	rotationSpeed             = 4
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
	screen.Fill(colornames.Indigo) // Fill the screen with purple color
	ebitenutil.DebugPrintAt(screen, "Embedded Image", 10, 10)
	spriteSize := embeddedAssets.SpriteStatic.Bounds().Size()
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(-spriteSize.X/2), float64(-spriteSize.Y/2)) // Move the center to upper-left corner to prepare for rotation
	options.GeoM.Rotate(float64(g.tick%360) * rotationSpeed * math.Pi / 360)   // Rotate the sprite
	options.GeoM.Translate(posX, posY)                                         // Move to the center of the screen
	screen.DrawImage(embeddedAssets.SpriteStatic, &options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return canvasWidth, canvasHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine - 16 - Rotate Image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
