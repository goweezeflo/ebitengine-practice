package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	AnimatedSprite *ebiten.Image
	err            error
	AnimationFPS   uint64
)

func init() {
	AnimatedSprite, _, err = ebitenutil.NewImageFromFile("10-sin-cos-v2/assets/sprites/breathe-in-breathe-out-64x64.png")
	if err != nil {
		log.Fatal(err)
	}
	AnimationFPS = 8
}
