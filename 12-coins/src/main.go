package main

import (
	"fmt"
	"github.com/goweezeflo/ebitengine-practice/12-coins/assets/sprites"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	sprite "image"
	"image/color"
	"log"
	"math"
	"time"
)

const (
	screenWidth, screenHeight = 640, 480
	imageSize                 = 64
	numOfFrames               = 11
)

var (
	animationSpeed uint64
	frameNum       int
	frameX         int
	options1       *ebiten.DrawImageOptions
	options2       *ebiten.DrawImageOptions
	subImg         sprite.Image
	initialPosX    int
	initialPosY    int
	posY1          float64
	posY2          float64
)

type Game struct {
	tick uint64
	sin  float64
	cos  float64
}

func init() {
	animationSpeed = 60 / sprites.AnimationFPS
	initialPosX = 320/2 - imageSize/2
	initialPosY = 240/2 - imageSize/2
}

func (g *Game) Update() error {
	g.tick++
	g.sin = (math.Sin(float64(time.Now().UnixMilli())/1000) + 1) / 2
	g.cos = (math.Cos(float64(time.Now().UnixMilli())/1000) + 1) / 2
	posY1 = float64(initialPosY-25) + (50 * g.sin)
	posY2 = float64(initialPosY-25) + (50 * g.cos)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff}) // Fill the screen with grey color
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick: %d", g.tick), 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sin: %.2f", g.sin), 10, 30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cos: %.2f", g.cos), 10, 50)
	frameNum = int(g.tick/animationSpeed) % numOfFrames
	frameX = frameNum * imageSize
	rect := sprite.Rect(frameX, 0, frameX+imageSize, imageSize)
	subImg = sprites.AnimatedSprite.SubImage(rect)

	options1 = &ebiten.DrawImageOptions{}
	options1.GeoM.Translate(float64(initialPosX), posY1)
	screen.DrawImage(subImg.(*ebiten.Image), options1)

	options2 = &ebiten.DrawImageOptions{}
	options2.GeoM.Translate(float64(initialPosX+64), posY2)
	screen.DrawImage(subImg.(*ebiten.Image), options2)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebitengine - 12 - Coins")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
