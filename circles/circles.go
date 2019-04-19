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
	sectorOne := pixel.Point{X: point.X + midpoint.X, Y: point.Y + midpoint.Y}
	sectorTwo := pixel.Point{X: point.Y + midpoint.Y, Y: point.X + midpoint.X}
	sectorThree := pixel.Point{X: -point.X + midpoint.X, Y: point.Y + midpoint.Y}
	sectorFour := pixel.Point{X: -point.Y + midpoint.Y, Y: point.X + midpoint.X}
	sectorFive := pixel.Point{X: -point.X + midpoint.X, Y: -point.Y + midpoint.Y}
	sectorSix := pixel.Point{X: -point.Y + midpoint.Y, Y: -point.X + midpoint.X}
	sectorSeven := pixel.Point{X: point.X + midpoint.X, Y: -point.Y + midpoint.Y}
	sectorEight := pixel.Point{X: point.Y + midpoint.Y, Y: -point.X + midpoint.X}

	pixel.DrawPixel(sectorOne)
	pixel.DrawPixel(sectorTwo)
	pixel.DrawPixel(sectorThree)
	pixel.DrawPixel(sectorFour)
	pixel.DrawPixel(sectorFive)
	pixel.DrawPixel(sectorSix)
	pixel.DrawPixel(sectorSeven)
	pixel.DrawPixel(sectorEight)
}
