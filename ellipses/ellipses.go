package ellipses

import "go-graphics/pixel"

func Draw(start pixel.Point, majorDistance float64) {
	minorDistance := start.Y
	var derivativeOfX float64 = 0
	derivativeOfY := 2 * majorDistance * majorDistance * minorDistance
	decisionVariable := (minorDistance * minorDistance) -
		(majorDistance * majorDistance * minorDistance) +
		(0.25 * majorDistance * majorDistance)
	currentPixel := start

	for derivativeOfX < derivativeOfY {
		drawMirroredPixels(currentPixel)
		currentPixel.X++

		derivativeOfX += minorDistance * minorDistance * 2

		if decisionVariable < 0 {
			decisionVariable += derivativeOfX + (minorDistance * minorDistance)
		} else {
			currentPixel.Y--
			derivativeOfY -= majorDistance * majorDistance * 2
			decisionVariable += derivativeOfX + (minorDistance * minorDistance) - derivativeOfY
		}
	}

	drawMirroredPixels(currentPixel)

	decisionVariable = minorDistance*minorDistance*(currentPixel.X+0.5)*(currentPixel.X+0.5) +
		majorDistance*majorDistance*(currentPixel.Y-1)*(currentPixel.Y-1) -
		minorDistance*minorDistance*majorDistance*majorDistance

	for currentPixel.Y > 0 {
		currentPixel.Y--
		derivativeOfY -= majorDistance * majorDistance * 2

		if decisionVariable >= 0 {
			decisionVariable -= derivativeOfY + majorDistance*majorDistance
		} else {
			currentPixel.X++
			derivativeOfX += minorDistance * minorDistance * 2
			decisionVariable += derivativeOfX - derivativeOfY + majorDistance*majorDistance
		}

		drawMirroredPixels(currentPixel)
	}
}

func drawMirroredPixels(point pixel.Point) {
	sectorOne := point
	sectorTwo := pixel.Point{X: -point.X, Y: point.Y}
	sectorThree := pixel.Point{X: -point.X, Y: -point.Y}
	sectorFour := pixel.Point{X: point.X, Y: -point.Y}

	pixel.DrawPixel(sectorOne)
	pixel.DrawPixel(sectorTwo)
	pixel.DrawPixel(sectorThree)
	pixel.DrawPixel(sectorFour)
}
