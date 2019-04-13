package lines

import "go-graphics/pixel"

func Draw(start pixel.Point, end pixel.Point) {
	slope := calculateSlope(start, end)

	switch slope {
	case 0:
		drawHorizontalLine(start, end)
	default:
		drawWithAlgorithm(start, end)
	}

	pixel.DrawPixel(start)
	pixel.DrawPixel(end)
}

func calculateSlope(start pixel.Point, end pixel.Point) float32 {
	differenceInY := end.Y - start.Y
	differenceInX := end.X - start.X

	return differenceInY / differenceInX
}

func drawHorizontalLine(start pixel.Point, end pixel.Point) {
	currentPixel := start

	for currentPixel.X != end.X {
		pixel.DrawPixel(currentPixel)
		currentPixel.X++
	}
}

func drawVerticalLine(start pixel.Point, end pixel.Point) {

}

func drawWithAlgorithm(start pixel.Point, end pixel.Point) {

}
