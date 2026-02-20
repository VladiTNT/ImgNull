package src

import (
	"image"
	"image/color"
)

func ProcessImage(img image.Image) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewGray16(img.Bounds())

	for i := range width {
		for j := range height {
			r, g, b, _ := img.At(i, j).RGBA()
			t := max(r, g, b)
			newImg.SetGray16(i, j, color.Gray16{Y: uint16(t)})
		}
	}

	return newImg
}
