package hud

import (
	"fmt"
	"github.com/goweezeflo/ebitengine-practice/07-hud/pkg/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type HUD struct {
	Tick string
}

func (i *HUD) Create() {
	i.Tick = "Game Tick"
}

func (i *HUD) DisplayGameTick(screen *ebiten.Image, tick uint64) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s: %d", i.Tick, tick), 10, 10)
}

func (i *HUD) DisplayPressedKeys(screen *ebiten.Image, playerInput input.PlayerInput) {
	keys := playerInput.PressedKeys
	pressedKeys := "Pressed Keys: "
	if len(keys) == 0 {
		ebitenutil.DebugPrintAt(screen, pressedKeys, 10, 30)
	} else if len(keys) == 1 {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%s", pressedKeys, keys[0]), 10, 30)
	} else if len(keys) > 1 {
		message := fmt.Sprintf("%s%s", pressedKeys, keys[0])
		for i := 1; i < len(keys); i++ {
			message = fmt.Sprintf("%s + %s", message, keys[i])
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s", message), 10, 30)
	}
}
