package main

import (
	embeddedAssets "github.com/goweezeflo/ebitengine-practice/13-embedding-images/assets/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"log"
)

const (
	screenWidth, screenHeight = 640, 480
	imageSize                 = 32
	posX                      = (screenWidth/2)/2 - imageSize/2
	posY                      = (screenHeight/2)/2 - imageSize/2
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
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(posX, posY)
	screen.DrawImage(embeddedAssets.SpriteStatic, &options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebitengine - 13 - Embedding Images")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
