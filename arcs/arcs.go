package arcs

import (
	"go-graphics/pixel"
	"math"
)

func Draw(start pixel.Point, end pixel.Point) {
	radius := end.Y - start.Y
	currentPixel := start

	for currentPixel.Y >= end.Y {
		currentPixel.Y = math.Sqrt(radius*radius - currentPixel.X*currentPixel.X)
		pixel.DrawPixel(currentPixel)
		currentPixel.X++
	}
}
