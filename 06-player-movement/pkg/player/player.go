package player

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Player struct {
	Width  int
	Height int
	PosX   float64
	PosY   float64
}

var (
	locationX, locationY float64
	movement             Movement
)

func (p *Player) Create(screenWidth int, screenHeight int) {
	p.Width = 32
	p.Height = 32
	p.PosX = float64((screenWidth - p.Width) / 2)
	p.PosY = float64((screenHeight - p.Height) / 2)
}

func (p *Player) Update() {
	movement.Update(p)
	if movement.IsUpKeyPressed() {
		fmt.Println(movement.GetUpKeyPressDuration())
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	drawPlayer(screen, p)
}

func drawPlayer(screen *ebiten.Image, p *Player) {
	sprite := ebiten.NewImage(p.Width, p.Height)
	sprite.Fill(color.RGBA{B: 0xff, A: 0xff})
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.PosX, p.PosY)
	screen.DrawImage(sprite, &options)
}
