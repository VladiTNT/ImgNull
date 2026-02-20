package src_test

import (
	"ImgNull/src"
	"image"
	"image/png"
	_ "image/png"
	"os"
	"testing"
)

func TestProcessImage(t *testing.T) {
	file, err := os.Open("./test.png")
	if err != nil {
		t.Errorf("Error opening file: %v\n", err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		t.Errorf("Error decoding image: %v\n", err)
	}

	t.Logf("Format %s\n", format)

	newImg := src.ProcessImage(img)

	newFile, err := os.Create("./testGray.png")
	if err != nil {
		t.Errorf("Error creating file: %v\n", err)
	}
	defer newFile.Close()

	err = png.Encode(newFile, newImg)
	if err != nil {
		t.Errorf("Error encoding image: %v\n", err)
	}
}
