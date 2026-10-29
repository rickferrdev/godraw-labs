package godraw

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

type GoDraw struct {
	Image *image.RGBA
}

func New(image *image.RGBA) *GoDraw {
	return &GoDraw{
		image,
	}
}

func (gd *GoDraw) Circle(radius, centerY, centerX int, color color.RGBA) {
	bounds := gd.Image.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dx, dy := x-centerX, y-centerY

			if dx*dx+dy*dy <= radius*radius {
				gd.Image.Set(x, y, color)
			}
		}
	}
}

func (gd *GoDraw) Line(x0, y0, x1, y1 int, color color.Color) {
	dx := x1 - x0
	dy := y1 - y0

	steps := int(math.Max(
		math.Abs(float64(dx)),
		math.Abs(float64(dy)),
	))

	if steps == 0 {
		gd.Image.Set(x0, y0, color)
		return
	}

	for index := 0; index <= steps; index++ {
		x := x0 + index*dx/steps
		y := y0 + index*dy/steps

		gd.Image.Set(x, y, color)
	}
}

func (gd *GoDraw) Rect(x0, y0, x1, y1 int, color color.Color) {
	rect := image.Rect(x0, y0, x1, y1)
	src := image.Uniform{
		C: color,
	}
	draw.Draw(gd.Image, rect.Bounds(), &src, image.Point{}, draw.Src)
}

func (gd *GoDraw) Save(name string, dest *image.RGBA) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, dest); err != nil {
		return err
	}

	return nil
}
