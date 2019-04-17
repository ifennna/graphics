package lines

import "go-graphics/pixel"

type Line struct {
	start pixel.Point
	end   pixel.Point
	slope float64
}

func Draw(start pixel.Point, end pixel.Point, algorithm string) {
	differenceInY, differenceInX := getXAndYDifferences(start, end)

	if differenceInY == 0 {
		drawHorizontalLine(start, end)
	} else if differenceInX == 0 {
		drawVerticalLine(start, end)
	} else {
		slope := differenceInY / differenceInX
		line := Line{start: start, end: end, slope: slope}

		drawWithAlgorithm(algorithm, line)
	}
}

func getXAndYDifferences(start pixel.Point, end pixel.Point) (float64, float64) {
	differenceInY := end.Y - start.Y
	differenceInX := end.X - start.X
	return differenceInY, differenceInX
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

func drawWithAlgorithm(algorithm string, line Line) {
	switch algorithm {
	case "bresenhams":
		drawBresenhams(line)
	case "dda":
		drawDDA(line)
	}
}
