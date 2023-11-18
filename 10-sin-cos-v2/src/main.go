package main

import (
	"fmt"
	"github.com/goweezeflo/ebitengine-practice/10-sin-cos-v2/assets/sprites"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"log"
	"math"
	"time"
)

var (
	currentEpochSecond  int64
	previousEpochSecond int64
	options             ebiten.DrawImageOptions
	subImg              image.Image
	frameX              int
	imageSize           = 64
)

type Game struct {
	tick    uint64
	seconds uint64
	sin     int64
	cos     int64
}

func init() {
	currentEpochSecond = time.Now().Unix()
	previousEpochSecond = currentEpochSecond
}

func (g *Game) Update() error {
	g.tick++
	g.sin = int64(math.Sin(float64(g.tick)/25) * 1000)
	g.cos = int64(math.Cos(float64(g.tick)/25) * 1000)
	currentEpochSecond = time.Now().Unix()
	if currentEpochSecond != previousEpochSecond {
		previousEpochSecond = currentEpochSecond
		g.seconds++
	}

	options = ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(140+g.sin/10), float64(105+g.cos/-10))

	//animationSpeed := 60 / sprites.AnimationFPS
	//numOfFrames := 9
	//frameNum := int(g.tick/animationSpeed) % numOfFrames
	frameX = int(((float64(g.sin)+1000)/2000)*9) * imageSize

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkslategray)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick: %d", g.tick), 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick (seconds): %d", g.seconds), 10, 30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sin: %d", g.sin), 10, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cos: %d", g.cos), 10, 70)
	sprite := ebiten.NewImage(32, 32)
	sprite.Fill(color.RGBA{G: 0xff, A: 0xff})
	screen.DrawImage(sprite, &options)
	rect := image.Rect(frameX, 0, frameX+imageSize, 64)
	subImg = sprites.AnimatedSprite.SubImage(rect)
	animSpriteOptions := ebiten.DrawImageOptions{}
	animSpriteOptions.GeoM.Translate(100, 100)
	screen.DrawImage(subImg.(*ebiten.Image), &animSpriteOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebitengine - 10 - Sin/Cos v2")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
