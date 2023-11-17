package player

import (
	"bytes"
	"fmt"
	sfx "github.com/goweezeflo/ebitengine-practice/03-controller-input/assets/audio"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"image/color"
	"log"

	"github.com/goweezeflo/ebitengine-practice/02-basic-movement/pkg/controller"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Name                             string
	Width, Height, PosX, PosY, Speed float64
	Error                            bool
	Controller                       controller.Controller
}

var (
	screenWidth, screenHeight float64
	bounceX, bounceY          float64 = 1, -1
	blipAudio                 *audio.Player
	audioContext              *audio.Context
	blipRaw                   = sfx.BlipOgg
)

func CreatePlayer(name string, controller controller.Controller) Player {
	enableSound()
	return Player{
		Name:       name,
		Width:      40,
		Height:     40,
		PosX:       320,
		PosY:       240,
		Speed:      3,
		Error:      false,
		Controller: controller,
	}
}

func playErrorSound() {
	if err := blipAudio.Rewind(); err != nil {
		log.Fatal(err)
	}
	blipAudio.Play()
}

func enableSound() {
	audioContext = audio.NewContext(44100)
	blip, err := vorbis.DecodeWithoutResampling(bytes.NewReader(blipRaw))
	if err != nil {
		log.Fatal(err)
	}
	blipAudio, err = audioContext.NewPlayer(blip)
	if err != nil {
		log.Fatal(err)
	}
	if err := blipAudio.Rewind(); err != nil {
		log.Fatal(err)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	sprite := ebiten.NewImage(int(p.Width), int(p.Height))
	if p.Error {
		sprite.Fill(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
		playErrorSound()
		fmt.Println(p.PosX, p.PosY)
		p.Error = false
	} else {
		sprite.Fill(color.RGBA{B: 0xff, A: 0xff})
	}
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.PosX, p.PosY)
	screen.DrawImage(sprite, &options)
	screenWidth = float64(screen.Bounds().Dx())
	screenHeight = float64(screen.Bounds().Dy())
}

func (p *Player) Update(time float64) {
	if p.PosX >= screenWidth-p.Width || p.PosX <= 0 {
		bounceX *= -1
		p.Error = true
	}
	if p.PosY >= screenHeight-p.Height || p.PosY <= 0 {
		bounceY *= -1
		p.Error = true
	}
	p.PosX += p.Speed * time * bounceX
	p.PosY -= p.Speed * time * bounceY
}
