package lines

import (
	"go-graphics/pixel"
	"math"
)

func drawDDA(line Line) {
	slope := math.Abs(line.slope)

	currentPixel := line.start
	pixel.DrawPixel(currentPixel)

	if slope <= 1 {
		var differenceInX float64 = 1
		differenceInY := currentPixel.Y + slope

		for currentPixel.X != line.end.X {
			pixel.DrawPixel(currentPixel)
			currentPixel = incrementCurrentPixel(currentPixel, differenceInX, differenceInY)
		}

	} else {
		differenceInX := currentPixel.X + (1 / slope)
		var differenceInY float64 = 1

		for currentPixel.Y != line.end.Y {
			pixel.DrawPixel(currentPixel)
			incrementCurrentPixel(currentPixel, differenceInX, differenceInY)
		}
	}
}

func incrementCurrentPixel(currentPixel pixel.Point, differenceInX float64, differenceInY float64) pixel.Point {
	currentPixel.X += differenceInX
	currentPixel.Y += differenceInY

	return currentPixel
}
