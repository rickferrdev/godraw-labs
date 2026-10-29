package main

import (
	"image"
	"image/color"

	"github.com/rickferrdev/godraw-labs"
)

func main() {
	dst := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw := godraw.New(dst)
	draw.Line(50, 10, 50, 90, color.White)

	if err := draw.Save("line.png", dst); err != nil {
		panic(err)
	}
}
