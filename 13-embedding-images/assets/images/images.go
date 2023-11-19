package images

// https://www.reddit.com/r/ebiten/comments/mc9joj/using_the_new_embed_package_in_go_116/

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

var (
	//go:embed sprite-static-32x32.png
	spriteStatic32x32 []byte
	SpriteStatic      *ebiten.Image
)

func init() {
	var err error
	spriteDecoded, _, err := image.Decode(bytes.NewReader(spriteStatic32x32))
	if err != nil {
		log.Fatal(err)
	}

	SpriteStatic = ebiten.NewImageFromImage(spriteDecoded)

}
