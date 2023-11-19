package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

const (
	screenWidth, screenHeight = 640, 480
)

var (
	screenColor    color.Color
	startX, startY int
	endX, endY     int
	isStart        bool = true
	isEnd          bool = false
)

type Game struct {
}

func init() {
	screenColor = color.RGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xff}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && isStart {
		isStart = false
		isEnd = true
		startX, startY = ebiten.CursorPosition()
		endX = 0
		endY = 0
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && isEnd {
		isStart = true
		isEnd = false
		endX, endY = ebiten.CursorPosition()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(screenColor)
	if endX != 0 && endY != 0 {
		vector.StrokeLine(screen, float32(startX), float32(startY), float32(endX), float32(endY), 2, color.White, true)
	} else if startX == 0 && startY == 0 {
		x, y := ebiten.CursorPosition()
		vector.StrokeLine(screen, screenWidth/2, screenHeight/2, float32(x), float32(y), 2, color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}, true)
	} else {
		x, y := ebiten.CursorPosition()
		vector.StrokeLine(screen, float32(startX), float32(startY), float32(x), float32(y), 2, color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}, true)
	}
	ebitenutil.DebugPrintAt(screen, "Left click: Start line", 10, 10)
	ebitenutil.DebugPrintAt(screen, "Right click: End line", 10, 30)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine - 15 - Drawing Lines")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
