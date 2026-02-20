package src_test

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"
)

func TestDecodePNG(t *testing.T) {
	file, err := os.Open("./test.png")
	if err != nil {
		t.Errorf("Error opening file: %+v\n", err)
	}
	defer file.Close()

	img, str, err := image.Decode(file)
	if err != nil {
		t.Errorf("Error decoding image: %+v\n", err)
	}

	t.Logf("Image format: %s\n", str)
	t.Logf("Width: %d, Height: %d\n", img.Bounds().Dx(), img.Bounds().Dy())
}

func TestDecodeJPEG(t *testing.T) {
	file, err := os.Open("./test.jpeg")
	if err != nil {
		t.Errorf("Error opening file: %+v\n", err)
	}
	defer file.Close()

	img, str, err := image.Decode(file)
	if err != nil {
		t.Errorf("Error decoding image: %+v\n", err)
	}

	t.Logf("Image format: %s\n", str)
	t.Logf("Width: %d, Height: %d\n", img.Bounds().Dx(), img.Bounds().Dy())
}
