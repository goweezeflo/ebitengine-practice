package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerInput struct {
	PressedKeys []ebiten.Key
}

func (input *PlayerInput) Create() {
	input.PressedKeys = make([]ebiten.Key, 0)
}

func (input *PlayerInput) DetectPressedKeys() []ebiten.Key {
	input.PressedKeys = inpututil.AppendPressedKeys(input.PressedKeys[:0])
	return input.PressedKeys
}
