package ellipses

import "go-graphics/pixel"

func Draw(midpoint pixel.Point, horizontalDistance float64, verticalDistance float64) {
	var derivativeOfX float64 = 0
	derivativeOfY := 2 * horizontalDistance * horizontalDistance * verticalDistance
	decisionVariable := (verticalDistance * verticalDistance) -
		(horizontalDistance * horizontalDistance * verticalDistance) +
		(0.25 * horizontalDistance * horizontalDistance)
	currentPixel := pixel.Point{X: 0, Y: verticalDistance}

	for derivativeOfX < derivativeOfY {
		drawMirroredPixels(midpoint, currentPixel)
		currentPixel.X++

		derivativeOfX += verticalDistance * verticalDistance * 2

		if decisionVariable < 0 {
			decisionVariable += derivativeOfX + (verticalDistance * verticalDistance)
		} else {
			currentPixel.Y--
			derivativeOfY -= horizontalDistance * horizontalDistance * 2
			decisionVariable += derivativeOfX + (verticalDistance * verticalDistance) - derivativeOfY
		}
	}

	drawMirroredPixels(midpoint, currentPixel)

	decisionVariable = verticalDistance*verticalDistance*(currentPixel.X+0.5)*(currentPixel.X+0.5) +
		horizontalDistance*horizontalDistance*(currentPixel.Y-1)*(currentPixel.Y-1) -
		verticalDistance*verticalDistance*horizontalDistance*horizontalDistance

	for currentPixel.Y > 0 {
		currentPixel.Y--
		derivativeOfY -= horizontalDistance * horizontalDistance * 2

		if decisionVariable >= 0 {
			decisionVariable -= derivativeOfY + horizontalDistance*horizontalDistance
		} else {
			currentPixel.X++
			derivativeOfX += verticalDistance * verticalDistance * 2
			decisionVariable += derivativeOfX - derivativeOfY + horizontalDistance*horizontalDistance
		}

		drawMirroredPixels(midpoint, currentPixel)
	}
}

func drawMirroredPixels(midpoint pixel.Point, point pixel.Point) {
	x := midpoint.X + point.X
	y := midpoint.Y + point.Y

	sectorOne := pixel.Point{X: x, Y: y}
	sectorTwo := pixel.Point{X: -x, Y: y}
	sectorThree := pixel.Point{X: -x, Y: -y}
	sectorFour := pixel.Point{X: x, Y: -y}

	pixel.DrawPixel(sectorOne)
	pixel.DrawPixel(sectorTwo)
	pixel.DrawPixel(sectorThree)
	pixel.DrawPixel(sectorFour)
}
