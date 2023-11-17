package main

import (
	"fmt"
	"github.com/goweezeflo/04-displaying-sprites/assets/sprites"
	sprite "image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 480
	imageSize                 = 32
	numOfFrames               = 16
)

var (
	animationSpeed uint64
	frameNum       int
	frameX         int
	options        *ebiten.DrawImageOptions
	subImg         sprite.Image
)

type Game struct {
	tick uint64
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine - 04 - Displaying Sprites")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func init() {
	animationSpeed = 60 / sprites.AnimationFPS
	fmt.Println("Displaying Sprites")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	options = &ebiten.DrawImageOptions{}
	if g.tick > 500 {
		options.GeoM.Rotate(float64(g.tick / 100))
		options.GeoM.Translate(300, 300)
	} else if g.tick > 250 {
		options.GeoM.Translate(200, 100)
	} else if g.tick > 100 {
		options.GeoM.Translate(100, 200)
	}
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.tick > 500 {
		rect := sprite.Rect(0, 0, imageSize, imageSize)
		subImg = sprites.StaticSprite.SubImage(rect)
	} else {
		frameNum = int(g.tick/animationSpeed) % numOfFrames
		frameX = frameNum * imageSize
		rect := sprite.Rect(frameX, 0, frameX+imageSize, imageSize)
		subImg = sprites.AnimatedSprite.SubImage(rect)
	}
	screen.DrawImage(subImg.(*ebiten.Image), options)
}
