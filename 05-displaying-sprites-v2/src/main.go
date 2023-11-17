package main

import (
	"fmt"
	"github.com/goweezeflo/ebitengine-practice/05-displaying-sprites-v2/pkg/player"
	"github.com/goweezeflo/ebitengine-practice/05-displaying-sprites-v2/pkg/sprites"
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
	ebiten.SetWindowTitle("Ebitengine - 05 - Displaying Sprites v2")
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
	if player.PosX < 0 || player.PosX > (screenWidth-imageSize) {
		player.VectorX *= -1
	}
	if player.PosY < 0 || player.PosY > (screenHeight-imageSize) {
		player.VectorY *= -1
	}
	player.PosX += player.VectorX
	player.PosY += player.VectorY
	options.GeoM.Translate(player.PosX, player.PosY)
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	frameNum = int(g.tick/animationSpeed) % numOfFrames
	frameX = frameNum * imageSize
	rect := sprite.Rect(frameX, 0, frameX+imageSize, imageSize)
	subImg = sprites.AnimatedSprite.SubImage(rect)
	screen.DrawImage(subImg.(*ebiten.Image), options)
}
