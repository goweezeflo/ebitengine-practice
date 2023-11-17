package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	StaticSprite   *ebiten.Image
	AnimatedSprite *ebiten.Image
	err            error
	AnimationFPS   uint64
)

func init() {
	StaticSprite, _, err = ebitenutil.NewImageFromFile("04-displaying-sprites/assets/sprites/sprite-static-32x32.png")
	if err != nil {
		log.Fatal(err)
	}
	AnimatedSprite, _, err = ebitenutil.NewImageFromFile("04-displaying-sprites/assets/sprites/sprite-animated-32x32.png")
	if err != nil {
		log.Fatal(err)
	}
	AnimationFPS = 8
}
