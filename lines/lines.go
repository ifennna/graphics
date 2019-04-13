package lines

import "go-graphics/pixel"

func Draw(start pixel.Point, end pixel.Point) {
	differenceInY := end.Y - start.Y
	differenceInX := end.X - start.X

	if differenceInY == 0 {
		drawHorizontalLine(start, end)
	} else if differenceInX == 0 {
		drawVerticalLine(start, end)
	} else {
		slope := differenceInY / differenceInX
		drawWithAlgorithm(start, end, slope)
	}
}

func drawHorizontalLine(start pixel.Point, end pixel.Point) {
	currentPixel := start

	for currentPixel.X != end.X {
		pixel.DrawPixel(currentPixel)
		currentPixel.X++
	}
}

func drawVerticalLine(start pixel.Point, end pixel.Point) {
	currentPixel := start

	for currentPixel.Y != end.Y {
		pixel.DrawPixel(currentPixel)
		currentPixel.Y++
	}
}

func drawWithAlgorithm(start pixel.Point, end pixel.Point, slope float32) {

}
