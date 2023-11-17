package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Movement struct{}

func (m *Movement) Update(p *Player) {
	if movement.IsUpKeyPressed() {
		movement.MoveUp(p)
	}
	if movement.IsDownKeyPressed() {
		movement.MoveDown(p)
	}
	if movement.IsLeftKeyPressed() {
		movement.MoveLeft(p)
	}
	if movement.IsRightKeyPressed() {
		movement.MoveRight(p)
	}
}

func (m *Movement) IsUpKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyUp)
}

func (m *Movement) IsDownKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyDown)
}

func (m *Movement) IsLeftKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyLeft)
}

func (m *Movement) IsRightKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyRight)
}

func (m *Movement) IsSpaceKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func (m *Movement) GetUpKeyPressDuration() int {
	return inpututil.KeyPressDuration(ebiten.KeyUp)
}

func (m *Movement) IsUpKeyJustReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyUp)
}

func (m *Movement) MoveUp(p *Player) {
	p.PosY -= float64(p.Height)
}

func (m *Movement) MoveDown(p *Player) {
	p.PosY += float64(p.Height)
}

func (m *Movement) MoveLeft(p *Player) {
	p.PosX -= float64(p.Width)
}

func (m *Movement) MoveRight(p *Player) {
	p.PosX += float64(p.Width)
}
