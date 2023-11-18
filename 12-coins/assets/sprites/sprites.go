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
	AnimatedSprite, _, err = ebitenutil.NewImageFromFile("12-coins/assets/sprites/sprite-coin-64x64.png")
	if err != nil {
		log.Fatal(err)
	}
	AnimationFPS = 8
}
