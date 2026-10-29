package main

import (
	"image"
	"image/color"

	"github.com/rickferrdev/godraw-labs"
)

func main() {
	dst := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw := godraw.New(dst)
	draw.Rect(2, 6, 7, 12, color.RGBA{
		R: 255,
		A: 255,
	})

	if err := draw.Save("rect.png", dst); err != nil {
		panic(err)
	}
}
