package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"image/color"
	"log"
)

const (
	screenWidth, screenHeight = 640, 480
)

var (
	screenColor color.Color
)

type Game struct {
}

func init() {
	screenColor = colornames.Black
}

func (g *Game) Update() error {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) &&
		!ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) &&
		!ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		screenColor = colornames.Black
	} else {
		screenColor = color.RGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xff}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(screenColor)
	// CursorPosition returns a position of a mouse cursor relative to the game screen (window). The cursor position is
	// 'logical' position and this considers the scale of the screen.
	//
	// CursorPosition returns (0, 0) before the main loop on desktops and browsers.
	//
	// CursorPosition always returns (0, 0) on mobiles.
	//
	// CursorPosition is concurrent-safe.
	cursorX, cursorY := ebiten.CursorPosition()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cursor Position: %d,%d", cursorX, cursorY), 10, 10)

	// Wheel returns x and y offsets of the mouse wheel or touchpad scroll.
	// It returns 0 if the wheel isn't being rolled.
	wheelX, wheelY := ebiten.Wheel()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mouse Wheel: %.2f,%.2f", wheelX, wheelY), 10, 30)

	// Detecting pressing mouse buttons
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Left Mouse Button: %t", ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)), 10, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Right Mouse Button: %t", ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)), 10, 70)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Middle Mouse Button: %t", ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)), 10, 90)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine - 14 - Mouse Events")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
