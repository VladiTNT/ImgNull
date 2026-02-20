package main

import (
	"ImgNull/src"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Print(`ImgNull is a basic tool to convert images into black and white versions,
COMMAND: "ImgNull <src> <out>", supported file formats: PNG and JPEG.`)
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments.")
		return
	}

	err := src.ImageToGray(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
	}
}
