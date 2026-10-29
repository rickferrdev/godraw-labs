package main

import (
	"image"
	"image/color"

	"github.com/rickferrdev/godraw-labs"
)

func main() {
	dst := image.NewRGBA(image.Rect(0, 0, 100, 100))

	leftX, leftY := 30, 50
	rightX, rightY := 70, 50

	radius := 30

	blue := color.RGBA{B: 255, A: 255}
	purple := color.RGBA{R: 128, B: 128, A: 255}

	geo := godraw.New(dst)

	geo.Circle(radius, leftY, leftX, blue)
	geo.Circle(radius, rightY, rightX, purple)
	if err := geo.Save("circle.png", dst); err != nil {
		panic(err)
	}
}
