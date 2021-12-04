package npc

import (
	"image"
	"image/png"
	"os"

	"github.com/df-mc/dragonfly/server/player/skin"
)

const (
	CustomSlimGeometry = "geometry.humanoid.customSlim"
	CustomGeometry     = "geometry.humanoid.custom"
)

func EncodeSkinPNG(skin skin.Skin, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, skin)
	if err != nil {
		return err
	}
	return nil
}
func DecodePNGSkin(path, geometry string) (skin.Skin, error) {
	var f *os.File
	var err error

	if f, err = os.OpenFile(path, os.O_RDWR, 0777); os.IsNotExist(err) {
		return skin.Skin{}, err
	}

	img, err := png.Decode(f)
	if err != nil {
		return skin.Skin{}, err
	}
	s := skin.New(img.Bounds().Max.X, img.Bounds().Max.Y)
	s.Pix = pix(img)
	s.ModelConfig = skin.ModelConfig{Default: geometry}
	return s, nil
}

func pix(i image.Image) (p []uint8) {
	for y := 0; y <= i.Bounds().Max.Y-1; y++ {
		for x := 0; x <= i.Bounds().Max.X-1; x++ {
			color := i.At(x, y)
			r, g, b, a := color.RGBA()
			p = append(p, uint8(r), uint8(g), uint8(b), uint8(a))
		}
	}
	return
}
