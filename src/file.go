package src

import (
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

func ImageToGray(path, newPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return err
	}

	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	newImg := ProcessImage(img)

	switch format {
	case "png":
		err = png.Encode(newFile, newImg)
		if err != nil {
			return err
		}

	case "jpeg":
		err = jpeg.Encode(newFile, newImg, &jpeg.Options{Quality: 100})
		if err != nil {
			return err
		}
	}

	return nil
}
