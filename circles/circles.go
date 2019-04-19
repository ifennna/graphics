package circles

import "go-graphics/pixel"

func Draw(midpoint pixel.Point, radius float64, algorithm string) {
	switch algorithm {
	case "bresenhams":
		drawBresenhams(midpoint, radius)
	case "midpoint":
		drawMidpoint(midpoint, radius)
	}
}

func drawMirroredPixels(midpoint pixel.Point, point pixel.Point) {
	x := midpoint.X + point.X
	y := midpoint.Y + point.Y

	sectorOne := pixel.Point{X: x, Y: y}
	sectorTwo := pixel.Point{X: y, Y: x}
	sectorThree := pixel.Point{X: -x, Y: y}
	sectorFour := pixel.Point{X: -y, Y: x}
	sectorFive := pixel.Point{X: -x, Y: -y}
	sectorSix := pixel.Point{X: -y, Y: -x}
	sectorSeven := pixel.Point{X: x, Y: -y}
	sectorEight := pixel.Point{X: y, Y: -x}

	pixel.DrawPixel(sectorOne)
	pixel.DrawPixel(sectorTwo)
	pixel.DrawPixel(sectorThree)
	pixel.DrawPixel(sectorFour)
	pixel.DrawPixel(sectorFive)
	pixel.DrawPixel(sectorSix)
	pixel.DrawPixel(sectorSeven)
	pixel.DrawPixel(sectorEight)
}
